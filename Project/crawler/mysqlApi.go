package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DBConn *gorm.DB

func init() {
	var err error
	dsn := "root:SNH48group!@tcp(127.0.0.1:3306)/douban?charset=utf8mb4&parseTime=True&loc=Local"
	DBConn, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Errorf("Connect to DB failed: %v", err)
	}

	DBConn.AutoMigrate(&CommentMysql{}, &TopicMysql{})
}

func InsertTopicData(data *TopicMysql) {

	if err := DBConn.Create(data).Error; err != nil {
		fmt.Printf("Insert Topic failed: %v; err: %v", data, err.Error())
		return
	}
}

func InsertCommentData(data *CommentMysql) {
	if err := DBConn.Create(data).Error; err != nil {
		fmt.Printf("Insert Comment failed: %v; err: %v", data, err.Error())
		return
	}
	defer DBConn.Close()
}

// func DeleteData(data string)
