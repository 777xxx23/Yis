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

func main() {
	s := ynet.NewServer("ross")
	s.AddRouter(1, &PingRouter{})
	s.AddRouter(2, &HelloRouter{})
	s.Serve()
}
