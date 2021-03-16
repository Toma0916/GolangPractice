package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // PostgreSQLのドライバ、インポートされたときに関数initが実行されて自身を登録
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=Toma dbname=Toma sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{}) // 自動マイグレーションを持っているため、setup.sqlの必要がない
}

func main() {
	post := Post{Content: "Hello, World", Author: "Sau Sheong"}
	fmt.Println(post)
	Db.Create(&post) // 投稿の作成
	fmt.Println(post)

	comment := Comment{Content: "良い投稿だね", Author: "Joe"}    // コメントの追加
	Db.Model(&post).Association("Comments").Append(comment) // 投稿へのコメントの取得、PostIdを直接操作していないことに注目

	var readPost Post
	Db.Where("author = $1", "Sau Sheong").First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])
}
