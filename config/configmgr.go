package config

import (
    "fmt"
    "io/ioutil"
    "os"
    "reflect"
    "github.com/golang/protobuf/proto"
)

type ConfigMgr struct {
    DirPath string
    ChessesStatisticsCfg    *ChessesStatisticsCfg
    ChessesSearchCfg    *ChessesSearchCfg
    GlobalErrorcodeCfg    *GlobalErrorcodeCfg
    MasterBaseCfg    *MasterBaseCfg
    MasterLevelCfg    *MasterLevelCfg
    RoundBaseCfg    *RoundBaseCfg
    TipErrorcodeCfg    *TipErrorcodeCfg
}

func (c *ConfigMgr) LoadAllCfg(dirPath string) error {
    c.DirPath = dirPath

    c.ChessesStatisticsCfg = &ChessesStatisticsCfg{}
    c.ChessesSearchCfg = &ChessesSearchCfg{}
    c.GlobalErrorcodeCfg = &GlobalErrorcodeCfg{}
    c.MasterBaseCfg = &MasterBaseCfg{}
    c.MasterLevelCfg = &MasterLevelCfg{}
    c.RoundBaseCfg = &RoundBaseCfg{}
    c.TipErrorcodeCfg = &TipErrorcodeCfg{}

    cfgs := map[string]proto.Message{
        "ChessesStatistics": c.ChessesStatisticsCfg,
        "ChessesSearch": c.ChessesSearchCfg,
        "GlobalErrorcode": c.GlobalErrorcodeCfg,
        "MasterBase": c.MasterBaseCfg,
        "MasterLevel": c.MasterLevelCfg,
        "RoundBase": c.RoundBaseCfg,
        "TipErrorcode": c.TipErrorcodeCfg,
    } 

    for fileName, msg := range cfgs {
        if err := LoadConfig(dirPath, fileName, msg); err != nil {
            return err
        }
    }

    return nil
}

func (c *ConfigMgr) GetChessesStatisticsByID(id uint32) *ChessesStatistics {
    if c.ChessesStatisticsCfg.Datas != nil {
        if info, ok := c.ChessesStatisticsCfg.Datas[id]; ok {
            return info
        }
    }
    return nil
}

func (c *ConfigMgr) GetChessesSearchByID(id uint32) *ChessesSearch {
    if c.ChessesSearchCfg.Datas != nil {
        if info, ok := c.ChessesSearchCfg.Datas[id]; ok {
            return info
        }
    }
    return nil
}

func (c *ConfigMgr) GetGlobalErrorcodeByID(id uint32) *GlobalErrorcode {
    if c.GlobalErrorcodeCfg.Datas != nil {
        if info, ok := c.GlobalErrorcodeCfg.Datas[id]; ok {
            return info
        }
    }
    return nil
}

func (c *ConfigMgr) GetMasterBaseByID(id uint32) *MasterBase {
    if c.MasterBaseCfg.Datas != nil {
        if info, ok := c.MasterBaseCfg.Datas[id]; ok {
            return info
        }
    }
    return nil
}

func (c *ConfigMgr) GetMasterLevelByID(id uint32) *MasterLevel {
    if c.MasterLevelCfg.Datas != nil {
        if info, ok := c.MasterLevelCfg.Datas[id]; ok {
            return info
        }
    }
    return nil
}

func (c *ConfigMgr) GetRoundBaseByID(id uint32) *RoundBase {
    if c.RoundBaseCfg.Datas != nil {
        if info, ok := c.RoundBaseCfg.Datas[id]; ok {
            return info
        }
    }
    return nil
}

func (c *ConfigMgr) GetTipErrorcodeByID(id uint32) *TipErrorcode {
    if c.TipErrorcodeCfg.Datas != nil {
        if info, ok := c.TipErrorcodeCfg.Datas[id]; ok {
            return info
        }
    }
    return nil
}

func (c *ConfigMgr) SetChessesStatisticsCfg(cfg *ChessesStatisticsCfg) { 
    c.ChessesStatisticsCfg = cfg
} 

func (c *ConfigMgr) SetChessesSearchCfg(cfg *ChessesSearchCfg) { 
    c.ChessesSearchCfg = cfg
} 

