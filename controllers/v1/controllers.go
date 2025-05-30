package v1

import (
	"database/sql"
	"log"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	"ezview.asia/ezview-web/ezview-lite-back-office/middlewares"
	response "ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func Admin(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		response.Unauthorized(c, "Missing claims")
		return
	}

	username := claims.(*middlewares.Claims).Username

	// Only allow admin operations for certain users
	if username != "admin" {
		response.Forbidden(c, "Unauthorized")
		return
	}

	// Parse JSON
	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		db[username] = json.Value
		response.OK(c, gin.H{"status": "ok"})
	} else {
		response.BadRequest(c, "Invalid request payload")
	}
}

func Ping(c *gin.Context) {
	response.OK(c, "pong")
}

func GetUser(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := db[user]
	if ok {
		response.OK(c, gin.H{"user": user, "value": value})
	} else {
		response.NotFound(c, "User not found")
	}
}

// GetUserFromDB retrieves user details from the database
func GetUserFromDB(c *gin.Context) {
	username := c.Query("username")

	db := database.ConnectDB()
	defer db.Close()

	// var id int
	// var username string

	// Query the database
	rows, err := db.Query("SELECT uid, username FROM users WHERE username = @username", sql.Named("username", username))

	// Call the stored procedure
	// err := db.QueryRow("CALL GetUserByUsername(?)", username).Scan(&id, &username)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		response.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	// Slice to store users
	var users []gin.H

	// Iterate through rows
	for rows.Next() {
		var id int
		var username string
		if err := rows.Scan(&id, &username); err != nil {
			log.Printf("Error scanning row: %v", err)
			response.InternalServerError(c, "Internal Server Error")
			return
		}
		users = append(users, gin.H{"id": id, "username": username})
	}

	// Check for empty result
	if len(users) == 0 {
		log.Printf("No users found with username: %s", username)
		response.NotFoundWithData(c, "User not found")
		return
	}

	// Respond with users
	response.OK(c, gin.H{"users": users})
}
