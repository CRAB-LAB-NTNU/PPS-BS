package types

// Stage is an enum defining which phase PPS is in. Can either be Push or Pull
type Stage int

const (
	// Push PPS is in the push stage and constraints are ignored
	Push Stage = iota
	// BorderSearch PPS is in the stage between Push and Pull
	BorderSearch
	// Pull PPS is in the pull stage and constraints are handled
	Pull
)
