package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMysql() {
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123456", "127.0.0.1", 3306, "transaction_center",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(3)
}

func getDB() *gorm.DB {
	return db
}
