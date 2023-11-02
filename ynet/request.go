package ynet

import "Yis/yiface"

type Request struct {
	Conn yiface.IConnection
	Msg  yiface.IMessage
}

// GetConn 获得连接
func (r *Request) GetConn() yiface.IConnection {
	return r.Conn
}

// GetData 获得连接的数据
func (r *Request) GetData() []byte {
	return r.Msg.GetData()
}

// 获得消息的Id
func (r *Request) GetMsgId() uint32 {
	return r.Msg.GetId()
}
