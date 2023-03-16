package config

import (
	"github.com/jmoiron/sqlx"
)

//mysql监听端口
var MysqlPort = ":3306"

//mysql ip
var MysqlIp = "localhost"

//mysql账户
var MysqlAcount = "root"

//mysql密码
var MysqlPassword = "753852941qx@"

// 连接信息
var MySQLInfo = MysqlAcount + ":" + MysqlPassword + "@tcp(" + MysqlIp + MysqlPort + ")/"

//是否允许注册
var AllowRegister = true

//发送方邮箱
var SenderAccount = "1327913121@qq.com"

//认证码
var SenderPassword = " sd"

// Redis ip
var RedisIp = "localhost"

// Redis 端口
var RedisPort = ":6379"

// info
var RedisInfo = RedisIp + RedisPort

// 域名
var Addr = "http://127.0.0.1/"

//数据库连接
var DB *sqlx.DB

// 管理员邮箱
var SystemAccount = "YourEmail@outlook.com"
