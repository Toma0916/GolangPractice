package main

import "fmt"

// アプリケーションのフォーラムの投稿を表現
type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post           // ユニークなidを投稿へのポインタへ紐付け
var PostsByAuthor map[string][]*Post // 著者名を投稿へのポインタのスライスに紐付け

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello, World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "Hola Mundo", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Bonjour Monde", Author: "Pedro"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings", Author: "Sau Sheong"}
	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
