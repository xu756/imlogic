package db

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/glebarez/sqlite"
	"github.com/xu756/imlogic/common/model"
	"gorm.io/gorm"
	"log"
	"time"
)

type customModel struct {
	Db *gorm.DB
}

func NewModel() Model {

	// FIXME: 先使用sqlite3
	//dsn := "host=%s user=%s password=%s dbname=%s port=%d  TimeZone=Asia/Shanghai"
	//c := config.RunData.DbConfig
	//dsn = fmt.Sprintf(dsn, c.Addr, c.Username, c.Password, c.DbName, c.Port)
	//db, err := gorm.Open(postgres.New(postgres.Config{
	//	DSN:                  dsn,
	//	PreferSimpleProtocol: true,
	//}), &gorm.Config{
	//	SkipDefaultTransaction: true,
	//})
	db, err := gorm.Open(sqlite.Open("./build/data.db"), &gorm.Config{
		TranslateError:         true, //  翻译错误
		SkipDefaultTransaction: true, //禁用默认事务
		PrepareStmt:            true, // 缓存预编译语句
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
	sqlDb.SetConnMaxLifetime(time.Second)
	log.Print("【 数据库连接成功 】")
	err = CreateTable(db)
	if err != nil {
		panic(err)
	}
	return &customModel{
		Db: db,
	}
}

func CreateTable(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		klog.Debugf("【 创建表失败 %s 】 ", "user")
		return err
	}
	return nil
}

func TestModel() Model {
	db, err := gorm.Open(sqlite.Open("../../build/data.db"), &gorm.Config{
		TranslateError:         true, //  翻译错误
		SkipDefaultTransaction: true, //禁用默认事务
		PrepareStmt:            true, // 缓存预编译语句
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
	sqlDb.SetConnMaxLifetime(time.Second)
	log.Print("【 数据库连接成功 】")
	err = CreateTable(db)
	if err != nil {
		panic(err)
	}
	return &customModel{
		Db: db,
	}
}
