package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"kratos_sim/app/ws/service/internal/pkg/ws_msg/interf"
	"reflect"
)

const (
	TextMessage  string = "text"
	ImageMessage string = "image"
)

type messageTypeMap map[string]interf.TypeInterface

var MessageTypes = messageTypeMap{
	TextMessage:  &Text{},
	ImageMessage: &Image{},
}

// Unmarshal 解码
func Unmarshal(data map[string]interface{}, x interf.TypeInterface) error {
	str, _ := json.Marshal(data)
	err := json.Unmarshal(str, &x)
	if err != nil {
		return errors.New(fmt.Sprintf("message.%s 数据绑定失败", reflect.TypeOf(x).Name()))
	}
	return nil
}
