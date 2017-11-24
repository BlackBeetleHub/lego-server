package entries

type Account struct {
	ID 	      int    `db:"id"`
	Details   string `db:"details"`
}

type Word struct {
	ID				int    `db:"id"`
	IDAccount		int    `db:"id_account"`
	Value    		string `db:"value"`
}