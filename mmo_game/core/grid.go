package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	GID       int          //格子Id
	PlayerIDs map[int]bool //格子玩家集合
	PlayLock  sync.RWMutex //玩家集合保护锁
	MinX      int          //左边界坐标
	MaxX      int          //右边界坐标
	MinY      int          //上边界坐标
	MaxY      int          //下边界坐标
}

func NewGrid(id int, minX int, maxX int, minY int, maxY int) *Grid {
	return &Grid{
		GID:       id,
		PlayerIDs: make(map[int]bool),
		PlayLock:  sync.RWMutex{},
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
	}
}

// AddPlayer 玩家的增加
func (g *Grid) AddPlayer(playerId int) {
	g.PlayLock.Lock()
	defer g.PlayLock.Unlock()
	if _, ok := g.PlayerIDs[playerId]; ok {
		fmt.Println("Already exit playerId ", playerId, " Add fail")
	}
	g.PlayerIDs[playerId] = true
}

// RemovePlayer 玩家的删除
func (g *Grid) RemovePlayer(playerId int) {
	g.PlayLock.Lock()
	defer g.PlayLock.Unlock()
	if _, ok := g.PlayerIDs[playerId]; !ok {
		fmt.Println("Didn't exit playerId ", playerId, " Remove fail")
	}
	delete(g.PlayerIDs, playerId)
}

// GetPlayerIds 获取格子所有玩家
func (g *Grid) GetPlayerIds() (playerIDs []int) {
	g.PlayLock.RLock()
	defer g.PlayLock.RUnlock()

	for id, _ := range g.PlayerIDs {
		playerIDs = append(playerIDs, id)
	}

	return
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grid id: %d, minX:%d, maxX:%d, minY:%d, maxY:%d, playerIDs:%v",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.PlayerIDs)
}
