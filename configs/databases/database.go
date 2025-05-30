package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// ConnectDB เชื่อมต่อฐานข้อมูล (singleton)
func ConnectDB() *sql.DB {
	once.Do(func() {
		server := os.Getenv("DB_SERVER")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		portNumber := os.Getenv("DB_PORT")
		database := os.Getenv("DB_DATABASE")

		// ตรวจสอบว่า env variable ถูกต้อง
		if server == "" || user == "" || password == "" || portNumber == "" || database == "" {
			log.Fatal("One or more required environment variables are missing")
		}

		portNumberInt, err := strconv.Atoi(portNumber)
		if err != nil {
			log.Fatalf("Error converting portNumber to integer: %v", err)
		}

		// Connection string
		connectionString := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			server, portNumberInt, user, password, database,
		)

		// เปิด Connection
		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			log.Printf("Error opening database connection: %v", err)
			db = nil
			return
		}

		// ตั้งค่า Connection Pool
		db.SetMaxOpenConns(50)                 // จำกัด connection สูงสุด
		db.SetMaxIdleConns(10)                 // จำกัด idle connection
		db.SetConnMaxLifetime(5 * time.Minute) // อายุ connection
		db.SetConnMaxIdleTime(1 * time.Minute) // ปิด connection ที่ idle นานเกินไป

		// Ping ตรวจสอบว่าสามารถเชื่อมต่อฐานข้อมูลได้
		if err = db.Ping(); err != nil {
			log.Printf("Error pinging database: %v", err)
			db = nil // ป้องกันการคืนค่า database ที่มีปัญหา
		}

		log.Println("Connected to PostgreSQL!")
	})

	return db
}

// GetDB ใช้สำหรับดึง instance ของ database
func GetDB() *sql.DB {
	if db == nil {
		log.Println("Warning: Database connection is not initialized.")
		return nil
	}
	return db
}
