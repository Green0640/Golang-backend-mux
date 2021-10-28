package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase() *sql.DB {
	serverName := "192.168.68.105:3306"
	user := "tctcore"
	password := "tctcore"
	dbName := "tctcore"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("Database connection failed cause by : ", err)
		return nil
	} else {
		log.Println("Database connection OK!")
	}

	return db
}
