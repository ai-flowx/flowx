package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbName   = "gorm"
	sslMode  = "disable"
	timeZone = "Asia/Shanghai"
)

type Postgres struct {
	Host string
	Port int
	User string
	Pass string

	db *gorm.DB
}

func (p *Postgres) Init(_ context.Context, _ string) error {
	var db *sql.DB
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		p.Host, p.User, p.Pass, dbName, p.Port, sslMode, timeZone)

	if p.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		return errors.Wrap(err, "failed to open database\n")
	}

	if db, err = p.db.DB(); err != nil {
		return errors.Wrap(err, "failed to verify database\n")
	}

	if err = db.Ping(); err != nil {
		return errors.Wrap(err, "failed to ping database\n")
	}

	return nil
}

func (p *Postgres) Deinit(_ context.Context) error {
	p.db = nil

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
