package types

import (
	"kratos_sim/app/sim/interface/internal/pkg/ws_msg/interf"
)

type Image struct {
	Url    string `json:"url"`
	Height int64  `json:"height"`
	Width  int64  `json:"width"`
}

func (i *Image) Parse(msg map[string]interface{}) (interf.TypeInterface, error) {
	return DataBindToType(msg, i)
}

func (i *Image) GetContent() string {
	return "[图片]"
}
