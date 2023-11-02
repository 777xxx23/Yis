package main

import (
	"Yis/ynet"
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("listen error : ", err)
	}
	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("accept error : ", err2)
			return
		}

		go func(conn net.Conn) {
			//创建解包对象
			dp := ynet.NewDataPack()
			for {
				//先读取头部
				dataHead := make([]byte, dp.GetHeadLen())
				_, err3 := io.ReadFull(conn, dataHead)
				if err3 != nil {
					fmt.Println("io readFull error : ", err3)
					return
				}
				//解包头部
				msgHead, err4 := dp.UnPack(dataHead)
				if err4 != nil {
					return
				}
				//再根据头部信息读取数据
				if msgHead.GetLength() > 0 {
					msg := msgHead.(*ynet.Message)
					msg.Data = make([]byte, msg.GetLength())
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("io readFull2 error : ", err)
						return
					}

					fmt.Println("=====>get msgId: ", msg.GetId(), " msgLength: ", msg.GetLength(), " msg: ", string(msg.GetData()))
				}
			}
		}(conn)

	}
}
