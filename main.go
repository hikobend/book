package main

import (
	"net/http"
	"strconv"

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
	// 指定したbookを表示する
	r.GET("/books/:id", getBookById)
	// bookを更新するハンドラ
	r.PATCH("/books/:id", patchBook)

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

// 指定したbookを表示する
func getBookById(c *gin.Context) {
	// URLからidを取得
	// strconv.Atoi()を利用して、idを文字列から数値に変換
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	for _, t := range books {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}

// 更新するハンドラ
func patchBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	var book book
	book.ID = id
	// JSONのリクエストボディをバインド
	if err = c.BindJSON(&book); err != nil {
		return
	}

	for i, t := range books {
		if t.ID == id {
			books[i] = book
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
