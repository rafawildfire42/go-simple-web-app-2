package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}

func ConnectDatabase() *sql.DB {

	dbDataConnection := os.Getenv("DB_CONNECTION")
	db, err := sql.Open("mysql", dbDataConnection)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db

}
