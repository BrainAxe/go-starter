package main

import (
	"go-starter/src/dao"
	"go-starter/src/routes"

	"github.com/joho/godotenv"
)

func main() {
	//Load env file
	er := godotenv.Load()
	if er != nil {
		panic(er)
	}

	//Connect to the database
	err := dao.InitDB()
	if err != nil {
		panic(err)
	}
	//The program exits and closes the database connection
	defer dao.Close()
	//Register routing
	r := routes.SetRouter()
	//Start the project with port 8000
	r.Run(":8000")
}
