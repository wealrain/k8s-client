package handler

import (
	"k8s-client/pkg/cluster"

	"github.com/gin-gonic/gin"
)

func HandleAddUser(c *gin.Context) {
	// 从json中获取数据
	var user cluster.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	err = cluster.AddUser(user)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func HandleDeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := cluster.DeleteUser(id)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func HandleUpdateUser(c *gin.Context) {
	var user cluster.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	err = cluster.UpdateUser(user)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func HandleListUser(c *gin.Context) {
	users, err := cluster.GetUserList()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, users)
}

func HandleLogin(c *gin.Context) {
	var user cluster.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}

	token, err := cluster.CheckUser(user.Username, user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if token {
		c.JSON(200, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(400, gin.H{
			"message": "用户名或密码错误",
		})
	}
}
