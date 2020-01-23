package testSuite

import "math"

func constraint1(x []float64, g func([]float64) float64) bool {
	return (0.51-g(x))*(g(x)-0.5) >= 0
}

func constraint2(x []float64) bool {
	return math.Sin(20*math.Pi*x[0])-0.5 >= 0
}

func constraint3(x, p, q, a, b []float64, k int, f1, f2 float64) bool {
	d := -0.25 * math.Pi
	z := math.Pow((f1-p[k])*math.Cos(d)-(f2-q[k])*math.Sin(d), 2) / math.Pow(a[k], 2)
	y := math.Pow((f1-p[k])*math.Sin(d)+(f2-q[k])*math.Cos(d), 2) / math.Pow(b[k], 2)
	return z+y >= 0.1
}

func constraint4(x []float64, f1, f2, z float64) bool {
	d := 0.25 * math.Pi
	a := f1*math.Sin(d) + f2*math.Cos(d)
	b := math.Sin(4*math.Pi*(f1*math.Cos(d)-f2*math.Sin(d))) - z
	return a-b >= 0
}

func constraint5(x []float64, g, a, b float64) float64 {
	return (g - a) * (g - b)
}
