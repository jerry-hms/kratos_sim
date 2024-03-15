package types

import (
	"kratos_sim/app/ws/service/internal/pkg/ws_msg/interf"
)

type Text struct {
	Content string `json:"content"`
}

func (t *Text) Parse(message map[string]interface{}) interf.TypeInterface {
	t.Content = message["text"].(string)
	return t
}

func (t *Text) GetContent() string {
	return t.Content
}
