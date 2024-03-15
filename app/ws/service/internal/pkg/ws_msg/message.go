package ws_msg

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"kratos_sim/app/ws/service/internal/pkg/ws_msg/interf"
	"kratos_sim/app/ws/service/internal/pkg/ws_msg/types"
	"time"
)

func NewMessage() *SendStru {
	m := &SendStru{}
	return m
}

// SendStru 推送消息体
type SendStru struct {
	Scene       string       `json:"scene"`
	Attachments *Attachments `json:"attachments"`
	Time        int64        `json:"time"`
	MessageId   string       `json:"message_id"`
}

// SetScene 设置发送场景
func (m *SendStru) SetScene(scene string) *SendStru {
	m.Scene = scene
	return m
}

// 设置Attachments
func (m *SendStru) setAttachments() {
	m.Attachments = createAttachments()
}

// 设置发送时间
func (m *SendStru) setTime() {
	m.Time = time.Now().Unix()
}

// 生成消息ID
func (m *SendStru) generateMessageId() {
	m.MessageId = uuid.NewString()
}

// GetMessageId 获取消息ID
func (m *SendStru) GetMessageId() string {
	return m.MessageId
}

// SetSender 设置sender数据
func (m *SendStru) SetSender(sender interface{}) *SendStru {
	if m.Attachments == nil {
		m.Attachments = createAttachments()
	}
	m.Attachments.setSender(sender)
	return m
}

// SetMessage 设置消息内容
func (m *SendStru) SetMessage(message interface{}) *SendStru {
	if m.Attachments == nil {
		m.setAttachments()
	}
	_ = m.Attachments.parseMessage(message)
	return m
}

// SetExtra 设置其他参数
func (m *SendStru) SetExtra(data interface{}) *SendStru {
	if values, ok := data.(map[string]interface{}); ok {
		if m.Attachments == nil {
			m.setAttachments()
		}
		for k, v := range values {
			m.Attachments.Extra[k] = v
		}
	}
	return m
}

func (m *SendStru) GetCtAs() (string, error) {
	return m.Attachments.GetContentAlias()
}

func (m *SendStru) GetMessage() map[string]interface{} {
	return m.Attachments.GetMessage()
}

// Parse 解析消息体
func (m *SendStru) Parse() *SendStru {
	m.setTime()
	m.generateMessageId()

	return m
}

type Message map[string]interface{}

func (m Message) Unmarshal(b interface{}) error {
	var msg Message
	err := msg.Scan(b)
	if err != nil {
		return err
	}
	x, err := msg.Resolve()
	if err != nil {
		return err
	}
	m["type"] = msg.GetType()
	message := x.Parse(msg)
	if err != nil {
		return err
	}
	m[msg.GetType()] = message
	return nil
}

func (m Message) GetType() string {
	return m["type"].(string)
}

func (m *Message) Scan(value interface{}) error {
	b, _ := json.Marshal(value)
	err := json.Unmarshal(b, &m)
	if err != nil {
		return errors.New("invalid Scan Source")
	}
	return nil
}

func (m Message) Resolve() (interf.TypeInterface, error) {
	ty, ets := m["type"]
	if !ets {
		return nil, errors.New("type cannot be empty")
	}

	t := ty.(string)
	x, ets := types.MessageTypes[t]
	if !ets {
		return nil, errors.New("message type is not defined")
	}
	return x, nil
}

type Attachments struct {
	Sender  map[string]interface{} `json:"sender"`
	Message Message                `json:"message"`
	Extra   map[string]interface{} `json:"extra"`
}

func createAttachments() *Attachments {
	return &Attachments{
		Sender:  make(map[string]interface{}),
		Message: make(Message),
		Extra:   make(map[string]interface{}),
	}
}

// 设置sender
func (a *Attachments) setSender(sender interface{}) *Attachments {
	str, _ := json.Marshal(sender)
	_ = json.Unmarshal(str, &a.Sender)
	return a
}

// 解析message，并且填充给Attachments
func (a *Attachments) parseMessage(message interface{}) error {
	msg := make(Message)
	err := msg.Unmarshal(message)
	if err != nil {
		return err
	}
	a.Message = msg
	return nil
}

func (a *Attachments) GetContentAlias() (string, error) {
	msg, err := a.Message.Resolve()
	if err != nil {
		return "", err
	}
	return msg.GetContent(), nil
}

func (m *Attachments) GetMessage() map[string]interface{} {
	return m.Message
}

// 设置其他参数
func (a *Attachments) setExtra(key string, value interface{}) {
	a.Message[key] = value
}
