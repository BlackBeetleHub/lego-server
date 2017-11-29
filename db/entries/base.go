package entries

type Account struct {
	ID          int    `db:"id"`
	Details     string `db:"details"`
	EncodedPass string `db:"hash"`
}

type Word struct {
	IDAccount int    `db:"id_account"`
	ID        int    `db:"id"`
	Value     string `db:"value"`
}
