package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < int(math.Pow(10, 6)); i++ {
		new_z := z - (z*z-x)/(2*z)
		if math.Abs(new_z-z) < math.Pow(10, -4) {
			return new_z
		} else {
			z = new_z
		}
	}
	return z
}

func main() {
	fmt.Println("my sqrt", Sqrt(2))
	fmt.Println("math package sqrt", math.Sqrt(2))
}
