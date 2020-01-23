package testSuite

import "math"

func CMOP1(x []float64) Fitness {
	fitness := Fitness{}
	fitness.InequalityConstraints = make(map[string]bool)
	fitness.Objectives = make(map[string]float64)

	fitness.Objectives["F1"] = objective1(x)
	fitness.Objectives["F2"] = objective2(x)

	fitness.InequalityConstraints["C1"] = constraint1(x, inner1)
	fitness.InequalityConstraints["C2"] = constraint1(x, inner2)

	return fitness
}

func CMOP2(x []float64) Fitness {
	fitness := Fitness{}
	fitness.InequalityConstraints = make(map[string]bool)
	fitness.Objectives = make(map[string]float64)

	fitness.Objectives["F1"] = objective1(x)
	fitness.Objectives["F2"] = objective3(x)

	fitness.InequalityConstraints["C1"] = constraint1(x, inner1)
	fitness.InequalityConstraints["C2"] = constraint1(x, inner2)

	return fitness
}

func CMOP3(x []float64) Fitness {
	fitness := CMOP1(x)

	fitness.InequalityConstraints["C3"] = constraint2(x)

	return fitness
}

func CMOP4(x []float64) Fitness {
	fitness := CMOP2(x)

	fitness.InequalityConstraints["C3"] = constraint2(x)

	return fitness
}

func CMOP5(x []float64) Fitness {
	fitness := Fitness{}

	fitness.InequalityConstraints = make(map[string]bool)
	fitness.Objectives = make(map[string]float64)

	fitness.Objectives["F1"] = objective4(x)
	fitness.Objectives["F2"] = objective5(x)

	p := []float64{1.6, 2.5}
	q := []float64{1.6, 2.5}
	a := []float64{2, 2}
	b := []float64{4, 8}

	fitness.InequalityConstraints["C1"] = constraint3(x, p, q, a, b, 0, fitness.Objectives["F1"], fitness.Objectives["F2"])
	fitness.InequalityConstraints["C2"] = constraint3(x, p, q, a, b, 1, fitness.Objectives["F1"], fitness.Objectives["F2"])

	return fitness
}

func CMOP6(x []float64) Fitness {
	fitness := Fitness{}

	fitness.InequalityConstraints = make(map[string]bool)
	fitness.Objectives = make(map[string]float64)

	fitness.Objectives["F1"] = objective4(x)
	fitness.Objectives["F2"] = objective6(x)

	p := []float64{1.8, 2.8}
	q := []float64{1.8, 2.8}
	a := []float64{2, 2}
	b := []float64{8, 8}

	fitness.InequalityConstraints["C1"] = constraint3(x, p, q, a, b, 0, fitness.Objectives["F1"], fitness.Objectives["F2"])
	fitness.InequalityConstraints["C2"] = constraint3(x, p, q, a, b, 1, fitness.Objectives["F1"], fitness.Objectives["F2"])

	return fitness
}

func CMOP7(x []float64) Fitness {
	fitness := Fitness{}

	fitness.InequalityConstraints = make(map[string]bool)
	fitness.Objectives = make(map[string]float64)

	fitness.Objectives["F1"] = objective4(x)
	fitness.Objectives["F2"] = objective5(x)

	p := []float64{1.2, 2.25, 3.5}
	q := []float64{1.2, 2.25, 3.5}
	a := []float64{2, 2.5, 2.5}
	b := []float64{6, 12, 10}

	fitness.InequalityConstraints["C1"] = constraint3(x, p, q, a, b, 0, fitness.Objectives["F1"], fitness.Objectives["F2"])
	fitness.InequalityConstraints["C2"] = constraint3(x, p, q, a, b, 1, fitness.Objectives["F1"], fitness.Objectives["F2"])
	fitness.InequalityConstraints["C3"] = constraint3(x, p, q, a, b, 2, fitness.Objectives["F1"], fitness.Objectives["F2"])

	return fitness
}

func CMOP8(x []float64) Fitness {
	fitness := Fitness{}

	fitness.InequalityConstraints = make(map[string]bool)
	fitness.Objectives = make(map[string]float64)

	fitness.Objectives["F1"] = objective4(x)
	fitness.Objectives["F2"] = objective6(x)

	p := []float64{1.2, 2.25, 3.5}
	q := []float64{1.2, 2.25, 3.5}
	a := []float64{2, 2.5, 2.5}
	b := []float64{6, 12, 10}

	fitness.InequalityConstraints["C1"] = constraint3(x, p, q, a, b, 0, fitness.Objectives["F1"], fitness.Objectives["F2"])
	fitness.InequalityConstraints["C2"] = constraint3(x, p, q, a, b, 1, fitness.Objectives["F1"], fitness.Objectives["F2"])
	fitness.InequalityConstraints["C3"] = constraint3(x, p, q, a, b, 2, fitness.Objectives["F1"], fitness.Objectives["F2"])

	return fitness
}

