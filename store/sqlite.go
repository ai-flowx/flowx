package store

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/pkg/errors"

	_ "github.com/mattn/go-sqlite3"
)

const (
	driverName = "sqlite3"

	createTableSql = `CREATE TABLE IF NOT EXISTS long_term_memory (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT,
		meta TEXT,
		agent TEXT
	)`
	deleteTableSql = `DELETE FROM long_term_memory;`
	insertTableSql = `INSERT INTO long_term_memory (text, meta, agent) VALUES (?, ?, ?)`
	selectTableSql = `SELECT text, meta, agent FROM long_term_memory LIMIT ?`
)

type Sqlite struct {
	Path string

	Client *sql.DB
}

func (s *Sqlite) Init(ctx context.Context, _ string) error {
	var err error

	if s.Client, err = sql.Open(driverName, s.Path); err != nil {
		return errors.Wrap(err, "failed to open client\n")
	}

	if _, err = s.Client.Exec(createTableSql); err != nil {
		return errors.Wrap(err, "failed to create table\n")
	}

	return nil
}

func (s *Sqlite) Deinit(_ context.Context) error {
	if s.Client == nil {
		return nil
	}

	if err := s.Client.Close(); err != nil {
		return errors.Wrap(err, "failed to close client\n")
	}

	s.Client = nil

	return nil
}

func (s *Sqlite) Reset(_ context.Context) error {
	if _, err := s.Client.Exec(deleteTableSql); err != nil {
		return errors.Wrap(err, "failed to delete table\n")
	}

	return nil
}

func (s *Sqlite) Save(_ context.Context, text string, meta map[string]interface{}, agent string) error {
	buf, err := json.Marshal(meta)
	if err != nil {
		return errors.Wrap(err, "failed to marshal json\n")
	}

	if _, err = s.Client.Exec(insertTableSql, text, string(buf), agent); err != nil {
		return errors.Wrap(err, "failed to insert table\n")
	}

	return nil
}

func (s *Sqlite) Search(_ context.Context, _ string, limit int32, _ float32) ([]interface{}, error) {
	var ret []interface{}

	rows, err := s.Client.Query(selectTableSql, limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query table\n")
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	for rows.Next() {
		var text, meta, agent string
		if err := rows.Scan(&text, &meta, &agent); err != nil {
			return nil, errors.Wrap(err, "failed to scan row\n")
		}
		var m map[string]interface{}
		if err := json.Unmarshal([]byte(meta), &m); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal json\n")
		}
		b := Collection{
			Meta:    m,
			Context: text,
		}
		ret = append(ret, b)
	}

	return ret, nil
}
