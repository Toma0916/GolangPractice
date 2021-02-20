package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!" // 初期化子がある場合は型を省略することが可能
	fmt.Println(i, j, c, python, java)
}
