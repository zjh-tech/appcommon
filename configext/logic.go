package configext

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
