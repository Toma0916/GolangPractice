package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 単純な型名であれば省略可能
var m = map[string]Vertex{
	"Bell Labs": {48.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
