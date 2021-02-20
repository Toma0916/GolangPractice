package main

import "fmt"

// 関数の２つ以上の引数の型が同じである場合、最後の型を残して省略できる
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(42, add(2, 3)))
}
