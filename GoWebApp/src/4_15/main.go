package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first cookie",
		Value:    "Go Web Programming",
		HttpOnly: false,
		Path:     "/",
	}
	c2 := http.Cookie{
		Name:     "second cookie",
		Value:    "manning Publications Co",
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h, err := r.Cookie("first cookie")
	//r.Header["Cookie"]
	if err != nil {
		fmt.Fprintln(w, "err")
	} else {
		fmt.Fprintln(w, h)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
