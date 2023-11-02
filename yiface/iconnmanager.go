package yiface

type IConnManger interface {
	// AddConn 添加一个连接
	AddConn(connId uint32, connection IConnection)
	// RemoveConn 删除一个连接
	RemoveConn(connection IConnection)
	// GetConn 通过id得到连接
	GetConn(connId uint32) (IConnection, error)
	// Len 得到所有连接个数
	Len() uint32
	// ClearConn 删除并停止所有连接
	ClearConn()
}
