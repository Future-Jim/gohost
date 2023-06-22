package userstore

import (
	"database/sql"
	"fmt"

	"github.com/future-jim/gohost/lib/types"
	_ "github.com/lib/pq"
)

// UserStorage acts as in interface on all functions necessary for an API user
type UserStorage interface {
	//accounts
	CreateAccount(*types.Account) error
	DeleteAccount(int) error
	GetAccounts() ([]*types.Account, error)
	GetAccountByID(int) (*types.Account, error)
	GetAccountByNumber(int) (*types.Account, error)
	//metrics queries
	GetMetric(int) (*types.QueryMetrics, error)
	GetMetricsAll() ([]*types.QueryMetrics, error)
	GetMetricsByDate(types.DateTimeQuery) ([]*types.QueryMetrics, error)
}

// PostgresStore is a type that contains the db for userstore
type PostgresStore struct {
	db *sql.DB
}

// NewPostgresStore creates the db connection and checks if it is live
func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gohost sslmode=disable host=gohost-db port=5432"
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

// Init initializes the user account table
func (s *PostgresStore) Init() error {
	return s.createAccountTable()

}

// Creates account table if it does not exist
func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists account (
           id serial primary key,
           first_name varchar(50),
           last_name varchar(50),
           number serial,
           encrypted_password varchar(100), 
           balance serial,
           created_at timestamp
    )`
	_, err := s.db.Exec(query)
	return err
}

// Creates user account entry
func (s *PostgresStore) CreateAccount(acc *types.Account) error {
	query := `insert into account
        (first_name, last_name, number, encrypted_password, balance, created_at)
        values ($1,$2,$3,$4,$5, $6)`
	_, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.EncryptedPassword,
		acc.Balance,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

// Hard deletes user account from account table
func (s *PostgresStore) DeleteAccount(id int) error {
	query := `delete from account where id=$1`
	_, err := s.db.Query(
		query,
		id)
	if err != nil {
		return err
	}

	return nil
}

// Gets a single user account by the user's id
func (s *PostgresStore) GetAccountByID(id int) (*types.Account, error) {
	query := `select * from account where id=$1`
	rows, err := s.db.Query(
		query,
		id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account %d not found", id)
}

// Gets a single user account by the account number
func (s *PostgresStore) GetAccountByNumber(number int) (*types.Account, error) {
	query := `select * from account where number=$1`
	rows, err := s.db.Query(
		query,
		number)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account %d not found", number)
}

// Gets all accounts
func (s *PostgresStore) GetAccounts() ([]*types.Account, error) {
	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}
	accounts := []*types.Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// Gets all metrics from the metrics table
func (s *PostgresStore) GetMetricsAll() ([]*types.QueryMetrics, error) {
	query := `select * from metrics`
	rows, err := s.db.Query(
		query)

	if err != nil {
		return nil, err
	}

	metrics := []*types.QueryMetrics{}
	for rows.Next() {
		metric, err := scanIntoQueryMetric(rows)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}

	return metrics, nil

}

// Gets a single metric from the metrics table based on the metric id
func (s *PostgresStore) GetMetric(id int) (*types.QueryMetrics, error) {
	query := `select * from metrics where id=$1`
	rows, err := s.db.Query(
		query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoQueryMetric(rows)
	}

	return nil, fmt.Errorf("metric %d not found", id)

}

// Gets metrics from the metrics table based on a time.Time range
func (s *PostgresStore) GetMetricsByDate(dates types.DateTimeQuery) ([]*types.QueryMetrics, error) {
	query := `SELECT * FROM metrics WHERE created_at BETWEEN $1 and $2`
	rows, err := s.db.Query(
		query, dates.Start, dates.End)
	if err != nil {
		return nil, err
	}

	metrics := []*types.QueryMetrics{}
	for rows.Next() {
		metric, err := scanIntoQueryMetric(rows)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}
	return metrics, nil
}

// helper function to cast a db row into a type.QueryMetrics type
func scanIntoQueryMetric(rows *sql.Rows) (*types.QueryMetrics, error) {
	metrics := new(types.QueryMetrics)
	err := rows.Scan(
		&metrics.ID,
		&metrics.AL1,
		&metrics.AL5,
		&metrics.AL15,
		&metrics.HUTD,
		&metrics.HUTH,
		&metrics.HUTM,
		&metrics.PMU,
		&metrics.CreatedAt)

	if err != nil {
		return nil, err
	}

	return metrics, nil
}

// helper function to cast a db row into a type.Accounts type
func scanIntoAccount(rows *sql.Rows) (*types.Account, error) {
	account := new(types.Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.EncryptedPassword,
		&account.Balance,
		&account.CreatedAt)
	if err != nil {
		return nil, err
	}
	return account, nil
}
