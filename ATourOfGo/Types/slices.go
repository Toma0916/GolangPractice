package main

import "fmt"

func main() {
	primes := [6]int{1, 2, 3, 4, 5, 6}
	// []TはTのスライスを表す、可変長
	var s []int = primes[1:4]
	fmt.Println(s)
}
