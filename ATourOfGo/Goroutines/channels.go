package main

import "fmt"

// channelは並列実行された関数間における値の送受信を提供する型
// 宣言はチャンネルで扱う型とともに行う
// ちなみにgoroutineは非同期処理にするために戻り値を書けない

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x := <-c
	y := <-c // receive from c
	fmt.Println(x, y, x+y)
}
