package yiface

type IRequest interface {
	// GetConn 获得连接
	GetConn() IConnection
	// GetData 获得消息的数据
	GetData() []byte
	// GetMsgId 获得消息的Id
	GetMsgId() uint32
}
