package main

import "fmt"

// 予め返り値の変数に名前を付けておくことが可能
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked returnという
}

func main() {
	fmt.Println(split(17))
}
