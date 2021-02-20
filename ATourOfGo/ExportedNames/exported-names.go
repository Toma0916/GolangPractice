package main

import (
	"fmt"
	"math"
)

// 大文字で始まる名前は外部のパッケージから参照できるエクスポート（公開）された名前
// 小文字で始まる名前は公開されていない名前となる
func main() {
	fmt.Println(math.Pi)
}
