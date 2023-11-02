package main

import (
	"fmt"
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
		_, err2 := conn.Write([]byte("Yis 0.3"))
		if err2 != nil {
			return
		}

		buf := make([]byte, 512)
		cnt, err3 := conn.Read(buf)
		if err3 != nil {
			fmt.Println("read buf error ", err3)
			return
		}
		fmt.Printf("receive from server data : %s , cnt = %d \n", buf, cnt)
		time.Sleep(1 * time.Second)
	}
}
