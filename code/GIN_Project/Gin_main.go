package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Title   string
	Content string
}

func test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "First Page"
	data.Content = "First Gin Project"

	//可以支援許多不同的形態: String JSON XML YAML ProtoBuf Redirect
	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/*")

	server.GET("/", test)
	server.Run(":8222")

}
