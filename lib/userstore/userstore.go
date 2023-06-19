package userstore

import (
	"database/sql"
	"fmt"

	"github.com/future-jim/gohost/lib/types"
	_ "github.com/lib/pq"
)

type UserStorage interface {
	CreateAccount(*types.Account) error
	DeleteAccount(int) error
	UpdateAccount(*types.Account) error
	GetAccounts() ([]*types.Account, error)
	GetAccountByID(int) (*types.Account, error)
	GetAccountByNumber(int) (*types.Account, error)
	GetQuery() error
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

func (s *PostgresStore) GetQuery() error {
	fmt.Println("hi")
	return nil
}

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

func (s *PostgresStore) UpdateAccount(*types.Account) error {
	return nil
}

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
