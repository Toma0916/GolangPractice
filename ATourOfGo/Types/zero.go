package main

import "fmt"

func main() {
	// 初期化しなければゼロ値が代入される
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
