package support

import (
	"testing"
/*	"github.com/easy/lego/db/entries"
	"github.com/stellar/go/support/log"*/
)

func TestConnectPostgres(t *testing.T)  {
	var connector PostgresConnector

	connector.Connect()

	entry := Account{}

	str_selector := "select id, details from account where id=1"
	/*connector.Session.Get(&entry, str_selector)*/

	/*str_insert := "insert into account (id, details) VALUES ('1', '{}')";
	connector.Insert(str_insert)*/
	connector.Get(str_selector, &entry)

	connector.Close()
}