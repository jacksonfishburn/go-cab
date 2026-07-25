package db

import (
	"database/sql"
	"errors"

	_ "modernc.org/sqlite"

	"github.com/jacksonfishburn/go-cab/internal/file"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewStore(db *sql.DB) SQLiteStore {
	return SQLiteStore{
		db: db,
	}
}

func Open(path string) (*SQLiteStore, error) {
    sqlDB, err := sql.Open("sqlite", path)
    if err != nil {
        return nil, err
    }
    if err := sqlDB.Ping(); err != nil {
        return nil, err
    }
    if err := createTables(sqlDB); err != nil {
        return nil, err
    }
    return &SQLiteStore{db: sqlDB}, nil
}

func createTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS records (
			name       TEXT PRIMARY KEY,
			size       INTEGER NOT NULL,
			md5        TEXT NOT NULL,
			created_at TEXT NOT NULL,
			updated_at TEXT NOT NULL
		)
	`)
	return err
}

func (s *SQLiteStore) Close() error {
    return s.db.Close()
}

func (s SQLiteStore) Put(name string, record file.Record) error {

	var exists int
	err := s.db.QueryRow(`SELECT 1 FROM records WHERE name = ?`, name).Scan(&exists)

	if err == nil {
		return errors.New("Name Already Taken")
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	_, err = s.db.Exec(
		`INSERT INTO records (name, size, md5, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`,
		name, record.Size, record.MD5, record.CreatedAt, record.UpdatedAt,
	)

	return err
}

func (s SQLiteStore) Get(name string) (file.Record, error) {
	var r file.Record
	err := s.db.QueryRow(
		`SELECT name, size, md5, created_at, updated_at FROM records WHERE name = ?`,
		name,
	).Scan(&r.Name, &r.Size, &r.MD5, &r.CreatedAt, &r.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return file.Record{}, errors.New("No such Record")
	}
	if err != nil {
		return file.Record{}, err
	}
	return r, nil
}

func (s SQLiteStore) List() (map[string]file.Record, error) {
	rows, err := s.db.Query(`SELECT name, size, md5, created_at, updated_at FROM records`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := make(map[string]file.Record)
	for rows.Next() {
		var r file.Record
		if err := rows.Scan(&r.Name, &r.Size, &r.MD5, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, err
		}
		list[r.Name] = r
	}
	return list, rows.Err()
}

func (s SQLiteStore) Delete(name string) error {
	res, err := s.db.Exec(`DELETE FROM records WHERE name = ?`, name)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("No such Record")
	}
	return nil
}