package main

import "net/http"

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second cookie",
		Value:    "manning Publications Co",
		HttpOnly: true,
	}
	// Setで１番目のクッキーを設定してからAddで２番めのクッキーを追加
	w.Header().Set("Set-Cookie", c1.String()) // メソッドStringでクッキーをシリアル化して返す
	w.Header().Add("Set-Cookie", c2.String())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	server.ListenAndServe()
}
