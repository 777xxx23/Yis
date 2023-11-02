package yiface

type IMsgHandler interface {
	// AddRouter 添加路由方法
	AddRouter(msgId uint32, router IRouter)

	// DoMsgHandler 处理消息方法
	DoMsgHandler(request IRequest)

	// StartWorkerPool 启动协程池的方法
	StartWorkerPool()

	// SendReqToTaskQueue 向消息队列发送req的方法
	SendReqToTaskQueue(request IRequest)
}
