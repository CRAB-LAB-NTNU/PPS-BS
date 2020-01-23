package main

import (
	"fmt"

	"github.com/CRAB-LAB-NTNU/PPS-BS/testSuite"
)

func main() {
	x := []float64{0.5, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 0.8, 0.9, 0.7}
	fmt.Println(testSuite.Cmop2(x))
}
