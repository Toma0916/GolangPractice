package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// 任意の型にメソッドを宣言可能
// レシーバ型が同じパッケージにある必要がある
// →組み込み型には設定出来ない
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
