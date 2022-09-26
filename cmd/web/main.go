package main

import (
	"log"
	"os"

	"github.com/ffelipelimao/delivery-service/internal/framework/database"
	"github.com/ffelipelimao/delivery-service/internal/framework/server"
	"github.com/joho/godotenv"
)

var db database.Database

func init() {

	//Remove this code to Debug Mode in Vscode
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.DsnTest = os.Getenv("DSN_TEST")
	db.Dsn = os.Getenv("DSN")
	db.DbTypeTest = os.Getenv("DB_TYPE_TEST")
	db.DbType = os.Getenv("DB_TYPE")
	db.Env = os.Getenv("ENV")
}

func main() {

	DBConnection, err := db.Connect()
	if err != nil {
		log.Fatalf("error connecting to DB")
	}
	defer DBConnection.Close()

	server.Start(DBConnection)

}
