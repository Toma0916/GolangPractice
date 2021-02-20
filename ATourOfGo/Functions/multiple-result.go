package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, _ := swap("hello", "world") // :=は変数を定義して代入することを意味する
	fmt.Println(a)
}
