package support

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/stellar/go/support/log"
)

type Account struct {
	Id 	      int    `db:"id"`
	Details   string `db:"details"`
}

type Connector interface {
	Connect() error
	Close()
}

type PostgresConnector struct {
	connectorString string
	Session         *sqlx.DB
}

func (psql *PostgresConnector) Connect() error {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=test_drive sslmode=disable")
	psql.Session = db
	return err
}

func (psql *PostgresConnector) Close() {
	psql.Session.Close()
}

func (psql *PostgresConnector) Get(selector string, dest interface{}) {
	err := psql.Session.Get(dest, selector)
	if err != nil {
		panic(err.Error())
	}
}

func (psql *PostgresConnector) Insert (insert string) {
	resultTx := psql.Session.MustExec(insert)
	log.Info(resultTx)
}