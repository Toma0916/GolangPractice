package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	// インターフェース自体の中にある具体的な値がnilの場合
	// メソッドはnilをレシーバとして呼び出される
	// Goではnilをレシーバとして呼び出されても適切に処理するメソッドを記述するのが一般的である
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	// nilを保持するインターフェースの値それ自体は非nil
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
