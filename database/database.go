package database

import (
	_ "github.com/denisenkom/go-mssqldb"

	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

type DatabaseCon struct {
	connection *sql.DB
	status     bool
}

func initDbConnection() (db *sql.DB, err error) {
	server := os.Getenv("DB_SERVER")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	database := os.Getenv("DB_DATABASE")

	connectionString := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s;", server, port, user, password, database)

	db, err = sql.Open("mssql", connectionString)
	err = db.Ping()

	return
}

func (db *DatabaseCon) ConnectToDatabaseWithRetry() error {
	var err error

	// Define the maximum number of retry attempts and the delay between each attempt
	maxRetries := 10
	retryDelay := time.Second

	for i := 0; i < maxRetries; i++ {
		db.connection, err = initDbConnection()
		err = db.connection.Ping()
		if err != nil {
			fmt.Println("Failed to connect to the database.")
			time.Sleep(retryDelay)
		} else {
			return nil
		}
	}

	return fmt.Errorf("failed to connect to the database after %d attempts", maxRetries)
}

func (db *DatabaseCon) GetDatabaseCon() *sql.DB {
	return db.connection
}

func (db *DatabaseCon) isDatabaseConnectionUp() bool {
	err := db.connection.Ping()
	if err != nil {
		return false
	}
	return true
}

func (db *DatabaseCon) StartConnectionCheck(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			if db.connection != nil {
				if err := db.connection.Ping(); err != nil {
					db.status = false
					log.Printf("Database connection is down: %s", err.Error())
				} else {
					db.status = true
					log.Printf("Database connection is up")
				}
			}

		}
	}()
}
