package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/sjug/am-go/structure"
)

const stmtGetUserDetailsFromNumber = "SELECT mem.firstName, acc.dreamBalance, acc.cashBalance FROM `apiPOC`.`ACCOUNT` acc JOIN `apiPOC`.`MEMBER` mem ON acc.collectorNumber = mem.collectorNumber WHERE acc.collectorNumber = ?"

// GetUserDetailsFromNumber gets the user details via MySQL connection
func GetUserDetailsFromNumber(num int) (*structure.CollectorDetails, error) {
	db, err := sql.Open("mysql", "root:password@tcp(tor-ovn-7004.loyalty.com:3306)/apiPOC")
	// db, err := sql.Open("mysql", "test:test_P455@tcp(127.0.0.1:3306)/apiPOC")
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
	defer stmt.Close()

	tempCollector := structure.CollectorDetails{}
	err = stmt.QueryRow(num).Scan(&tempCollector.CollectorName, &tempCollector.DreamBalance, &tempCollector.CashBalance)
	if err != nil {
		return &tempCollector, err
	}
	return &tempCollector, nil
}
