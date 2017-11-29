package migration

import (
	"github.com/jmoiron/sqlx"
	"log"
)

var schema = `
CREATE TABLE account (
    id      SERIAL  NOT NULL,
    details TEXT    NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE account_words (
    account_id   int,
    word_id      int
);

CREATE TABLE service (
    id           SERIAL   NOT NULL,
    name         TEXT     NOT NULL,
	token        TEXT     NOT NULL
);

CREATE TABLE account_service (
    account_id int NOT NULL,
    service_id int NOT NULL
);

CREATE TABLE word (
    id      SERIAL NOT NULL ,
    value   text   NOT NULL,
    PRIMARY KEY (ID)
)`

func MigrateDB(connectStr string){
	db, err := sqlx.Connect("postgres", connectStr)
	if err != nil {
		log.Fatalln("Cannot migrate: " + err.Error())
	}
	resultMigration := db.MustExec(schema)
	log.Println(resultMigration)
	db.Close()
}