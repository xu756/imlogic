package db

import (
	"context"
	"imlogic/ent"
	"imlogic/ent/migrate"
	"log"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
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
