package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

// ServeHTTPを実装
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler, // ハンドラを作成してサーバに割り当てた、全てのリクエストがこのハンドラで処理される
	}
	server.ListenAndServe()
}
