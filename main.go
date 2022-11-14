package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 構造体
type book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Describe string `json:"describe"`
}

// データ作成
var books = []book{
	{
		ID:       1,
		Title:    "タイトルA",
		Describe: "Taro",
	},
	{
		ID:       2,
		Title:    "タイトルB",
		Describe: "Jiro",
	},
	{
		ID:       3,
		Title:    "タイトルC",
		Describe: "Takuya",
	},
}

func main() {
	r := gin.Default()

	// 全てのbookを返すハンドラ
	r.GET("/books", getBooks)
	// bookを追加
	r.POST("/books", postBook)

	r.Run("localhost:8080")
}

// 全てのtodoを返すハンドラ
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBook(c *gin.Context) {
	var newBook book

	// リクエストボディをnewBookにバインド
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// newBookを追加
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
