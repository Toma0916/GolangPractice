package main

import "fmt"

func main() {
	var i interface{} // 空のインターフェースで任意の型を保持出来る、未知の型を扱うコードで使用される
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
