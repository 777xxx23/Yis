package test

import (
	"Yis/mmo_game/core"
	"fmt"
	"testing"
)

func TestAOIMgr(t *testing.T) {
	aoiMgr := core.NewAOIManager(0, 250, 5, 0, 150, 3)
	fmt.Println(aoiMgr)
}

func TestAOIMgrGetSurGids(t *testing.T) {
	aoiMgr := core.NewAOIManager(0, 250, 5, 0, 150, 3)
	fmt.Println(aoiMgr.GetSurGIDs(6))
	fmt.Println(aoiMgr.GetSurGIDs(0))
	fmt.Println(aoiMgr.GetSurGIDs(5))
	fmt.Println(aoiMgr.GetSurGIDs(10))
	fmt.Println(aoiMgr.GetSurGIDs(11))
}

func TestAOIMgrGetSurPlayIDs(t *testing.T) {
	aoiMgr := core.NewAOIManager(0, 250, 5, 0, 150, 3)
	aoiMgr.Grids[0].AddPlayer(0)
	aoiMgr.Grids[1].AddPlayer(1)
	aoiMgr.Grids[2].AddPlayer(2)
	aoiMgr.Grids[6].AddPlayer(6)
	aoiMgr.Grids[7].AddPlayer(7)
	fmt.Println(aoiMgr.GetGidByPos(175, 50))
	fmt.Println(aoiMgr.GetSurPlayIDsByPos(175, 50))
}