func (c *ConfigMgr) SetGlobalErrorcodeCfg(cfg *GlobalErrorcodeCfg) { 
    c.GlobalErrorcodeCfg = cfg
} 

func (c *ConfigMgr) SetMasterBaseCfg(cfg *MasterBaseCfg) { 
    c.MasterBaseCfg = cfg
} 

func (c *ConfigMgr) SetMasterLevelCfg(cfg *MasterLevelCfg) { 
    c.MasterLevelCfg = cfg
} 

func (c *ConfigMgr) SetRoundBaseCfg(cfg *RoundBaseCfg) { 
    c.RoundBaseCfg = cfg
} 

func (c *ConfigMgr) SetTipErrorcodeCfg(cfg *TipErrorcodeCfg) { 
    c.TipErrorcodeCfg = cfg
} 

func LoadConfig(dirPath string, fileName string, message proto.Message) error {
    PbName := dirPath + "/" + fileName + ".pb"
    file, openErr := os.Open(PbName)
    if openErr != nil {
        return openErr
    }
    defer file.Close()
    datas, readErr := ioutil.ReadAll(file)
    if readErr != nil {
        return readErr
    }
    return proto.Unmarshal(datas, message)
}

func ReloadConfig(dirPath string, fileName string) (error, proto.Message) {
    className := fileName + "Cfg"
    obj := GStringToStructMgr.GetStructObj(className)
    if err := LoadConfig(dirPath, fileName, obj.(proto.Message)); err != nil {
        return err, nil
    } 

    return nil, obj.(proto.Message)
} 
func NewChessesStatisticsCfg() interface{} { 
    return &ChessesStatisticsCfg{}
} 

func NewChessesSearchCfg() interface{} { 
    return &ChessesSearchCfg{}
} 

func NewGlobalErrorcodeCfg() interface{} { 
    return &GlobalErrorcodeCfg{}
} 

func NewMasterBaseCfg() interface{} { 
    return &MasterBaseCfg{}
} 

func NewMasterLevelCfg() interface{} { 
    return &MasterLevelCfg{}
} 

func NewRoundBaseCfg() interface{} { 
    return &RoundBaseCfg{}
} 

func NewTipErrorcodeCfg() interface{} { 
    return &TipErrorcodeCfg{}
} 

func RegisterAllExcelConfig() { 
    GStringToStructMgr.Register("ChessesStatisticsCfg", NewChessesStatisticsCfg)
    GStringToStructMgr.Register("ChessesSearchCfg", NewChessesSearchCfg)
    GStringToStructMgr.Register("GlobalErrorcodeCfg", NewGlobalErrorcodeCfg)
    GStringToStructMgr.Register("MasterBaseCfg", NewMasterBaseCfg)
    GStringToStructMgr.Register("MasterLevelCfg", NewMasterLevelCfg)
    GStringToStructMgr.Register("RoundBaseCfg", NewRoundBaseCfg)
    GStringToStructMgr.Register("TipErrorcodeCfg", NewTipErrorcodeCfg)
} 

type ExcelCfgReloadMgr struct {
    funcs map[string]reflect.Value
} 

func NewExcelCfgReloadMgr() *ExcelCfgReloadMgr {
    return &ExcelCfgReloadMgr{
        funcs: make(map[string]reflect.Value),
    } 
} 

func (r *ExcelCfgReloadMgr) Init(x interface{}) {
    RegisterAllExcelConfig()
    t := reflect.TypeOf(x)
    v := reflect.ValueOf(x)
    for i := 0; i < t.NumMethod(); i++ {
        methodName := t.Method(i).Name
        r.funcs[methodName] = v.MethodByName(methodName)
    } 
} 

func (r *ExcelCfgReloadMgr) Call(methodName string, message proto.Message) error {
    if info, ok := r.funcs[methodName]; ok {
        info.Call([]reflect.Value{reflect.ValueOf(message)})
        return nil
    } 

    return fmt.Errorf("ExcelCfgReloadMgr Not Find MethodName: %v", methodName)
} 


var GConfigMgr *ConfigMgr
var GExcelCfgReloadMgr *ExcelCfgReloadMgr

func init() {
    GConfigMgr = &ConfigMgr{}
    GExcelCfgReloadMgr = NewExcelCfgReloadMgr()
}
