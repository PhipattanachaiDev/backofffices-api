package dealersController

import (
	"database/sql"
	"log"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/dealers"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/dealers"
	"ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	"github.com/gin-gonic/gin"
)

func GetDealersByConditions(c *gin.Context) {

	var reqBody dto.GetDealersByConditionsRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	dealerId := sql.NullString{}
	if reqBody.DealerId != "" {
		dealerId = sql.NullString{String: reqBody.DealerId, Valid: true}
	}

	dealerName := sql.NullString{}
	if reqBody.DealerName != "" {
		dealerName = sql.NullString{String: reqBody.DealerName, Valid: true}
	}

	dealerGroup := sql.NullInt64{}
	if reqBody.DealerGroup != 0 {
		dealerGroup = sql.NullInt64{Int64: reqBody.DealerGroup, Valid: true}
	}

	dealerDetail := sql.NullString{}
	if reqBody.DealerDetail != "" {
		dealerDetail = sql.NullString{String: reqBody.DealerDetail, Valid: true}
	}

	dealerStatus := sql.NullInt64{}
	if reqBody.DealerStatus != 0 {
		dealerStatus = sql.NullInt64{Int64: reqBody.DealerStatus, Valid: true}
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
	query := `SELECT * FROM ezw_search_dealers($1, $2, $3, $4, $5)`
	rows, err := db.Query(query, dealerId, dealerName, dealerGroup, dealerDetail, dealerStatus)

	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch dealers")
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูล dealer
	var response []dto.GetDealersByConditionsResponse

	// วนลูปเพื่อดึงข้อมูล dealer
	for rows.Next() {
		var dealer models.EZWSearchDealers
		err := rows.Scan(&dealer.DealerId,
			&dealer.DealerName,
			&dealer.DealerGroupId,
			&dealer.DealerGroup,
			&dealer.DealerDetail,
			&dealer.DealerStatus)
		if err != nil {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch dealers")
			return
		}

		response = append(response, dto.GetDealersByConditionsResponse{
			DealerId:      dealer.DealerId,
			DealerName:    dealer.DealerName,
			DealerGroupId: dealer.DealerGroupId,
			DealerGroup:   dealer.DealerGroup,
			DealerDetail:  dealer.DealerDetail,
			DealerStatus:  dealer.DealerStatus,
		})
	}

	if len(response) == 0 {
		responses.OK(c, []dto.GetDealersByConditionsResponse{})
		return
	}

	responses.OK(c, response)
}

func GetDealerGroups(c *gin.Context) {

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}
	// ใช้ Prepared Statement (แบบไม่ต้องเรียก `Prepare()` เอง)
	query := `select type_id, type_name from system_master_types smt where smt.category_id = 109`
	rows, err := db.Query(query)

	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch customers")
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลลูกค้า
	var response []dto.GetDealerGroupsResponse

	// วนลูปเพื่อดึงข้อมูลลูกค้า
	for rows.Next() {
		var dealerGroup models.EZWGetDealerGroup
		err := rows.Scan(&dealerGroup.DealerGroupId,
			&dealerGroup.DealerGroupName)
		if err != nil {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch dealers")
			return
		}

		response = append(response, dto.GetDealerGroupsResponse{
			DealerGroupId:   dealerGroup.DealerGroupId,
			DealerGroupName: dealerGroup.DealerGroupName,
		})
	}

	if len(response) == 0 {
		responses.OK(c, []dto.GetDealerGroupsResponse{})
		return
	}

	responses.OK(c, response)
}

func GetDealerStatus(c *gin.Context) {

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}
	// ใช้ Prepared Statement (แบบไม่ต้องเรียก `Prepare()` เอง)
	query := `select type_id, type_name from system_master_types smt where smt.category_id = 108`
	rows, err := db.Query(query)

	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch dealers")
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลลูกค้า
	var response []dto.GetDealerStatusResponse

	// วนลูปเพื่อดึงข้อมูลลูกค้า
	for rows.Next() {
		var dealerStatus models.EZWDealerStatus
		err := rows.Scan(&dealerStatus.StatusId,
			&dealerStatus.StatusName)
		if err != nil {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch dealers")
			return
		}

		response = append(response, dto.GetDealerStatusResponse{
			StatusId:   dealerStatus.StatusId,
			StatusName: dealerStatus.StatusName,
		})
	}

	if len(response) == 0 {
		responses.OK(c, []dto.GetDealerStatusResponse{})
		return
	}

	responses.OK(c, response)
}
