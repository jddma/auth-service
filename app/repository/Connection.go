package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func openConnection() (*gorm.DB, error) {
	URI := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("db_user"),
		os.Getenv("db_pwd"),
		os.Getenv("db_host"),
		os.Getenv("db_name"),
	)
	db, err := gorm.Open(mysql.Open(URI), &gorm.Config{})
	return db, err
}
