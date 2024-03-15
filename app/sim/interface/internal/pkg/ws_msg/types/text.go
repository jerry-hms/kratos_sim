package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"kratos_sim/app/sim/interface/internal/pkg/ws_msg/interf"
	"reflect"
)

type Text struct {
	Content string `json:"content"`
}

func (t *Text) Parse(msg map[string]interface{}) (interf.TypeInterface, error) {
	str, _ := json.Marshal(msg)
	err := json.Unmarshal(str, &t)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("message.%s 数据绑定失败", reflect.TypeOf(t).Name()))
	}
	return t, nil
}

func (t *Text) GetContent() string {
	return t.Content
}
