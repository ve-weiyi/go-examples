package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var db *gorm.DB
var sqlDB *sql.DB

const dsn = "root:mysql7914@(127.0.0.1:3306)/anker?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() {
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("GORM 数据库连接失败: %v", err)
	}

	sqlDB, err = db.DB()
	if err != nil {
		log.Fatalf("SQL 数据库连接失败: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Printf("Mysql数据库连接成功:%v", dsn)
}

func CloseDB() {
	_ = sqlDB.Close()
}

func DB() *gorm.DB {
	return db
}

func SqlDB() *sql.DB {
	return sqlDB
}
