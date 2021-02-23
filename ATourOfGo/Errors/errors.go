package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil { // errorがnilか確認することでエラーをハンドルする
		fmt.Println(err)
	}
}

// error型は組み込みのインターフェース
// type error interface {
// 	Error() string
// }
