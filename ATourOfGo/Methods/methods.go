package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Vertexにメソッドを紐付けている
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// Goにはクラスはないが、型にメソッドを定義できる
	// メソッドは特別なレシーバ関数を引数を関数にとる
	// レシーバはfuncキーワードとメソッド名の間にじしんの引数リストで表現する
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
