package configext

import (
	"fmt"
	"sort"
	"strings"

	"github.com/zjh-tech/appcommon/config"
	"github.com/zjh-tech/go-ebase/convert"
)

type ConfigExtMgr struct {
	ChessExchangeMap map[uint32]uint32              //RealId-Tid
	RoundCreeps      map[uint32][]*RoundSummonCreep //RoundId -RoundSummonCreep
	SortFetters      map[uint32]*SortFetters        //ReadID - SortFetters
	ExtRetrofits     map[uint32]*ExtRetrofit        //EquipType-ExtRetrofit
}

func NewConfigExtMgr() *ConfigExtMgr {
	return &ConfigExtMgr{
		ChessExchangeMap: make(map[uint32]uint32),
		RoundCreeps:      make(map[uint32][]*RoundSummonCreep),
		SortFetters:      make(map[uint32]*SortFetters),
	}
}

func (c *ConfigExtMgr) Init() bool {
	for _, chessBaseInfo := range config.GConfigMgr.ChessesBaseCfg.Datas {
		if chessBaseInfo != nil && chessBaseInfo.Star == 1 {
			c.ChessExchangeMap[chessBaseInfo.RealID] = chessBaseInfo.ID
		}
	}

	for _, roundBaseInfo := range config.GConfigMgr.RoundBaseCfg.Datas {
		if roundBaseInfo != nil {
			creeps := make([]*RoundSummonCreep, 0)
			creepStrs := strings.Split(roundBaseInfo.Creeps, "#")
			for _, creepStr := range creepStrs {
				creepInfoStr := strings.Split(creepStr, "-")
				if len(creepInfoStr) == 2 {
					creepId, _ := convert.StrToUint32(creepInfoStr[0])
					pos, _ := convert.StrToUint32(creepInfoStr[1])
					creep := NewRoundSummonCreep(creepId, pos)
					creeps = append(creeps, creep)
				}
			}
			if len(creeps) != 0 {
				c.RoundCreeps[roundBaseInfo.ID] = creeps
			}
		}
	}

	for _, fetterBaseInfo := range config.GConfigMgr.FettersBaseCfg.Datas {
		if sortFetter, ok := c.SortFetters[fetterBaseInfo.RealID]; ok {
			(*sortFetter).FettersBases = append((*sortFetter).FettersBases, fetterBaseInfo)
		} else {
			sortFetter := NewSortFetters()
			(*sortFetter).FettersBases = append((*sortFetter).FettersBases, fetterBaseInfo)
			c.SortFetters[fetterBaseInfo.RealID] = sortFetter
		}
	}

	for _, SortFetter := range c.SortFetters {
		sort.Sort(SortFettersByNumActivations{*SortFetter})
	}

	for _, retrofitBaseInfo := range config.GConfigMgr.RetrofitBaseCfg.Datas {
		var extRetrofit *ExtRetrofit
		if tempExtRetrofit, ok := c.ExtRetrofits[retrofitBaseInfo.Type]; !ok {
			extRetrofit = NewExtRetrofit()
			c.ExtRetrofits[retrofitBaseInfo.Type] = extRetrofit
		} else {
			extRetrofit = tempExtRetrofit
		}

		var extRetrofitBelongs *ExtRetrofitBelong
		if tempExtRetrofitBelongs, ok := extRetrofit.Belongs[retrofitBaseInfo.Belong]; !ok {
			extRetrofitBelongs = NewExtRetrofitBelong()
			extRetrofit.Belongs[retrofitBaseInfo.Belong] = extRetrofitBelongs
		} else {
			extRetrofitBelongs = tempExtRetrofitBelongs
		}

		extRetrofitBelongs.RetrofitBases = append(extRetrofitBelongs.RetrofitBases, retrofitBaseInfo)
	}

	for _, extRetrofit := range c.ExtRetrofits {
		for _, extRetrofitBelong := range extRetrofit.Belongs {
			for _, extRetrofitBelongInfo := range extRetrofitBelong.RetrofitBases {
				extRetrofitBelong.TotalWeight += int(extRetrofitBelongInfo.Weights)
			}
		}
	}

	return true
}

func (c *ConfigExtMgr) GetExchangeInfo(realId uint32) (uint32, error) {
	tid, ok := c.ChessExchangeMap[realId]
	if ok {
		return tid, nil
	}

	return 0, fmt.Errorf("[ConfigExtMgr] GetExchangeInfo RealId=%v Not Find", realId)
}

func (c *ConfigExtMgr) GetRoundSummonCreeps(roundId uint32) ([]*RoundSummonCreep, bool) {
	if creeps, ok := c.RoundCreeps[roundId]; ok {
		return creeps, true
	} else {
		return nil, false
	}
}

var GConfigExtMgr *ConfigExtMgr

func init() {
	GConfigExtMgr = NewConfigExtMgr()
}
