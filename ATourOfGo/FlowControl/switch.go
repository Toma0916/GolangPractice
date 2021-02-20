package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on")
	// 選択されたcaseのみ実行
	// breakが要らない
	// 上から下へと評価
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
