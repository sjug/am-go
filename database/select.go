package database

import _ "github.com/go-sql-driver/mysql"

const stmtSelectUserByNumber = "SELECT name, cash, dream FROM users WHERE number = ?"

type User struct {
	name         string
	dreamBalance int
	cashBalance  int
}

func SelectUserByNumber(num int) (User, error) {
	stmt, err := db.Prepare(stmtSelectUserByNumber)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(1).Scan(&user)
	if err != nil {
		return nil, err
	}
	return User, nil
}
