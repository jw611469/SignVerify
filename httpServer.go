package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"signverify/userTABLE"
)

func AUTH(name string, pass string) bool {
	if buff, exist := userTABLE.GETDATA(name); exist && (buff == pass) {
		return true
	} else {
		return false
	}
}
func loginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
func loginAuth(c *gin.Context) {
	var user string
	var pass string
	if buff, exist := c.GetPostForm("username"); exist && buff != "" {
		user = buff
	} else {
		c.HTML(http.StatusBadRequest, "login.html", nil)
	}
	if buff, exist := c.GetPostForm("password"); exist && buff != "" {
		pass = buff
	} else {
		c.HTML(http.StatusBadRequest, "login.html", nil)
	}
	if AUTH(user, pass) {
		c.HTML(http.StatusOK, "login.html", gin.H{"success": "success"})
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{"error": "error"})
	}
}
func main() {
	fmt.Println(userTABLE.GETDATA("test"))
	router := gin.Default()
	router.LoadHTMLGlob("./temp/*")
	router.GET("/login", loginPage)
	router.GET("/", loginPage)
	router.POST("/login", loginAuth)
	router.Run(":8080")
}
