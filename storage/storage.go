package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	AddEntry() error
	GetEntries() error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gohost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()

}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists account (
           id serial primary key,
           AverageLoad_1 int,
           AverageLoad_5 int,
           AverageLoad_15 int,
           HostUpTime_Days int,
           HostUpTime_Hours int,
           HostUpTime_Minutes int,
           MemPercentUsed decimal(3,2)
           created_at timestamp
    )`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) AddEntry() error {

	return nil
}
