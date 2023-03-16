package routers

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"test/config"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	// 获取表单数据
	username := c.PostForm("account")
	password := c.PostForm("password")
	fmt.Println(1)
	// 验证用户名和密码是否为空
	fmt.Println(username)
	fmt.Println(password)
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": "Username is required",
		})
		return
	}
	fmt.Println(2)
	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": "Password is required",
		})
		return
	}
	fmt.Println(3)
	// 验证用户名和密码是否合法
	match, err := regexp.MatchString("^[a-zA-Z0-9_]+$", username)
	if err != nil || !match {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": "Invalid username",
		})
		return
	}
	fmt.Println(4)
	match, err = regexp.MatchString("^[a-zA-Z0-9_@#$%^&+=]+$", password)
	if err != nil || !match {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": "Invalid password",
		})
		return
	}

	var dbUsername string
	var dbHashedPassword []byte
	err = config.DB.QueryRow("SELECT username, password FROM registers WHERE username = ?", username).Scan(&dbUsername, &dbHashedPassword)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failure",
				"message": "Invalid username or password",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failure",
				"message": "Failed to query database",
			})
		}
		return
	}
	fmt.Println(8)
	// 用户名和密码正确
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login successful!",
	})

}
