package yiface

type IMessage interface {
	/*id,data,length 的getter和setter*/
	GetId() uint32
	GetData() []byte
	GetLength() uint32

	SetId(uint32)
	SetData([]byte)
	SetLength(uint32)
}
