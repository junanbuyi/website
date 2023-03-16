package db

import (
	"test/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDB() string {
	var err error
	config.DB, err = sqlx.Connect("mysql", config.MySQLInfo+"huajiaofen")
	if err != nil {
		return err.Error()
	}
	return ""
}
