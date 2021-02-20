package main

import (
	"fmt"
	"math"
)

// 型T、変数vに対してT(v)で型変換可能
// golangでは明示的な型変換が必要
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
