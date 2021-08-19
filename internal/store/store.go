package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
	LinkRepository *LinkRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres",
		"postgres://postgres:postgres@localhost:5432/link?sslmode=disable")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS link (" +
			"url varchar not null unique," +
			"shortlink varchar not null unique" +
			");")
	if err != nil {
		db.Close()
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Link() *LinkRepository {
	if s.LinkRepository == nil {
		s.LinkRepository = &LinkRepository{
			store: s,
		}
	}
	return s.LinkRepository
}