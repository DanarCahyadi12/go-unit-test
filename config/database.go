package config

import "database/sql"

func NewDatabase() *sql.DB {
	db, err := sql.Open("mysql", "<your_username>:<your_password>@tcp(<your_host>:<your_port>)/<your_database>")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}
