package main

import "net/http"

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	// cert.pemがSSL証明書
	// key.pemがサーバ用の秘密鍵
	// 実運用の場合は認証局から証明書を取得する
	// 試して見るだけなら自分で証明書を作成出来る
	server.ListenAndServeTLS("cert.pen", "key.pem")
}
