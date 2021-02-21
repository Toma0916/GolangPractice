package main

import "fmt"

// 構造体はフィールドの集まり
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
