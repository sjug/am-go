package database

import _ "github.com/go-sql-driver/mysql"

const stmtSelectUserByNumber = "SELECT name, cash, dream FROM users WHERE number = ?"

type user struct {
	name         string
	dreamBalance int
	cashBalance  int
}

// SelectUserByNumber function selects user info via collector number
func SelectUserByNumber(num int) (user, error) {
	stmt, err := db.Prepare(stmtSelectUserByNumber)
	if err != nil {
		return &user, err
	}

	err = stmt.QueryRow(1).Scan(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}
