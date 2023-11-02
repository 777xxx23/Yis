package ynet

import (
	"Yis/utils"
	"Yis/yiface"
	"fmt"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	//多路由管理器
	msgHandler yiface.IMsgHandler
	//链接管理器
	connMgr yiface.IConnManger

	//创建连接时的Hook函数
	OnConnStart func(connection yiface.IConnection)
	//创建断开时的Hook函数
	OnConnStop func(connection yiface.IConnection)
}

func NewServer(name string) yiface.IServer {
	//初始化全局参数
	utils.GlobalObject.Reload()
	//初始化server
	s := &Server{
		Name:        utils.GlobalObject.Name,
		IPVersion:   "tcp4",
		IP:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.TcpPort,
		msgHandler:  NewMsgHandler(),
		connMgr:     NewConnManger(),
		OnConnStart: nil,
		OnConnStop:  nil,
	}
	return s
}

// Start 开始方法
func (s *Server) Start() {
	fmt.Printf("[Start] listening at IP:%s, Port:%d, Name:%s\n", s.IP, s.Port, s.Name)
	fmt.Printf("[Yis] Version: %s, MaxConn: %d,  MaxPacketSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize)
	//创建协程持续监听
	go func() {
		//初始化协程池
		s.msgHandler.StartWorkerPool()

		//初始化套接字地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err ", err)
			return
		}
		//开始监听
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp err ", err)
		}
		fmt.Println("[Server] now is listening...")

		var ConnID uint32
		ConnID = 0

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept tcp err ", err)
				continue
			}
			//判断连接是否超过最大值
			if s.connMgr.Len() > utils.GlobalObject.MaxConn {
				err := conn.Close()
				if err != nil {
					fmt.Println("close conn fail when too many conn...")
					continue
				}
				fmt.Println("too many conn ... ")
				continue
			}

			fmt.Println("[Server] Get One conn ,ConnId:", ConnID, "...")
			//封装得到新的conn
			dealConn := NewConnection(s, conn, ConnID, s.msgHandler)
			//添加进connMgr
			fmt.Println("[Server] Add connId:", ConnID, " to ConnMgr ...")
			s.connMgr.AddConn(ConnID, dealConn)

			//让conn开始工作
			go dealConn.Start()
			//ConnId自增+1
			ConnID++
		}
	}()
}

// Stop 服务器停止方法
func (s *Server) Stop() {
	fmt.Println("[Server] ServerName:", s.Name, ", has Exit")
	//TODO 释放资源
	s.connMgr.ClearConn()
}

// Serve 服务器服务方法
func (s *Server) Serve() {
	s.Start()
	//TODO 做一些服务器启动后的额外业务
	select {}
}

// AddRouter 向路由管理器中添加路由
func (s *Server) AddRouter(msgId uint32, router yiface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
	fmt.Println("add router succ! ...")
}

// GetConnMgr 获取connManger
func (s *Server) GetConnMgr() yiface.IConnManger {
	return s.connMgr
}

// AddOnConnStart 添加Conn创建时的hook函数
func (s *Server) AddOnConnStart(onStart func(yiface.IConnection)) {
	s.OnConnStart = onStart
}

// AddOnConnStop 添加Conn断开时的hook函数
func (s *Server) AddOnConnStop(onStop func(yiface.IConnection)) {
	s.OnConnStop = onStop
}

// CallOnConnStart 执行Conn创建时的hook函数
func (s *Server) CallOnConnStart(connection yiface.IConnection) {
	if s.OnConnStart != nil {
		s.OnConnStart(connection)
	}
}

// CallOnConnStop 执行Conn断开时的hook函数
func (s *Server) CallOnConnStop(connection yiface.IConnection) {
	if s.OnConnStop != nil {
		s.OnConnStop(connection)
	}
}
