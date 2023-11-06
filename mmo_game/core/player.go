package core

import (
	"Yis/mmo_game/pb"
	"Yis/yiface"
	"fmt"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"sync"
)

type Player struct {
	Pid  int
	Conn yiface.IConnection //对应的IConn
	X    float32            //平面的x坐标
	Y    float32            //高度
	Z    float32            //平面的y坐标
	V    float32            //旋转角度0-360度
}

var PIDGen int = 0     //PID生成器
var PIDLock sync.Mutex //PID保护锁

func NewPlayer(conn yiface.IConnection) *Player {
	PIDLock.Lock()
	PIDGen++
	pid := PIDGen
	PIDLock.Unlock()

	player := &Player{
		Pid:  pid,
		Conn: conn,
		X:    float32(150 + rand.Intn(20)),
		Y:    0, //高度为0
		Z:    float32(150 + rand.Intn(20)),
		V:    0, //旋转角度为0
	}

	return player
}

// SendMsg 向客户端发送消息接口
func (p *Player) SendMsg(msgID uint32, data proto.Message) {
	fmt.Printf("before Marshal data = %+v\n", data)
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("PlayID:", p.Pid, " Marshal msg fail, err:", err)
		return
	}
	fmt.Printf("after Marshal data = %+v\n", msg)

	if p.Conn == nil {
		fmt.Println("PlayerID:", p.Pid, " Conn is nil, can't SendMsg...")
		return
	}

	err = p.Conn.SendMsg(msgID, msg)
	if err != nil {
		fmt.Println("PlayID:", p.Pid, " Conn SendMsg fail, err:", err)
		return
	}
}

// SyncPid 向客户端同步PID，msgID：1
func (p *Player) SyncPid() {
	data := &pb.SyncPid{
		Pid: int32(p.Pid),
	}
	p.SendMsg(1, data)
}

// BroadCastStartPosition 广播起始位置，msgID：200
func (p *Player) BroadCastStartPosition() {
	data := &pb.BroadCast{
		Pid: int32(p.Pid),
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	p.SendMsg(200, data)
}
