package employee

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// GetDBconn :
func GetDBconn() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost/employee")

	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()

	// Ping for checking database connected or not
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Print("Successfully Postgresql Database Connected!")
	return db
}
