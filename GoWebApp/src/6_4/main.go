package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func store(data interface{}, filename string) {
	// 空のinterfaceを受け取る -> つまり、なんでも受け取れる
	buffer := new(bytes.Buffer)       // bytes.Bufferはバイトデータの可変バッファ
	encoder := gob.NewEncoder(buffer) // gobエンコーダを生成
	err := encoder.Encode(data)       // Postをバッファ内にエンコード
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600) // バッファから書き込み
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer) // バッファからデコーダを生成
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, Content: "Hello World", Author: "Sau Sheong"}
	store(post, "post1")
	var postRead Post
	load(&postRead, "post1")
	fmt.Println(postRead)
}
