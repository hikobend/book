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

	r.GET("/books", getbooks)

	r.Run("localhost:8080")
}

// 全てのtodoを返すハンドラ
func getbooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
