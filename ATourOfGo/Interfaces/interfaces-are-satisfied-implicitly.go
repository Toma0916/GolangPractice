package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 型にメソッドを実装していくことによりインタフェースを実装する
// インタフェースを実装することを明示的に宣言する必要はない
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
