package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	// ハンドラ関数はハンドラのように祓魔う関数
	// メソッドserveHTTPと同じシグネチャを持つ
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello) // 内部でハンドラ関数をハンドラに変換する
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
