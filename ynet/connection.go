package ynet

import (
	"Yis/utils"
	"Yis/yiface"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

type Connection struct {
	//所属的服务器
	TcpServer yiface.IServer
	//连接
	Conn *net.TCPConn
	//连接id
	ConnID uint32
	//是否关闭
	isClosed bool

	//是否退出通道
	ExitChan chan bool
	//读写routine通信通道，无缓冲
	msgChan chan []byte
	//读写routine通信通道，有缓冲
	msgBuffChan chan []byte

	//处理消息的一个多路由
	msgHandler yiface.IMsgHandler

	//连接的属性
	property map[string]interface{}
	//保护连接属性的读写锁
	propertyLock sync.RWMutex
}

// NewConnection 初始化并返回一个连接
func NewConnection(s yiface.IServer, conn *net.TCPConn, id uint32, msgHandler yiface.IMsgHandler) yiface.IConnection {
	c := &Connection{
		TcpServer:   s,
		Conn:        conn,
		ConnID:      id,
		isClosed:    false,
		ExitChan:    make(chan bool, 1),
		msgChan:     make(chan []byte),
		msgBuffChan: make(chan []byte, 10),
		msgHandler:  msgHandler,
		property:    make(map[string]interface{}),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("[Reader] goroutine is running ... ")
	defer fmt.Println("[Reader] connId : ", c.ConnID, " reader exit ...")

	//循环读取msg
	for {
		dp := NewDataPack()
		//获取头部字节
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Println("ReadFull in headData error : ", err)
			c.ExitChan <- true
			return
		}
		//解析头部字节
		Msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("dp unpack headData error : ", err)
			c.ExitChan <- true
			continue
		}
		//获取数据信息
		var data []byte
		if Msg.GetLength() > 0 {
			data = make([]byte, Msg.GetLength())
			_, err = io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("ReadFull in data error : ", err)
				c.ExitChan <- true
				continue
			}
		}
		Msg.SetData(data)

		//包装成request传给router处理
		req := &Request{
			Conn: c,
			Msg:  Msg,
		}

		//判断是否使用协程池工作
		if utils.GlobalObject.MaxPackageSize > 0 {
			//将请求发送至消息队列，让协程池处理业务
			c.msgHandler.SendReqToTaskQueue(req)
		} else {
			go c.msgHandler.DoMsgHandler(req)
		}
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("[Writer] goroutine is running ... ")
	defer fmt.Println("[Writer] connId:", c.ConnID, " reader exit ...")
	for {
		//阻塞等待消息的发送
		select {
		//无缓冲通道
		case data := <-c.msgChan:
			_, err := c.GetTCPConnection().Write(data)
			if err != nil {
				fmt.Println("Send msg fail,err:", err)
				return
			}
		//有缓冲通道
		case data, ok := <-c.msgBuffChan:
			if !ok {
				fmt.Println("msgBuffChan has closed")
				break
			}
			_, err := c.GetTCPConnection().Write(data)
			if err != nil {
				fmt.Println("Send msg fail,err:", err)
				return
			}
		//退出通道
		case <-c.ExitChan:
			return
		}
	}
}

// Start 启动连接，让当前的链接开始工作
func (c *Connection) Start() {
	fmt.Println("[Conn] Conn Start() connID: ", c.ConnID)
	//结束后释放资源
	defer c.Stop()
	//启动连接的读业务
	go c.StartReader()
	go c.StartWriter()

	//执行创建连接时的Hook函数
	c.TcpServer.CallOnConnStart(c)

	for {
		//监听连接是否断开
		select {
		case <-c.ExitChan:
			return
		}
	}
}

// Stop 停止连接，结束当前连接的工作
func (c *Connection) Stop() {
	fmt.Println("[Conn] Conn Stop() connID: ", c.ConnID)

	if c.isClosed == true {
		return
	}
	//关闭标志置为true
	c.isClosed = true
	//执行连接关闭时的Hook函数
	c.TcpServer.CallOnConnStop(c)

	//关闭连接
	err := c.Conn.Close()
	if err != nil {
		return
	}
	//将连接从连接管理器中删除
	c.TcpServer.GetConnMgr().RemoveConn(c)
	//关闭通道
	close(c.ExitChan)
	close(c.msgChan)
	close(c.msgBuffChan)
}

// GetTCPConnection 返回当前连接
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 返回连接id
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远端地址
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg 通过无缓冲管道发送数据方法
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	dp := NewDataPack()
	//将msg打包成二进制流
	msg := NewMessagePack(msgId, data)
	byteMsg, err := dp.Pack(msg)
	if err != nil {
		fmt.Println("pack error when SendMsg ,err: ", err)
		return err
	}
	//让写routine去发送数据
	c.msgChan <- byteMsg

	return nil
}

// SendBuffMsg 通过有缓冲管道发送数据方法
func (c *Connection) SendBuffMsg(msgId uint32, data []byte) error {

	dp := NewDataPack()
	//先将数据包装成msg
	msg := NewMessagePack(msgId, data)
	//再将数据打包成字节流
	byteMsg, err := dp.Pack(msg)
	if err != nil {
		fmt.Println("pack error when SendMsg ,err: ", err)
		return err
	}
	//让写routine去发送数据
	c.msgBuffChan <- byteMsg

	return nil
}

// AddProperty 添加属性
func (c *Connection) AddProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	if _, ok := c.property[key]; ok {
		fmt.Println("property key:", key, " already exit, Add fail...")
		return
	}
	c.property[key] = value
}

// RemoveProperty 删除属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	if _, ok := c.property[key]; !ok {
		fmt.Println("property key:", key, " didn't exit, Remove fail...")
		return
	}
	delete(c.property, key)
}

// GetProperty 获得属性值
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	value, ok := c.property[key]
	if !ok {
		fmt.Println("property key:", key, " didn't exit, Remove fail...")
		return nil, errors.New(fmt.Sprintf("property key:%s ,didn't exit, Get fail...", key))
	}
	return value, nil
}
