package config

import (
	setting "bamboo.com/pipeline/Go-assault-squad/config"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func DBInit(cfg *setting.MySQLConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.DB)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("sql.Open err, \v", err))
	}
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns) //最大连接数
	sqlDB.SetMaxOpenConns(cfg.MaxIdleConns)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //禁用外键生成
	})
	if err != nil {
		panic(fmt.Sprintf("链接数据库失败\v", err))
	}
	DBConn = gormDB
}
