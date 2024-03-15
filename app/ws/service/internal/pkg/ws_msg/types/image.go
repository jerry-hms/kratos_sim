package types

import "kratos_sim/app/ws/service/internal/pkg/ws_msg/interf"

type Image struct {
	Url    string `json:"url"`
	Height int64  `json:"height"`
	Width  int64  `json:"width"`
}

func (i *Image) Parse(message map[string]interface{}) interf.TypeInterface {
	_ = Unmarshal(message, i)
	return i
}

func (i *Image) GetContent() string {
	return "[图片]"
}
