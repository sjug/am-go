package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const stmtSelectUserByNumber = "SELECT name, cash, dream FROM users WHERE number = ?"

// User database structure
type User struct {
	name         string
	dreamBalance int
	cashBalance  int
}

// SelectUserByNumber function selects user info via collector number
func SelectUserByNumber(num int) (*User, error) {
	tempUser := User{}
	stmt, err := db.Prepare(stmtSelectUserByNumber)
	if err != nil {
		return &tempUser, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(1).Scan(&tempUser)
	if err != nil {
		return &tempUser, err
	}
	return &tempUser, nil
}
