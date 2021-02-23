package main

import "fmt"

func main() {
	var i interface{} = "hello"
	// 型アサーションはインターフェースを値の基となる具体的な値を利用する手段を提供する
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
