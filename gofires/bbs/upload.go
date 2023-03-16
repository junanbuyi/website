package bbs

import (
	"net/http"
	"strconv"
	"test/config"
	"test/gofires/algorithm"
	"test/gofires/user"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 上传文件
func Upload(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())

	result := map[string]interface{}{
		"msg": "success",
	}
	// 获取上传的文件
	res, err := ctx.MultipartForm()
	// 错误处理
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	files := res.File["file"]

	for _, file := range files {
		// 保存文件
		err = ctx.SaveUploadedFile(file, "userFile"+"/"+file.Filename)
		if err != nil {
			result["msg"] = err.Error()
			continue
		}
	}

	ctx.JSON(http.StatusOK, result)
}

// 上传头像
func UploadProfile(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())

	result := map[string]interface{}{
		"msg": "success",
	}
	// 检查登录状态
	if !user.IsLogin(ctx) {
		result["msg"] = "未登录"
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 获取图片
	pic, _ := ctx.FormFile("pic")
	// 连接数据库
	conn, err := sqlx.Connect("mysql", config.MysqlAcount+":"+config.MysqlPassword+"@tcp("+config.MysqlIp+config.MysqlPort+")/user")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer conn.Close()
	// 获取 id
	var id int
	// 获取 account
	account, err := user.GetUserEmail(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	conn.Get(&id, "SELECT id FROM user WHERE account = ?", account)
	// 保存文件
	go func() {
		ctx.SaveUploadedFile(pic, `bbsFile/`+strconv.Itoa(id)+`/`+pic.Filename)
	}()

	conn.Exec("UPDATE user SET pic = ? WHERE ID = ?;", config.Addr+`bbsFile/`+strconv.Itoa(id)+`/`+pic.Filename, id)
	ctx.JSON(http.StatusOK, result)
}
