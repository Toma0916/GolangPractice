package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("post.csv") // csvファイルの作成
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello, World!", Author: "Sau Sheong"},
		Post{Id: 2, Content: "Hola Mundo", Author: "Pierre"},
		Post{Id: 3, Content: "Bonjour Monde", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush() // バッファにあった全てのデータを確実に書き込むためのwriterのメソッド

	file, err := os.Open("post.csv") // csvファイルの読み込み
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // レコード内に全てのフィールドが揃っていなくても良いことを示す、正の場合は想定しているフィールド数を表す
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
