package mysql

import (
	"bluebell/models"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

func Init() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bluebell?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	sqlDB, err = db.DB()
	if err != nil {
		return err
	}
	if err = sqlDB.Ping(); err != nil {
		return errors.WithMessage(err, "connect to db failed")
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(50)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(200)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	if err = db.AutoMigrate(&models.User{}); err != nil {
		return errors.WithMessage(err, "autoMigrate failed")
	}

	return nil
}

func Close() {
	_ = sqlDB.Close()
}

func GetDB() *gorm.DB {
	return db
}
