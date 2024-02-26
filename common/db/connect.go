package db

import (
	"context"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/xu756/imlogic/ent"
	"github.com/xu756/imlogic/ent/migrate"
	"log"
	"time"
)

type customModel struct {
	client *ent.Client
}

func NewModel() Model {
	dsn := "host=%s user=%s password=%s dbname=%s port=%d  TimeZone=Asia/Shanghai sslmode=disable"
	drv, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panic(err)
	}
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return &customModel{
		client: ent.NewClient(ent.Driver(drv)),
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
