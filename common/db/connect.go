package db

import (
	"context"
	"fmt"
	"imlogic/common/cache"
	"imlogic/common/config"
	"imlogic/ent"
	"imlogic/ent/migrate"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type customModel struct {
	client *ent.Client
	cache  *cache.Client
}

func NewModel() Model {
	dsn := "host=%s user=%s password=%s dbname=%s port=%d  TimeZone=Asia/Shanghai sslmode=disable"
	dsn = fmt.Sprintf(dsn, config.RunData().DbConfig.Addr, config.RunData().DbConfig.Username, config.RunData().DbConfig.Password, config.RunData().DbConfig.DbName, config.RunData().DbConfig.Port)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)          // 设置最大打开连接数
	db.SetConnMaxLifetime(time.Hour) // 设置连接最大存活时间

	drive := sql.OpenDB(dialect.Postgres, db.DB)
	client := ent.NewClient(ent.Driver(drive))
	err = CreateTable(client)
	if err != nil {
		log.Panic(err)

	}
	return &customModel{
		client: client,
		cache:  cache.NewClient(),
	}
}

func CreateTable(client *ent.Client) error {
	err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropIndex(true),
	)
	return err
}
