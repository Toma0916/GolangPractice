package main

import (
	"fmt"
	"io" // データストリームを読むことを表現するio.Readerインターフェースを規定している
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b) // データを与えられたバイトスライスに入れ、入れたバイトのサイズとエラーを返す
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] | %q\n", b[:n])
		if err == io.EOF { // ストリーム終端はio.EOFで返す
			break
		}
	}
}
