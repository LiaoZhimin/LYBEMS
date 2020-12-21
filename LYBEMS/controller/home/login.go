package login

import (
	"fmt"

	"net/http"

	ss "../../services"

	ms "../../models"

	"github.com/gin-gonic/gin"
)

//Routers 传入路由组，往该路由组添加路由项
func Routers(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.POST("/lgion", Function)
	return
}

// Function is a function
func Function(c *gin.Context) {
	var input ms.Input
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
	}
	if output, err := ss.Function(c, input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "获取数据失败",
			"data": output,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "获取数据成功",
			"data": output,
		})
	}
	return
}
