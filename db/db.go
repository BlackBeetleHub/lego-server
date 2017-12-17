package db

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
	Get(selector string, dest interface{})
	Select(selector string, dest interface{})
	Insert(inset string)
}

type PostgresConnector struct {
	ConnectorString string
	session         *sqlx.DB
}

func (psql *PostgresConnector) Connect() error {
	db, err := sqlx.Connect("postgres", psql.ConnectorString)
	psql.session = db
	return err
}

func (psql *PostgresConnector) Close() {
	psql.session.Close()
}

func (psql *PostgresConnector) Get(selector string, dest interface{}) {
	err := psql.session.Get(dest, selector)
	if err != nil {
		panic(err.Error())
	}
}

func (psql *PostgresConnector) Select(selector string, dest interface{}) {
	err := psql.session.Select(dest, selector)
	if err != nil {
		panic(err.Error())
	}
}

func (psql *PostgresConnector) Insert (insert string) {
	resultTx := psql.session.MustExec(insert)
	log.Info(resultTx)
}