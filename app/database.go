package app

import (
	"time"
	"userloginapi/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/userloginapi?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	helper.PanicIfError(err)

	sql, err := db.DB()
	helper.PanicIfError(err)

	sql.SetMaxOpenConns(100)
	sql.SetMaxIdleConns(100)
	sql.SetConnMaxLifetime(30 * time.Minute)
	sql.SetConnMaxIdleTime(100)

	return db
}
