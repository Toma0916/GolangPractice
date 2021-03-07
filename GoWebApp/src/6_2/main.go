package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello World!\n")
	err := ioutil.WriteFile("data1", data, 0644) // WriteFileでファイルに書き込む、0644はパーミッション
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1") // ReadFileでファイルを読み込む
	fmt.Print(string(read1))

	file1, _ := os.Create("data2") // 構造体Fileを使ってファイルを読み書き
	defer file1.Close()            // ファイルのクローズを忘れないようにdeferを使ってクローズしておく

	bytes, _ := file1.Write(data) // Writeで書き込み可能
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
