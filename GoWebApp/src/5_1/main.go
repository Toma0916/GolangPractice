package main

import (
	"net/http"
	"text/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html") // テンプレートファイルを解析
	t.Execute(w, "Hello, World!")            // Executeでデータを当てはめる
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
