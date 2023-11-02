package main

import (
	"Yis/yiface"
	"Yis/ynet"
	"fmt"
)

type PingRouter struct {
	ynet.BaseRouter
}

func (mr *PingRouter) Handle(request yiface.IRequest) {
	fmt.Println("===>recv from client msgId:", request.GetMsgId(), " data:", string(request.GetData()))
	err := request.GetConn().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

type HelloRouter struct {
	ynet.BaseRouter
}

func (mr *HelloRouter) Handle(request yiface.IRequest) {
	fmt.Println("===>recv from client msgId:", request.GetMsgId(), " data:", string(request.GetData()))
	err := request.GetConn().SendMsg(2, []byte("Hello World!"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func OnConnStart(connection yiface.IConnection) {
	fmt.Println("[Conn] OnConnStart func is Called...")
	err := connection.SendBuffMsg(3, []byte("OnConnStart send msg to you"))
	if err != nil {
		return
	}
}

func OnConnStop(connection yiface.IConnection) {
	fmt.Println("[Conn] OnConnStop func is Called...")
}

func main() {
	s := ynet.NewServer("ross")
	//设置连接Hook函数
	s.AddOnConnStart(OnConnStart)
	s.AddOnConnStop(OnConnStop)

	//添加路由
	s.AddRouter(1, &PingRouter{})
	s.AddRouter(2, &HelloRouter{})
	s.Serve()
}
