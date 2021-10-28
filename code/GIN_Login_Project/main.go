package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func LoginAuth(c *gin.Context) {
	var (
		username   string
		userpasswd string
	)

	// 有無輸入帳號
	if Namein, isExist := c.GetPostForm("username"); isExist && Namein != "" {
		username = Namein
	} else {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": errors.New("must key in user name"),
		})
		return
	}

	// 有無輸入密碼
	if Passin, isExist := c.GetPostForm("password"); isExist && Passin != "" {
		userpasswd = Passin
	} else {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": errors.New("must key in user password"),
		})
	}

	if err := Auth(username, userpasswd); err == nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"success": "登入成功",
		})
	} else {

		c.HTML(http.StatusUnauthorized, "index.html", gin.H{
			"error": err,
		})
	}
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")

	// 靜態檔
	server.Static("/assets", "./template/assets")
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth)
	server.Run(":8999")

}
