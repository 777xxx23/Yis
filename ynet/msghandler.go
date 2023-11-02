package ynet

import (
	"Yis/utils"
	"Yis/yiface"
	"fmt"
)

type MsgHandler struct {
	Apis         map[uint32]yiface.IRouter
	WorkPoolSize uint32
	TaskQueue    []chan yiface.IRequest
}

func NewMsgHandler() yiface.IMsgHandler {
	mh := &MsgHandler{
		Apis:         make(map[uint32]yiface.IRouter),
		WorkPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:    make([]chan yiface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
	return mh
}

func (mh *MsgHandler) StartOneWorker(workerId int, taskQueue chan yiface.IRequest) {
	fmt.Println("WorkerId:", workerId, " is working...")
	for {
		select {
		case req := <-taskQueue:
			mh.DoMsgHandler(req)
		}
	}
}

// StartWorkerPool 启动协程池的方法
func (mh *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(mh.WorkPoolSize); i++ {
		//每个协程的消息队列需要初始化
		mh.TaskQueue[i] = make(chan yiface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// SendReqToTaskQueue 向消息队列发送req的方法
func (mh *MsgHandler) SendReqToTaskQueue(request yiface.IRequest) {
	workerId := request.GetConn().GetConnID() % mh.WorkPoolSize
	fmt.Println("Send req to WorkerId:", workerId, ",connId:", request.GetConn().GetConnID())
	mh.TaskQueue[workerId] <- request
}

// AddRouter 添加路由方法
func (mh *MsgHandler) AddRouter(msgId uint32, router yiface.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		fmt.Println("add router error : repeat api")
		return
	}
	mh.Apis[msgId] = router
	fmt.Println("Add router succ handler MsgId:", msgId)
}

// DoMsgHandler 处理消息方法
func (mh *MsgHandler) DoMsgHandler(request yiface.IRequest) {
	router, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("Can't handler MsgId:", request.GetMsgId())
		return
	}
	router.PostHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}
