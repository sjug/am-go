package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Init func initializes mysql db connection
func Init() error {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
