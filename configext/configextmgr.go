package configext

import (
	"fmt"

	"github.com/zjh-tech/appcommon/config"
)

type ConfigExtMgr struct {
	ChessExchangeMap map[uint32]uint32 //realId-tid
}

func NewConfigExtMgr() *ConfigExtMgr {
	return &ConfigExtMgr{
		ChessExchangeMap: make(map[uint32]uint32),
	}
}

func (c *ConfigExtMgr) Init() bool {
	for _, chessBaseInfo := range config.GConfigMgr.ChessesBaseCfg.Datas {
		if chessBaseInfo != nil && chessBaseInfo.Star == 1 {
			c.ChessExchangeMap[chessBaseInfo.RealID] = chessBaseInfo.ID
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

var GConfigExtMgr *ConfigExtMgr

func init() {
	GConfigExtMgr = NewConfigExtMgr()
}
