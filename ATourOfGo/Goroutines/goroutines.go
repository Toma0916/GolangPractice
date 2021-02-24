package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// go f(x, y ,z)と書くと新しいgoroutinesが実行される
	// f, x, y, xの評価は実行元のgoroutineで行われ、fの実行は新しいgoroutineで実行される
	// goroutineは同じアドレス空間で実行されるため共有メモリへのアクセスは必ず同期する必要がある
	go say("Hello")
	say("Hello")
}
