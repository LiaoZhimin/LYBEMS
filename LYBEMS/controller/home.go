package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	//name := c.Query("name") //没传则为空
	name := c.DefaultQuery("name", "游客")      //Get 传入的参数
	role := c.DefaultQuery("role", "teacher") //Get 传入的参数
	//返回 String
	//c.String(200, "Hello, Geektutu "+name+" "+role)
	c.HTML(http.StatusOK, "index.html", gin.H{"name": name, "role": role})
}

/* p1 router func   */

func p1IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "lzm", "code": "asdqwe"})
}
