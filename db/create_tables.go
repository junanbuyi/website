package db

import (
	"errors"
	"fmt"
	"log"
	"test/config"
)

// 创建bbs表
func Create_database() error {
	if config.DB == nil {
		return errors.New("数据库连接失败")
	}
	_, err := config.DB.Exec(`CREATE TABLE IF NOT EXISTS bbs (
        id int NOT NULL AUTO_INCREMENT COMMENT 'primary key',
        Title varchar(255) NOT NULL COMMENT '作者',
        Content varchar(255) NOT NULL
    )`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully")
	return nil
}

//创建users_register表
func Create_users_rigister() error {
	if config.DB == nil {
		return errors.New("数据库连接失败")
	}
	_, err := config.DB.Exec(`CREATE TABLE IF NOT EXISTS registers (
		id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'primary key',
		username VARCHAR(255) NOT NULL COMMENT '用户名',
		email VARCHAR(255) NOT NULL COMMENT '邮箱',
		password VARCHAR(255) NOT NULL COMMENT '密码',
		confirmPassword VARCHAR(255) NOT NULL COMMENT '验证密码'
	);`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully")
	return nil
}

//创建personalemail表，用来管理用户个人信息
func Personalemail() error {
	if config.DB == nil {
		return errors.New("数据库连接失败")
	}
	_, err := config.DB.Exec(`CREATE TABLE IF NOT EXISTS personalemail (
		id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'primary key',
		sign TEXT NOT NULL COMMENT '签名',
		email VARCHAR(255) NOT NULL COMMENT '邮箱'
	);`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully")
	return nil
}
