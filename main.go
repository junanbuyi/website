package main

import (
	"net/http"
	"test/db"
	"test/gofires"
	"test/gofires/bbs"
	"test/gofires/user"
	"test/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	//数据库初始化
	db.InitDB()
	//创建数据库
	//db.Create_database()
	//db.Create_users_rigister()
	//db.Personalemail()
	//路由
	router := gin.Default()
	router.POST("/bbs_createartical", routers.Getallartical)

	//加载HTML模块
	// 加载 html 文件
	router.LoadHTMLFiles("/static_rouses/html/homepage.html")
	// js 文件
	router.Static("/static_rouses/js", "./js")
	//css文件
	router.Static("/static_rouses/css", "/static_rouses/css/blog.css")
	// 系统图片文件
	router.Static("/image/", "./image/")
	// GET 请求根路由，返回 HTML 登录页面
	router.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{})
	})
	// POST 请求路由，处理表单数据
	router.POST("/login", routers.Admin)

	//GET请求根目录，返回HTML注册页面
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})

	//
	router.POST("/register", routers.Register)

	//GET请求根目录，返回HTML注册页面
	router.GET("/homepage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", gin.H{})
	})

	router.POST("/firstpage", routers.Homepage)
	// GET 请求根路由，返回 HTML 登录页面
	router.GET("/personalpage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "personalpage.html", gin.H{})
	})

	router.GET("/private", func(c *gin.Context) {
		c.HTML(http.StatusOK, "private.html", gin.H{})
	})
	router.POST("/personalinfo", routers.Private)

	// bbs 专区
	bbsRouter := router.Group("/bbs")
	{
		// 查询所有分类
		bbsRouter.GET("/InquireClassification", bbs.GetClassification)
		// 添加评论
		bbsRouter.POST("/AddComment", bbs.AddComment)
		// 作者
		bbsRouter.POST("/Author", bbs.Author)
		// 新建文章
		bbsRouter.POST("/CreateTexts", bbs.CreateText)
		// 删除文章
		bbsRouter.POST("/DeleteBlog", bbs.DeleteFromBlog)
		// 删除评论
		bbsRouter.POST("/DeleteComment", bbs.DeleteComment)
		// 获取评论 ID
		bbsRouter.POST("/GetCommentsID", bbs.GetCommentsID)
		// 获取用户文章
		bbsRouter.POST("/GetUserText", bbs.GetUserText)
		// 获取头像
		bbsRouter.POST("/GetProfile", bbs.GetProfile)
		// 获取最后一次编辑时间
		bbsRouter.POST("/GetLastModify", bbs.GetLastModify)
		// 获取修改文章信息
		bbsRouter.POST("/GetModifyBlog", bbs.GetModifyBlog)
		// 获取图片 url
		bbsRouter.POST("/Getpicurl", bbs.Getpicurl)
		// 获取页面数量
		bbsRouter.POST("/InquirePageNums", bbs.GetPageNums)
		// 获取文章内容
		bbsRouter.POST("/InquireText", bbs.GetText)
		// 编辑文章
		bbsRouter.POST("/ModifyBlog", bbs.ModifyBlog)
		// 赞
		bbsRouter.POST("/Parise", bbs.Parise)
		// 获取赞数
		bbsRouter.POST("/PariseNum", bbs.PariseNum)
		// 搜索
		bbsRouter.POST("/Search", bbs.Search)
		// 评论
		bbsRouter.POST("/TextComment", bbs.TextComment)
		// 浏览量
		bbsRouter.POST("/Views", bbs.Views)
	}

	// 用户路由
	users := router.Group("/user")
	{
		users.GET("/signAddScore", user.SignAddScore)
		users.GET("/getUsersName", user.GetUsersNames)
		users.GET("/checkPermission", user.CheckPermission)
		users.GET("/IsSystem", user.IsSystem)

		users.POST("/IsSystems", user.IsSystemOrAuthor)
		users.POST("/changePassword", user.ChangePassWord)
		users.POST("/login", user.Login)
		users.POST("/register", user.Register)
		users.POST("/UploadProfile", bbs.UploadProfile)
	}

	// 收藏路由
	collections := router.Group("/collections")
	{
		collections.GET("/", gofires.ToCollections)
	}

	//机器对话
	router.GET("/chargpt", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chatgpt.html", gin.H{})
	})
	//启动服务器
	router.Run(":8080")
}
