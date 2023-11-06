package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type GlobalObj struct {
	//ip地址
	Host string
	//端口号
	TcpPort int
	//服务器名字
	Name string
	//服务器版本
	Version string

	//最大包大小
	MaxPackageSize uint32
	//最多连接数量
	MaxConn uint32

	//协程池数量
	WorkerPoolSize uint32
	//任务队列长度
	MaxWorkerTaskLen uint32
}

var GlobalObject *GlobalObj

// Reload 从cong/yis.json文件中重新加载数据
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("./conf/yis.json")
	if err != nil {
		fmt.Println("readFile error, ", err)
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		fmt.Println("unmarshal error, ", err)
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Host:             "127.0.0.1",
		TcpPort:          7777,
		Name:             "ross",
		Version:          "V0.8",
		MaxPackageSize:   4096,
		MaxConn:          200,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}
	GlobalObject.Reload()
}
