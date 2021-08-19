package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", "host=localhost dbname=short_link sslmode=disable")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
