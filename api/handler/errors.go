package handler

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			c.Abort()
		}

	}()

	c.Next()
}
