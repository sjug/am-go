package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const stmtGetUserDetailsFromNumber = "SELECT mem.firstName, acc.dreamBalance, acc.cashBalance FROM `apiPOC`.`ACCOUNT` acc JOIN `apiPOC`.`MEMBER` mem ON acc.collectorNumber = mem.collectorNumber WHERE acc.collectorNumber = ?"

type Collector struct {
	firstName    string
	dreamBalance int
	cashBalance  int
}

func GetUserDetailsFromNumber(num int) (*Collector, error) {
	db, err := sql.Open("mysql",
		"test:test_P455@tcp(127.0.0.1:3306)/apiPOC")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare(stmtGetUserDetailsFromNumber)
	if err != nil {
		log.Fatal(err)
	}
	tempCollector := Collector{}
	err = stmt.QueryRow(num).Scan(&tempCollector.firstName, &tempCollector.dreamBalance, &tempCollector.cashBalance)
	if err != nil {
		log.Fatal(err)
	}
	return &tempCollector, nil
}
