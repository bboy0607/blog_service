package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() (*gorm.DB, error) {
	s := "root:1234@tcp(127.0.0.1:3306)/blog_service?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", s)
	if err != nil {
		return db, err
	}
	return db, nil
}
