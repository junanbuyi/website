package routers

import (
	"net/http"
	"test/config"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	// 获取表单数据
	username := c.PostForm("username")
	password := c.PostForm("password")
	var db_usename string
	var db_password string
	err := config.DB.QueryRow("SELECT * FROM users WHERE username = ? AND password = ?", username, password).Scan(&db_usename, &db_password)
	if err != nil {
		// 用户不存在或密码不匹配
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failure",
			"message": "Invalid username or password",
		})
		return
	}

	// 用户名和密码正确
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login successful!",
	})
}
