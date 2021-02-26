package main

import (
	"net/http"
)

// 最も単純なウェブサーバ
func main() {
	http.ListenAndServe("", nil)
}