func CMOP9(x []float64) Fitness {
	fitness := Fitness{}

	fitness.InequalityConstraints = make(map[string]bool)
	fitness.Objectives = make(map[string]float64)

	fitness.Objectives["F1"] = objective7(x)
	fitness.Objectives["F2"] = objective8(x)

	p := []float64{1.4}
	q := []float64{1.4}
	a := []float64{1.5}
	b := []float64{6.0}

	fitness.InequalityConstraints["C1"] = constraint3(x, p, q, a, b, 0, fitness.Objectives["F1"], fitness.Objectives["F2"])
	fitness.InequalityConstraints["C2"] = constraint4(x, fitness.Objectives["F1"], fitness.Objectives["F2"], 2)

	return fitness
}

func CMOP10(x []float64) Fitness {
	fitness := Fitness{}

	fitness.InequalityConstraints = make(map[string]bool)
	fitness.Objectives = make(map[string]float64)

	fitness.Objectives["F1"] = objective7(x)
	fitness.Objectives["F2"] = objective9(x)

	p := []float64{1.1}
	q := []float64{1.2}
	a := []float64{2.0}
	b := []float64{4.0}

	fitness.InequalityConstraints["C1"] = constraint3(x, p, q, a, b, 0, fitness.Objectives["F1"], fitness.Objectives["F2"])
	fitness.InequalityConstraints["C2"] = constraint4(x, fitness.Objectives["F1"], fitness.Objectives["F2"], 1)

	return fitness
}

func CMOP11(x []float64) Fitness {
	fitness := Fitness{}

	fitness.InequalityConstraints = make(map[string]bool)
	fitness.Objectives = make(map[string]float64)

	fitness.Objectives["F1"] = objective7(x)
	fitness.Objectives["F2"] = objective9(x)

	p := []float64{1.2}
	q := []float64{1.2}
	a := []float64{1.5}
	b := []float64{5.0}

	fitness.InequalityConstraints["C1"] = constraint3(x, p, q, a, b, 0, fitness.Objectives["F1"], fitness.Objectives["F2"])
	fitness.InequalityConstraints["C2"] = constraint4(x, fitness.Objectives["F1"], fitness.Objectives["F2"], 2.1)

	return fitness
}

func CMOP12(x []float64) Fitness {
	fitness := Fitness{}

	fitness.InequalityConstraints = make(map[string]bool)
	fitness.Objectives = make(map[string]float64)

	fitness.Objectives["F1"] = objective7(x)
	fitness.Objectives["F2"] = objective8(x)

	p := []float64{1.6}
	q := []float64{1.6}
	a := []float64{1.5}
	b := []float64{6.0}

	fitness.InequalityConstraints["C1"] = constraint3(x, p, q, a, b, 0, fitness.Objectives["F1"], fitness.Objectives["F2"])
	fitness.InequalityConstraints["C2"] = constraint4(x, fitness.Objectives["F1"], fitness.Objectives["F2"], 2.5)

	return fitness
}

func CMOP13(x []float64) Fitness {
	fitness := Fitness{}

	fitness.SoftConstraints = make(map[string]float64)
	fitness.Objectives = make(map[string]float64)

	fitness.Objectives["F1"] = objective10(x)
	fitness.Objectives["F2"] = objective11(x)
	fitness.Objectives["F3"] = objective12(x)

	g := math.Pow(fitness.Objectives["F1"], 2) + math.Pow(fitness.Objectives["F2"], 2) + math.Pow(fitness.Objectives["F3"], 2)

	fitness.SoftConstraints["C1"] = constraint5(x, g, 9, 4)
	fitness.SoftConstraints["C2"] = constraint5(x, g, 3.61, 3.24)

	return fitness
}

func CMOP14(x []float64) Fitness {
	fitness := Fitness{}

	fitness.SoftConstraints = make(map[string]float64)
	fitness.Objectives = make(map[string]float64)

	fitness.Objectives["F1"] = objective10(x)
	fitness.Objectives["F2"] = objective11(x)
	fitness.Objectives["F3"] = objective12(x)

	g := math.Pow(fitness.Objectives["F1"], 2) + math.Pow(fitness.Objectives["F2"], 2) + math.Pow(fitness.Objectives["F3"], 2)

	fitness.SoftConstraints["C1"] = constraint5(x, g, 9, 4)
	fitness.SoftConstraints["C2"] = constraint5(x, g, 3.61, 3.24)
	fitness.SoftConstraints["C3"] = constraint5(x, g, 3.0625, 2.56)

	return fitness
}
