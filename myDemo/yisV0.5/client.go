package main

import (
	"Yis/ynet"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("dial error : ", err)
		return
	}
	dp := ynet.NewDataPack()

	msg := &ynet.Message{
		Id:     1,
		Length: 6,
		Data:   []byte{'h', 'e', 'l', 'l', 'o', ' '},
	}

	msg2 := &ynet.Message{
		Id:     2,
		Length: 6,
		Data:   []byte{'w', 'o', 'r', 'l', 'd', '!'},
	}
	//封包出两个切片
	pack1, err := dp.Pack(msg)
	if err != nil {
		fmt.Println("pack error ", err)
		return
	}
	pack2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("pack2 error ", err)
		return
	}
	mPackge := append(pack1, pack2...)
	//将切片合并后发送
	_, err = conn.Write(mPackge)
	if err != nil {
		return
	}

	for {
	}
}
