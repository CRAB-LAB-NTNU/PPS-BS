package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/CRAB-LAB-NTNU/PPS-BS/testSuite"
)

func main() {
	x := generate()
	fmt.Println(testSuite.CMOP1(x))
	fmt.Println(testSuite.CMOP2(x))
	fmt.Println(testSuite.CMOP3(x))
	fmt.Println(testSuite.CMOP4(x))
	fmt.Println(testSuite.CMOP5(x))
	fmt.Println(testSuite.CMOP6(x))
	fmt.Println(testSuite.CMOP7(x))
	fmt.Println(testSuite.CMOP8(x))
	fmt.Println(testSuite.CMOP9(x))
	fmt.Println(testSuite.CMOP10(x))
	fmt.Println(testSuite.CMOP11(x))
	fmt.Println(testSuite.CMOP12(x))
	fmt.Println(testSuite.CMOP13(x))
	fmt.Println(testSuite.CMOP14(x))
}

func generate() []float64 {
	rand.Seed(time.Now().UnixNano())
	arr := make([]float64, 30)
	for i := 0; i < 30; i++ {
		arr[i] = rand.Float64()
	}
	return arr
}
