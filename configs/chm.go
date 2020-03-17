package configs

type Epsilon struct {
	Cp, Tau float64
	Tc      int
}
type ImprovedEpsilon struct {
	Tau, Alpha, Cp float64
	Tc             int
}
type R2S struct {
	FESc   int
	FESacd int
	Cs     int
	Val    float64
	ZMin   float64
}
