package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const stmtSelectUserByNumber = "SELECT name, dream, cash FROM users WHERE number = ?"

// User database structure
type User struct {
	number       int
	name         string
	dreamBalance int
	cashBalance  int
}

// SelectUserByNumber function selects user info via collector number
func SelectUserByNumber(num int) (*User, error) {
	tempUser := User{}
	stmt, err := db.Prepare(stmtSelectUserByNumber)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(1).Scan(&tempUser.name, &tempUser.dreamBalance, &tempUser.cashBalance)
	if err != nil {
		return nil, err
	}
	return &tempUser, nil
}
