package entries

type QueryBuilder interface {

}

const InsertString = "INSERT INTO "

type Account struct {
	ID          int    `db:"id"`
	Details     string `db:"details"`
	Hash 		string `db:"hasj"`
}

type Word struct {
	Word_id        int    `db:"id"`
	Word_value     string `db:"value"`
}

func AddUserWord(account_id, word string) string {
	query := InsertString + "account_word"
	query = query + " (account_id, word_id) "
	query = query + "VALUES('"+account_id+"','"+word+"');"
	return query
}

func AddWord(word string) string{
	query := InsertString + "word"
	query = query + " (value) "
	query = query + "VALUES('"+word+"');"
	return query
}

func ExistWord(word string) string {
	query := "select exists (select null from word where value='"
	query = query + word + "');"
	return query
}

func GetWordID(word string) string {
	query := "select id from word where value='"
	query = query + word +"';"
	return query
}

func GetAllCustomWords(account_id string) string {
	query := "select value from word, account_word where word.id=account_word.word_id AND account_id='"
	query = query + account_id + "';"
	return query
}