package yiface

type IServer interface {
	Start()
	Stop()
	Serve()
	// AddRouter 向路由管理器中添加路由
	AddRouter(msgId uint32, router IRouter)
	// GetConnMgr 获取connManger
	GetConnMgr() IConnManger

	// AddOnConnStart 添加Conn创建时的hook函数
	AddOnConnStart(func(IConnection))
	// AddOnConnStop 添加Conn断开时的hook函数
	AddOnConnStop(func(IConnection))
	// CallOnConnStart 执行Conn创建时的hook函数
	CallOnConnStart(connection IConnection)
	// CallOnConnStop 执行Conn断开时的hook函数
	CallOnConnStop(connection IConnection)
}
