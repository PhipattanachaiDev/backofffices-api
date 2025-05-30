package menusController

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/menus"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/menus"
	"ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	"github.com/gin-gonic/gin"
)

func GetMenus(c *gin.Context) {
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
	query := `SELECT * FROM ezw_get_menus_by_user($1)`
	row := db.QueryRow(query, userID)

	// สร้าง response สำหรับข้อมูลลูกค้า
	var response dto.GetMenusResponse

	var menus models.EZWGetMenusByUser

	// ดึงข้อมูลลูกค้า
	err := row.Scan(&menus.MenuIdList)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.NotFound(c, "User not found")
			return
		} else {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch user menus")
			return
		}
	}

	jsonData, err := json.Marshal(menus.MenuIdList)
	if err != nil {
		log.Println("Error marshalling MenuIdList to JSON:", err)
		responses.InternalServerError(c, "Failed to process menu data")
		return
	}

	// Check if the JSON data represents an empty array
	if string(jsonData) == "[]" {
		responses.Unauthorized(c, "Unauthorized access")
		return
	}

	fmt.Println("MenuIdList JSON:", string(jsonData))
	if err == nil && len(menus.MenuIdList) == 0 {
		responses.Unauthorized(c, "Unauthorized access")
		return
	}
	if err != nil {
		if err == sql.ErrNoRows {
			responses.NotFound(c, "User not found")
		} else {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch user menus")
		}
		return
	}

	response = dto.GetMenusResponse{
		Menus: menus.MenuIdList,
	}

	responses.OK(c, response)
}
