package queries

import "database/sql"

func CreateUsersTable(db *sql.DB) error {
	q := `CREATE TABLE IF NOT EXISTS users (
		ID SERIAL NOT NULL,
		username varchar(50) NOT NULL,
		password varchar(50) NOT NULL,
		PRIMARY KEY (ID)
	);`

	_, err := db.Exec(q)
	return err
}

func DropUsersTable(db *sql.DB) error {
	q := `DROP TABLE users`

	_, err := db.Exec(q)
	return err
}
