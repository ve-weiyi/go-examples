package db

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var db *gorm.DB
var sqlDB *sql.DB

func Connect(dsn string) {
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//PrepareStmt:            true, // 缓存预编译语句
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Info),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
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
