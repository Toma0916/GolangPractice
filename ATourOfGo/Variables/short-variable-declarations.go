package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// 関数内では:=による暗黙的な型宣言が可能
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
