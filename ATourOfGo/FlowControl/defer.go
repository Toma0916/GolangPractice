package main

import "fmt"

func main() {
	// deferへ渡した関数の実行を呼び出し元の関数の終わりまで遅延する
	// 関数の引数の評価自体は即座に行う
	defer fmt.Println("world")
	fmt.Println("hello")
}
