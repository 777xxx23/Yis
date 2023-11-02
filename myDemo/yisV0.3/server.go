package main

import (
	"Yis/yiface"
	"Yis/ynet"
	"fmt"
)

type MyRouter struct {
	ynet.BaseRouter
}

func (mr *MyRouter) PreHandle(request yiface.IRequest) {
	fmt.Println("Pre Handle Start ...")
	_, err := request.GetConn().GetTCPConnection().Write([]byte("before Ping ....\n"))
	if err != nil {
		fmt.Println("PreHandle error !!!")
	}
}
func (mr *MyRouter) Handle(request yiface.IRequest) {
	fmt.Println("Handle Start ...")
	_, err := request.GetConn().GetTCPConnection().Write([]byte("ping ping Ping ....\n"))
	if err != nil {
		fmt.Println("Handle error !!!")
	}
}
func (mr *MyRouter) PostHandle(request yiface.IRequest) {
	fmt.Println("Post Handle Start...")
	_, err := request.GetConn().GetTCPConnection().Write([]byte("after Ping ....\n"))
	if err != nil {
		fmt.Println("PostHandle error !!!")
	}
}

func main() {
	s := ynet.NewServer("ross")
	s.AddRouter(&MyRouter{})
	s.Serve()
}
