package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "OK", "code": "sign in"})
}

func logoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "OK", "code": "sign out"})
}
