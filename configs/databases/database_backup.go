package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// var db *sql.DB

// func ConnectDB() *sql.DB {
// 	// Get environment variables
// 	server := os.Getenv("DB_SERVER")
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	portNumber := os.Getenv("DB_PORT")
// 	database := os.Getenv("DB_DATABASE")

// 	// Validate environment variables
// 	if server == "" || user == "" || password == "" || portNumber == "" || database == "" {
// 		log.Fatal("One or more required environment variables are missing")
// 	}

// 	portNumberInt, err := strconv.Atoi(portNumber)
// 	if err != nil {
// 		log.Printf("Error converting portNumber to integer: %v", err)
// 	}

// 	// Setup database connection
// 	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		server, portNumberInt, user, password, database)

// 	db, err = sql.Open("postgres", connectionString)
// 	if err != nil {
// 		log.Printf("Error opening database connection: %v", err)
// 	}

// 	// Configure connection pooling
// 	db.SetMaxOpenConns(100)
// 	db.SetMaxIdleConns(10)
// 	db.SetConnMaxLifetime(5 * time.Minute)

// 	// Ping the database to verify connection
// 	if err = db.Ping(); err != nil {
// 		log.Printf("Error pinging database: %v", err)
// 	}

// 	fmt.Println("Connected to PostgreSQL!")

// 	return db
// }
