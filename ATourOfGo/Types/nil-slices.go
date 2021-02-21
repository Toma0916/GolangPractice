package main

import "fmt"

func main() {
	// sliceのゼロ値はnil
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
