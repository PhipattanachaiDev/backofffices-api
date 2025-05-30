package usersController

import (
	"database/sql"
	"log"
	"strconv"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/users"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/users"
	"ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
		return
	}

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}
	// ใช้ Prepared Statement (แบบไม่ต้องเรียก `Prepare()` เอง)
	query := `SELECT * FROM ezw_get_user_profile_by_id($1)`
	row := db.QueryRow(query, userID)

	// สร้าง response สำหรับข้อมูลลูกค้า
	var response dto.GetUserResponse

	var user models.EZWGetUser

	// ดึงข้อมูลลูกค้า
	err := row.Scan(&user.UserId, &user.UserName, &user.Name, &user.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.NotFound(c, "User not found")
			return
		} else {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch user data")
			return
		}
	}

	response.UserId = user.UserId
	response.UserName = user.UserName
	response.Name = user.Name
	response.Description = user.Description

	responses.OK(c, response)
}

func UpdateUser(c *gin.Context) {
	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
		return
	}

	var reqBody dto.UpdateUserRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	id := reqBody.UserId
	username := reqBody.Username
	password := reqBody.Password
	name := reqBody.Name
	description := reqBody.Description

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	query := `SELECT * FROM ezw_update_user($1, $2, $3, $4, $5, $6)`
	row := db.QueryRow(query, id, username, password, name, description, userID)

	var updatedUserID int
	err := row.Scan(&updatedUserID)

	if err != nil {
		// Check for specific error message
		if err.Error() == "pq: Login name "+username+" is already taken by another user." {
			responses.BadRequest(c, "Login name already exists")
			return
		}

		if err.Error() == "pq: User ID "+strconv.Itoa(id)+" does not exist." {
			responses.BadRequest(c, "User ID does not exist")
			return
		}

		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.OK(c, gin.H{"message": "User updated successfully", "user_id": updatedUserID})
}

func DeleteUser(c *gin.Context) {
	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
		return
	}

	id := c.Param("id")

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	query := `SELECT * FROM ezw_delete_user($1, $2)`
	row := db.QueryRow(query, id, userID)

	var deletedUserID int
	err := row.Scan(&deletedUserID)

	if err != nil {
		if err.Error() == "pq: User ID "+id+" does not exist." {
			responses.BadRequest(c, "User ID does not exist")
			return
		}
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.OK(c, gin.H{"message": "User deleted successfully", "user_id": deletedUserID})
}
