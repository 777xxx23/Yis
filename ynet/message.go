package ynet

import "Yis/yiface"

type Message struct {
	Length uint32
	Id     uint32
	Data   []byte
}

func NewMessagePack(id uint32, data []byte) yiface.IMessage {
	msg := &Message{
		Id:     id,
		Length: uint32(len(data)),
		Data:   data,
	}
	return msg
}

func (m *Message) GetId() uint32 {
	return m.Id
}
func (m *Message) GetData() []byte {
	return m.Data
}
func (m *Message) GetLength() uint32 {
	return m.Length
}

func (m *Message) SetId(id uint32) {
	m.Id = id
}
func (m *Message) SetData(data []byte) {
	m.Data = data
}
func (m *Message) SetLength(length uint32) {
	m.Length = length
}
