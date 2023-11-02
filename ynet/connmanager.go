package ynet

import (
	"Yis/yiface"
	"errors"
	"fmt"
	"sync"
)

type ConnManger struct {
	//connId到conn的映射map
	connections map[uint32]yiface.IConnection

	//保护map的读写锁
	connLock sync.RWMutex
}

func NewConnManger() yiface.IConnManger {
	return &ConnManger{
		connections: make(map[uint32]yiface.IConnection),
		connLock:    sync.RWMutex{},
	}
}

// AddConn 添加一个连接
func (cm *ConnManger) AddConn(connId uint32, connection yiface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	if _, ok := cm.connections[connId]; ok {
		fmt.Println("ConnId:", connId, ",already exit, add fail")
		return
	}
	cm.connections[connId] = connection
}

// RemoveConn 删除一个连接
func (cm *ConnManger) RemoveConn(connection yiface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	connId := connection.GetConnID()
	if _, ok := cm.connections[connId]; !ok {
		fmt.Println("ConnId:", connId, ",didn't exit, remove fail")
		return
	}
	delete(cm.connections, connId)
}

// GetConn 通过id得到连接
func (cm *ConnManger) GetConn(connId uint32) (yiface.IConnection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	conn, ok := cm.connections[connId]
	if !ok {
		fmt.Println("ConnId:", connId, ", didn't exit, get fail")
		return nil, errors.New(fmt.Sprintf("ConnId:%d, didn't exit, get fail", connId))
	}
	return conn, nil
}

// Len 得到所有连接个数
func (cm *ConnManger) Len() uint32 {
	return uint32(len(cm.connections))
}

// ClearConn 删除并停止所有连接
func (cm *ConnManger) ClearConn() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for connId, conn := range cm.connections {
		conn.Stop()
		delete(cm.connections, connId)
	}

	fmt.Println("[Cons] All connections have closed ...")
}
