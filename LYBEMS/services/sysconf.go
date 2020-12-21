package sysconf

import (
	ms "../models"

	"github.com/gin-gonic/gin"
)

// GetConfig is a func to link database
func GetConfig(c *gin.Context, input ms.I) (output ms.Output, err error) {
	return output, nil
}
