package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router 函数内配置 Url
func Router(r *gin.Engine) {
	r.GET("/", homeHandler)

	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")                  //Post 传入的参数
		password := c.DefaultPostForm("username", "000000") // 可设置默认值
		//返回Json
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	p1 := r.Group("/p1")
	{
		p1.GET("/", p1IndexHandler)
		p1.GET("/Index", p1IndexHandler)
	}
	sign := r.Group("/sign")
	{
		sign.GET("/", loginHandler)
		sign.GET("/in", loginHandler)
		sign.GET("/out", logoutHandler)
	}
}

//
//c.Redirect(http.StatusMovedPermanently, "/index") 重定向
