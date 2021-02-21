package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	// sliceの長さは含まれる要素の数
	// sliceの容量ははスライスの最初の要素から数えて元となる配列の要素数
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}
