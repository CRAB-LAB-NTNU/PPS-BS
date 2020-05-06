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
	NUMacd int
	Val    float64
	Z      float64
}
