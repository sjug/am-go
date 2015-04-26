package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Init func initializes mysql db connection
func Init() error {
	con, err := sql.Open("mysql", store.user+":"+store.password+"@/"+store.database)
	defer con.Close()
	if err != nil {
		return err
	}
}
