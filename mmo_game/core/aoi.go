package core

import "fmt"

type AOIManager struct {
	MinX  int           //左边界坐标
	MaxX  int           //右边界坐标
	CntX  int           //x方向的格子数量
	MinY  int           //上边界坐标
	MaxY  int           //下边界坐标
	CntY  int           //y方向的格子数量
	Grids map[int]*Grid //格子集合
}

func NewAOIManager(minX int, maxX int, cntX int, minY int, maxY int, cntY int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntX:  cntX,
		MinY:  minY,
		MaxY:  maxY,
		CntY:  cntY,
		Grids: make(map[int]*Grid),
	}

	for i := 0; i < cntY; i++ {
		for j := 0; j < cntX; j++ {
			gid := i*cntX + j
			aoiMgr.Grids[gid] = NewGrid(gid,
				minX+j*aoiMgr.getWidth(),
				minX+(j+1)*aoiMgr.getWidth(),
				minY+i*aoiMgr.getLength(),
				minY+(i+1)*aoiMgr.getLength())
		}
	}
	return aoiMgr
}

func (aoiMgr *AOIManager) getWidth() int {
	return (aoiMgr.MaxX - aoiMgr.MinX) / aoiMgr.CntX
}

func (aoiMgr *AOIManager) getLength() int {
	return (aoiMgr.MaxY - aoiMgr.MinY) / aoiMgr.CntY
}

// GetSurGIDs 通过GID获取周围格子
func (aoiMgr *AOIManager) GetSurGIDs(gid int) (gids []int) {
	//获取x坐标
	x := gid % aoiMgr.CntX
	gids = append(gids, gid)
	//获取x方向的gid
	if x == 0 && aoiMgr.CntX != 1 {
		gids = append(gids, gid+1)
	} else if x == aoiMgr.CntX-1 && aoiMgr.CntX != 1 {
		gids = append(gids, gid-1)
	} else if aoiMgr.CntX > 2 {
		gids = append(gids, gid-1)
		gids = append(gids, gid+1)
	}
	tmp := gids
	//获取y方向的gid
	for _, v := range tmp {
		//获取y坐标
		y := v / aoiMgr.CntX
		if y == 0 && aoiMgr.CntY != 1 {
			gids = append(gids, v+aoiMgr.CntX)
		} else if y == aoiMgr.CntY-1 && aoiMgr.CntY != 1 {
			gids = append(gids, v-aoiMgr.CntX)
		} else if aoiMgr.CntY > 2 {
			gids = append(gids, v+aoiMgr.CntX)
			gids = append(gids, v-aoiMgr.CntX)
		}
	}
	return
}

// GetGidByPos 通过pos获取GID
func (aoiMgr *AOIManager) GetGidByPos(posX, posY float32) int {
	//获取gid
	x := (int(posX) - aoiMgr.MinX) / aoiMgr.getWidth()
	y := (int(posY) - aoiMgr.MinY) / aoiMgr.getLength()
	gid := y*aoiMgr.CntX + x
	return gid
}

// GetSurPlayIDsByPos 通过玩家position获取周围玩家ID
func (aoiMgr *AOIManager) GetSurPlayIDsByPos(posX, posY float32) (playIDs []int) {
	//获取gid
	gid := aoiMgr.GetGidByPos(posX, posY)
	//获取周围gid
	gids := aoiMgr.GetSurGIDs(gid)
	for _, v := range gids {
		playIDs = append(playIDs, aoiMgr.Grids[v].GetPlayerIds()...)
	}

	return
}

// GetPidsByGid 通过GID获取当前格子的全部playerID
func (m *AOIManager) GetPidsByGid(gID int) (playerIDs []int) {
	playerIDs = m.Grids[gID].GetPlayerIds()
	return
}

// RemovePidFromGrid 移除一个格子中的PlayerID
func (m *AOIManager) RemovePidFromGrid(pID, gID int) {
	m.Grids[gID].RemovePlayer(pID)
}

// AddPidToGrid 添加一个PlayerID到一个格子中
func (m *AOIManager) AddPidToGrid(pID, gID int) {
	m.Grids[gID].AddPlayer(pID)
}

// AddToGridByPos 通过横纵坐标添加一个Player到一个格子中
func (m *AOIManager) AddToGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	m.Grids[gID].AddPlayer(pID)
}

// RemoveFromGridByPos 通过横纵坐标把一个Player从对应的格子中删除
func (m *AOIManager) RemoveFromGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	m.Grids[gID].RemovePlayer(pID)
}

func (aoiMgr *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager MinX:%d,MaxX:%d,CntX:%d,MinY:%d,MaxY:%d,CntY:%d\n",
		aoiMgr.MinX, aoiMgr.MaxX, aoiMgr.CntX, aoiMgr.MinY, aoiMgr.MaxY, aoiMgr.CntY)
	for _, v := range aoiMgr.Grids {
		s += fmt.Sprintln(v)
	}

	return s
}
