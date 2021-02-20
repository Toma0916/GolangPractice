package main

import "fmt"

func main() {
	sum := 1
	// 初期化と後処理ステートメントは省略可能
	// 所謂while文となる
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
