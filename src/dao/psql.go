package dao // Data Access Object

import (
	"database/sql"
	_ "github.com/lib/pq"
	db "go-starter/src/db/sqlc"
	"log"
	"os"
)

var SqlSession *db.SQLStore
var conn *sql.DB

// Initialize the connection to the database and generate variables that can operate basic addition, deletion, and modification of the structure
func InitDB() (err error) {
	//Connect to the database
	dbSource := os.Getenv("DB_SOURCE")
	conn, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	SqlSession = db.NewStore(conn)

	//Verify whether the database connection is successful, if successful, there is no exception
	return conn.Ping()
}

// Close the database connection
func Close() {
	conn.Close()
}
