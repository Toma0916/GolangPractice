package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	// makeによって動的サイズの配列を作成することができる
	// makeはゼロ化された配列を割り当て、その配列を指すスライスを返す
	// makeの3番目の引数にスライスの容量を指定する
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}
