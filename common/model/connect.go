package model

import (
	"fmt"
	"github.com/xu756/imlogic/common/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var Db *gorm.DB

func InitModel() {
	connect()
	err := create()
	if err != nil {
		panic(err)
	}
}

func connect() *gorm.DB {
	dsn := "host=%s user=%s password=%s dbname=%s port=%d  TimeZone=Asia/Shanghai"
	c := config.RunData.DbConfig
	dsn = fmt.Sprintf(dsn, c.Addr, c.Username, c.Password, c.DbName, c.Port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDb.SetConnMaxLifetime(time.Hour)
	Db = db
	return Db
}

// create 创建数据表
func create() error {

	return nil
}

func GetDb() *gorm.DB {
	db, err := Db.DB()
	if err != nil {
		log.Print(err, "获取数据库连接失败")
		return connect()
	}
	err = db.Ping()
	if err != nil {
		log.Print(err, "ping 数据库连接失败")
		return connect()
	}
	return Db
}
