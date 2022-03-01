package configext

import "github.com/zjh-tech/appcommon/config"

type RoundSummonCreep struct {
	Id  uint32
	Pos uint32
}

func NewRoundSummonCreep(id uint32, pos uint32) *RoundSummonCreep {
	return &RoundSummonCreep{
		Id:  id,
		Pos: pos,
	}
}

//------------------------------------
type SortFetters struct {
	FettersBases []*config.FettersBase
}

func NewSortFetters() *SortFetters {
	return &SortFetters{
		FettersBases: make([]*config.FettersBase, 0),
	}
}

func (s SortFetters) Len() int {
	return len(s.FettersBases)
}

func (s SortFetters) Swap(i, j int) {
	s.FettersBases[i], s.FettersBases[j] = s.FettersBases[j], s.FettersBases[i]
}

type SortFettersByNumActivations struct {
	SortFetters
}

func (s SortFettersByNumActivations) Less(i, j int) bool {
	return s.FettersBases[i].NumActivations >= s.FettersBases[j].NumActivations
}

//---------------------------------

type ExtRetrofitBelong struct {
	RetrofitBases []*config.RetrofitBase //Belong
	TotalWeight   int
}

func NewExtRetrofitBelong() *ExtRetrofitBelong {
	return &ExtRetrofitBelong{
		RetrofitBases: make([]*config.RetrofitBase, 0),
	}
}

type ExtRetrofit struct {
	Belongs map[uint32]*ExtRetrofitBelong //Belong - ExtRetrofitBelong
}

func NewExtRetrofit() *ExtRetrofit {
	return &ExtRetrofit{
		Belongs: make(map[uint32]*ExtRetrofitBelong),
	}
}
