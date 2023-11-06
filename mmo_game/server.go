package main

import (
	"Yis/mmo_game/core"
	"Yis/yiface"
	"Yis/ynet"
	"fmt"
)

func OnConnStart(connection yiface.IConnection) {
	player := core.NewPlayer(connection)
	player.SyncPid()
	player.BroadCastStartPosition()
	fmt.Println("====> Player PID:", player.Pid, " arrived...")
}

func main() {
	server := ynet.NewServer("ross")
	server.AddOnConnStart(OnConnStart)
	server.Serve()
}
