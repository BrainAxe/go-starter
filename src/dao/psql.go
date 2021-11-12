package dao

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Specify driver
const DRIVER = "postgres" // Untyped named string constant

var SqlSession *gorm.DB

//Initialize the connection to the database and generate variables that can operate basic addition, deletion, and modification of the structure
func InitDB() (err error) {
	//Connect to the database
	SqlSession, err = gorm.Open(DRIVER, ConnectUrl())
	if err != nil {
		panic(err)
	}

	//Verify whether the database connection is successful, if successful, there is no exception
	return SqlSession.DB().Ping()
}

//Database connection string
func ConnectUrl() string {
	//env configuration parameters into a url to connect to the database
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("url"),
		os.Getenv("port"),
		os.Getenv("username"),
		os.Getenv("password"),
		os.Getenv("dbname"),
	)
}

//Close the database connection
func Close() {
	SqlSession.Close()
}
