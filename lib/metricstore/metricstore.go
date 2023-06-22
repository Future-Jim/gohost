package metricstore

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/future-jim/gohost/lib/types"
	_ "github.com/lib/pq"
)

type MetricStorage interface {
	createMetricTable() error
	AddEntry(*types.Metrics) error
	GetEntry() error
}

type PostgresStore struct {
	db *sql.DB
}

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

func (s *PostgresStore) Init() error {
	return s.createMetricTable()

}

func (s *PostgresStore) createMetricTable() error {
	query := `create table if not exists metrics (
           id serial primary key,
           AverageLoad_1 decimal(3,2),
           AverageLoad_5 decimal(3,2),
           AverageLoad_15 decimal(3,2),
           HostUpTime_Days int,
           HostUpTime_Hours int,
           HostUpTime_Minutes int,
           MemPercentUsed int,
           Created_At timestamp
           );`
	_, err := s.db.Exec(query)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func (s *PostgresStore) AddEntry(metrics *types.Metrics) error {
	query := `insert into metrics
            (AverageLoad_1,
             AverageLoad_5,
             AverageLoad_15,
             HostUpTime_Days,
             HostUpTime_Hours,
             HostUpTime_Minutes,
             MemPercentUsed,
             Created_At)
             values ($1, $2, $3, $4, $5, $6, $7, $8);`
	rows, err := s.db.Query(query,
		metrics.AL.One,
		metrics.AL.Five,
		metrics.AL.Fifteen,
		metrics.HUT.Days,
		metrics.HUT.Hours,
		metrics.HUT.Minutes,
		metrics.PMU.PMU,
		time.Now().UTC())

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer rows.Close()

	return nil
}

func (s *PostgresStore) GetEntry() error {
	query := `select * from metrics;`
	rows, err := s.db.Query(query)
	//todo - load the data into a metrics struct
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()

	return nil

}
