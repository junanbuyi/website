package routers

import (
	"fmt"
	"log"
	"net/http"
	"test/config"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID              int64  `DB:"id"`
	Username        string `DB:"username"`
	Email           string `DB:"email"`
	Password        string `DB:"password"`
	ConfirmPassword string `DB:"confirmPassword"`
}

var users User

func Register(c *gin.Context) {
	// 从body中获取username, email, password, confirm_password参数
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	if username == "" {
		c.JSON(http.StatusOK, gin.H{
			"error": "用户名为空值",
		})
		fmt.Println("1")
		return
	}
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"error": "邮箱为空值",
		})
		fmt.Println("2")
		return
	}
	if password == "" {
		c.JSON(http.StatusOK, gin.H{
			"error": "密码为空值",
		})
		fmt.Println("3")
		return
	}
	if confirmPassword == "" {
		c.JSON(http.StatusOK, gin.H{
			"error": "验证密码为空值",
		})
		fmt.Println("4")
		return
	}

	users.Email = email
	users.Password = password
	users.Username = username
	users.ConfirmPassword = confirmPassword
	if len(email) < 8 || len(password) < 8 || len(username) < 8 || len(confirmPassword) < 8 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failure",
			"message": "Invalid username or password",
		})
	}
	//表单数据写进数据库
	_, err := config.DB.Exec(`INSERT INTO registers (password, email, username,confirmPassword )VALUES (?,?,?,?);`, password, email, username, confirmPassword)
	if err != nil {
		log.Fatalln(err)
	}
	// 返回响应数据
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})

}
