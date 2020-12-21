package controller

import (
	"fmt"

	lgn "../controller/login"

	"github.com/gin-gonic/gin"
)

//GinRouter is a func
func GinRouter(r *gin.Engine) *gin.Engine {
	rr := r.Group("/")
	rr.GET("/first", func(c *gin.Context) {
		fmt.Println("first .........")
	})
	rr = r.Group("/a")
	lgn.Routers(rr)
	return r
}
