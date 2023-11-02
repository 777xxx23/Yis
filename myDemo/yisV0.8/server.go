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

func OnConnStart(conn yiface.IConnection) {
	fmt.Println("[Conn] OnConnStart func is Called...")
	//在连接建立时设置属性
	conn.AddProperty("name", "Ross")
	conn.AddProperty("school", "Dgut")
	fmt.Println("[Conn] AddProperty succ...")
}

func OnConnStop(conn yiface.IConnection) {
	fmt.Println("[Conn] OnConnStop func is Called...")
	//在连接断开时打印属性
	if name, err := conn.GetProperty("name"); err == nil {
		fmt.Println("[Conn] Conn Property Name = ", name)
	}

	if school, err := conn.GetProperty("school"); err == nil {
		fmt.Println("[Conn] Conn Property Home = ", school)
	}
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
