package testSuite

import (
	"fmt"
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func fuzzyTest(x types.Genotype) {
	if x[0] < 4*x[1] && x[2] < x[0]+x[1] {
		x[4] = 2*x[3] + 10
	} else if x[0] < 4*x[1] && x[2] > x[0]+x[1] {
		x[4] = x[1] + 50
	}
}

func convertGenotype(x types.Genotype) types.Genotype {
	a := make([]float64, len(x))
	a[0] = conv(x[0], 10.0, 150.0)
	a[1] = conv(x[1], 10.0, 150.0)
	a[2] = conv(x[2], 100.0, 200.0)
	a[3] = conv(x[3], 0.0, 50.0)
	a[4] = conv(x[4], 10.0, 150.0)
	a[5] = conv(x[5], 100.0, 300.0)
	a[6] = conv(x[6], 1.0, 3.14)
	return a
}

func conv(val, min, max float64) float64 {
	return min + ((max - min) * val)
}

func fk(x types.Genotype, z, P float64) float64 {
	_, b, c, _, _, _, _ := x[0], x[1], x[2], x[3], x[4], x[5], x[6]
	alpha, beta := alpha(x, z), beta(x, z)

	return P * b * math.Sin(alpha+beta) / (2 * c * math.Cos(alpha))

}

func ang(x types.Genotype, z float64) float64 {
	_, _, _, e, _, l, _ := x[0], x[1], x[2], x[3], x[4], x[5], x[6]
	fmt.Println("ANG", math.Atan(e/(l-z)))
	return math.Atan(e / (l - z))
}

func alpha(x types.Genotype, z float64) float64 {
	a, b, _, e, _, l, _ := x[0], x[1], x[2], x[3], x[4], x[5], x[6]
	g := math.Sqrt(math.Pow(l-z, 2) + math.Pow(e, 2))
	alpha := math.Acos((math.Pow(a, 2) + math.Pow(g, 2) - math.Pow(b, 2)) / (2 * a * g))
	fmt.Println("ALPHA", alpha)
	return alpha + ang(x, z)
}

func beta(x types.Genotype, z float64) float64 {
	a, b, _, e, _, l, _ := x[0], x[1], x[2], x[3], x[4], x[5], x[6]
	g := math.Sqrt(math.Pow(l-z, 2) + math.Pow(e, 2))
	beta := math.Acos((math.Pow(b, 2) + math.Pow(g, 2) - math.Pow(a, 2)) / (2 * b * g))
	fmt.Println("BETA", beta)
	return beta - ang(x, z)
}

func robotObjective1(x types.Genotype, z, P float64) float64 {
	return P / fk(x, z, P)
}

func robotObjective2(x types.Genotype) float64 {
	a, b, c, e, _, l, _ := x[0], x[1], x[2], x[3], x[4], x[5], x[6]
	return a + b + c + e + l
}

func y(x types.Genotype, z float64) float64 {
	_, _, c, e, f, _, delta := x[0], x[1], x[2], x[3], x[4], x[5], x[6]
	beta := beta(x, z)
	return 2 * (e + f + c*math.Sin(beta+delta))
}

func robotConstraint1(x types.Genotype, yMin, zMax float64) float64 {
	return yMin - y(x, zMax)
}

func robotConstraint2(x types.Genotype, zMax float64) float64 {
	return y(x, zMax)
}

func robotConstraint3(x types.Genotype, yMax float64) float64 {
	return y(x, 0) - yMax
}

func robotConstraint4(x types.Genotype, yG float64) float64 {
	return yG - y(x, 0)
}

func robotConstraint5(x types.Genotype) float64 {
	a, b, _, e, _, l, _ := x[0], x[1], x[2], x[3], x[4], x[5], x[6]
	return math.Pow(a+b, 2) - math.Pow(l, 2) - math.Pow(e, 2)
}

func robotConstraint6(x types.Genotype, zMax float64) float64 {
	a, b, _, e, _, l, _ := x[0], x[1], x[2], x[3], x[4], x[5], x[6]
	return math.Pow(l-zMax, 2) + math.Pow(a-e, 2) - math.Pow(b, 2)
}

func robotConstraint7(x types.Genotype, zMax float64) float64 {
	_, _, _, _, _, l, _ := x[0], x[1], x[2], x[3], x[4], x[5], x[6]
	return l - zMax
}
