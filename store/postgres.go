package store

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Host string
	Port int
}

func (p *Postgres) Init(_ context.Context, name string) error {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	_, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return nil
}

func (p *Postgres) Deinit(_ context.Context) error {
	return nil
}

func (p *Postgres) Reset(_ context.Context) error {
	return nil
}

func (p *Postgres) Save(_ context.Context, text string, meta map[string]interface{}, agent string) error {
	return nil
}

func (p *Postgres) Search(_ context.Context, query string, limit int32, threshold float32) ([]interface{}, error) {
	return nil, nil
}
