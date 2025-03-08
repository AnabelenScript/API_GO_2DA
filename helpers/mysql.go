package helpers

import (
	"database/sql"
	"fmt"
	"log"
	
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQL() *sql.DB {
	dsn := "root:Ana0507belen@tcp(localhost:3306)/api_hexagonal"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping MySQL: %v", err)
	}

	fmt.Println("Connected to MySQL successfully")
	return db
}
