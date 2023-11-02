package main

import (
	"Yis/ynet"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("Client is Start ...")
	time.Sleep(2 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("dial error", err)
		return
	}
	for {
		dp := ynet.NewDataPack()
		//发包
		msg := ynet.NewMessagePack(2, []byte("22222222"))
		binaryMsg, err2 := dp.Pack(msg)
		if err2 != nil {
			return
		}
		_, err := conn.Write(binaryMsg)
		if err != nil {
			return
		}
		//拆包
		dataHead := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, dataHead)
		if err != nil {
			return
		}
		msgRecv, err := dp.UnPack(dataHead)
		if err != nil {
			return
		}
		var data []byte
		if msgRecv.GetLength() > 0 {
			data = make([]byte, msgRecv.GetLength())
			_, err := io.ReadFull(conn, data)
			if err != nil {
				return
			}
		}
		msgRecv.SetData(data)

		fmt.Println("===>recv from client msgId:", msgRecv.GetId(), " data:", string(msgRecv.GetData()))
		time.Sleep(time.Second * 2)
	}
}
