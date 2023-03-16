package routers

import (
	"fmt"
	"log"
	"net/http"
	"test/config"

	"github.com/gin-gonic/gin"
)

type Personal struct {
	sign  string
	email string
}

func Private(c *gin.Context) {
	// 获取表单信息
	sign := c.PostForm("sign")
	email := c.PostForm("email")
	if sign == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "1",
		})
	}
	if email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "1",
		})
	}
	fmt.Println(sign, email, 1)
	// 将表单数据写入数据库
	_, err := config.DB.Exec(`INSERT INTO personalemail (sign, email) VALUES (?, ?);`, sign, email)
	if err != nil {
		log.Fatalln(err)
	}
	// 返回响应数据
	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}
