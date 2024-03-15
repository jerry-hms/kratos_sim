package websocket

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
)

type Any interface{}

type MessagePayload Any
type MessageType uint32

type BinaryMessage struct {
	Type MessageType
	Body []byte
}

func (b *BinaryMessage) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, uint32(b.Type)); err != nil {
		return nil, err
	}

	buf.Write(b.Body)
	return buf.Bytes(), nil
}

func (b *BinaryMessage) Unmarshal(buf []byte) error {
	network := new(bytes.Buffer)
	network.Write(buf)
	if err := binary.Read(network, binary.LittleEndian, &b.Type); err != nil {
		return err
	}
	b.Body = network.Bytes()
	return nil
}

type TextMessage struct {
	Type MessageType            `json:"type" xml:"type"`
	Body map[string]interface{} `json:"body" xml:"body"`
}

func (t *TextMessage) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func (t *TextMessage) Unmarshal(buf []byte) error {
	return json.Unmarshal(buf, t)
}
