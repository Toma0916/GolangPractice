package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// indexのみが必要な場合は2つ目の値を省略
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	// アンダーバーを使って値を捨てられる
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
