package BatteryControllers

import (
	"database/sql"
	"log"
	"strconv"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/battery"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/battery"
	responses "ezview.asia/ezview-web/ezview-lite-back-office/types/responses"

	"github.com/gin-gonic/gin"
)

// GetBatteryStatus ดึงข้อมูล Battery Status
// @Summary ดึงข้อมูล Battery Status
// @Description API นี้ใช้สำหรับดึงข้อมูลสถานะแบตเตอรี่จากฐานข้อมูล
// @Tags Battery
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWBatteryStatusModelResponse] "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/battery/status [get]
func GetBatteryStatus(c *gin.Context) {
	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM public.ezw_get_batteries_status()")
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var statuses []dto.EZWBatteryStatusModelResponse

	statuses = append(statuses, dto.EZWBatteryStatusModelResponse{

		StatusId:   0,
		StatusName: "ทั้งหมด",
	})
	for rows.Next() {
		var battery models.EZWBatteryStatus
		if err := rows.Scan(&battery.StatusId, &battery.StatusName); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		status := dto.EZWBatteryStatusModelResponse{
			StatusId:   battery.StatusId,
			StatusName: battery.StatusName,
		}
		statuses = append(statuses, status)
	}

	responses.OK(c, statuses)
}

// SearchBattery ค้นหา Battery ตามเงื่อนไขพารามิเตอร์
// @Summary ค้นหา Battery
// @Description รับ JSON และส่งพารามิเตอร์ไปยังฟังก์ชัน ezw_search_battery
// @Tags Battery
// @Accept json
// @Produce json
// @Security Bearer
// @Param searchRequest body dto.EZWSearchBatteryModelRequest true "Payload สำหรับค้นหา Battery"
// @Success 201 {object} responses.SuccessResponseSwagger[dto.EZWSearchBatteryModelResponse] "ผลลัพธ์การค้นหา Battery"
// @Failure 400 {object} responses.BadRequestResponseSwagger "ข้อมูลที่ส่งมาไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.Response "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/battery/search [post]
func SearchBattery(c *gin.Context) {
	var reqBody dto.EZWSearchBatteryModelRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
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

	// เตรียมตัวแปรเป็น sql.Nullxxx
	var (
		SerialNoNull sql.NullString
		statusIDNull sql.NullInt32
		remarkNull   sql.NullString
	)

	// ถ้า string ไม่ว่าง
	if reqBody.SerialNo != "" {
		SerialNoNull = sql.NullString{String: reqBody.SerialNo, Valid: true}
	}
	// ถ้าเป็น 0 → ถือว่า NULL, แต่ถ้าเป็น 99 → ส่งเป็น valid เพื่อให้ PostgreSQL มองว่า “เลือกทั้งหมด”
	if reqBody.StatusId != 0 {
		statusIDNull = sql.NullInt32{Int32: int32(reqBody.StatusId), Valid: true}
	}
	if reqBody.Remark != "" {
		remarkNull = sql.NullString{String: reqBody.Remark, Valid: true}
	}

	query := `
		SELECT * FROM public.ezw_search_batteries(
			$1::varchar,
			$2::int,
			$3::varchar
		)
	`
	rows, err := db.Query(query,
		SerialNoNull,
		statusIDNull,
		remarkNull,
	)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var results []dto.EZWSearchBatteryModelResponse
	for rows.Next() {
		var battery models.EZWSearchBattery
		if err := rows.Scan(
			&battery.BatteryId,  // column 1: battery_id
			&battery.SerialNo,   // column 2: serial_no
			&battery.StatusId,   // column 3: status_id
			&battery.StatusName, // column 4: status_name
			&battery.Remark,     // column 5: remark

		); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		results = append(results, dto.EZWSearchBatteryModelResponse{
			BatteryId:  battery.BatteryId,
			SerialNo:   battery.SerialNo,
			StatusId:   battery.StatusId,
			StatusName: battery.StatusName,
			Remark:     battery.Remark,
		})
	}

	if len(results) == 0 {
		responses.OK(c, []dto.EZWSearchBatteryModelResponse{})
		return
	}

	responses.OK(c, results)

}

// InsertBattery เพิ่มข้อมูล Battery
// @Summary เพิ่มข้อมูล Battery
// @Description เพิ่ม Battery ใหม่ลงในระบบ
// @Tags Battery
// @Accept json
// @Produce json
// @Security Bearer
// @Param gsmAntenna body dto.EZWInsertBatteryRequest true "ข้อมูล Battery ที่ต้องการเพิ่ม"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWInsertBatteryResponse]
// @Failure 400 {object} responses.BadRequestResponseSwagger
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/battery/insert [post]
func InsertBattery(c *gin.Context) {
	var req dto.EZWInsertBatteryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// ตรวจสอบข้อมูลแบบง่าย ๆ
	if req.SerialNo == "" {
		responses.BadRequest(c, "serial_no is required and should not be empty")
		return
	}
	if req.StatusId <= 0 {
		responses.BadRequest(c, "status_id must be a positive integer")
		return
	}
	// Remark จะเป็น string เสมอ แต่อาจเช็คเพิ่มได้ตาม business logic

	// ดึง user_id จาก context (JWT middleware)
	userIDVal, exists := c.Get("user_id")
	if !exists {
		responses.Unauthorized(c, "User ID not found in context")
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		responses.Unauthorized(c, "User ID in context is invalid")
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

	var result string
	// เรียกใช้งานฟังก์ชัน ezw_insert_battery
	err := db.QueryRow(`
        SELECT public.ezw_insert_battery($1, $2, $3, $4)
    `,
		req.SerialNo, // p_serial_no
		req.StatusId, // p_status_id
		req.Remark,   // p_remark
		userID,       // p_user_id
	).Scan(&result)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	// ตอบกลับ
	responses.OK(c, dto.EZWInsertBatteryResponse{Message: result})
}

// UpdateBattery แก้ไขข้อมูล Battery
// @Summary แก้ไขข้อมูล Battery
// @Description แก้ไขข้อมูล Battery ที่มีอยู่ในระบบ
// @Tags Battery
// @Accept json
// @Produce json
// @Security Bearer
// @Param gsmAntenna body dto.EZWUpdateBatteryRequest true "ข้อมูล Battery ที่ต้องการอัปเดต"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWUpdateBatteryResponse]
// @Failure 400 {object} responses.BadRequestResponseSwagger
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/battery/update [put]
func UpdateBattery(c *gin.Context) {
	var req dto.EZWUpdateBatteryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// ตรวจสอบข้อมูลแบบง่าย ๆ
	if req.BatteryId <= 0 {
		responses.BadRequest(c, "battery_id must be a positive integer")
		return
	}
	if req.SerialNo == "" {
		responses.BadRequest(c, "serial_no is required and should not be empty")
		return
	}
	if req.StatusId <= 0 {
		responses.BadRequest(c, "status_id must be a positive integer")
		return
	}

	// ดึง user_id จาก context (JWT middleware)
	userIDVal, exists := c.Get("user_id")
	if !exists {
		responses.Unauthorized(c, "User ID not found in context")
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		responses.Unauthorized(c, "User ID in context is invalid")
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

	var result string
	// เรียกใช้งานฟังก์ชัน ezw_update_battery
	err := db.QueryRow(`
        SELECT public.ezw_update_battery($1, $2, $3, $4, $5)
    `,
		req.BatteryId, // p_battery_id
		req.SerialNo,  // p_serial_no
		req.StatusId,  // p_status_id
		req.Remark,    // p_remark
		userID,        // p_user_id
	).Scan(&result)

	if err != nil {
		// กรณี record ไม่พบ หรือ error อื่นๆ
		if err == sql.ErrNoRows {
			responses.NotFound(c, "Battery not found")
			return
		}
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.OK(c, dto.EZWUpdateBatteryResponse{Message: result})
}

// GetBatteryGeneralById godoc
// @Summary      Get Battery general information
// @Tags         Battery
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        battery_id query int true "Battery ID"
// @Success      200 {object} responses.SuccessResponseSwagger[dto.EZWGetBatteryGeneralResponse]
// @Failure      400 {object} responses.BadRequestResponseSwagger
// @Failure      500 {object} responses.InternalServerErrorResponseSwagger
// @Router       /v1/battery/general [get]
func GetBatteryGeneralById(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("battery_id"))
	if err != nil || id <= 0 {
		responses.BadRequest(c, "invalid battery_id")
		return
	}

	db := database.GetDB()
	if db == nil {
		responses.InternalServerError(c, "database connection error")
		return
	}

	var m models.BatteryGeneralModel
	if err := db.QueryRow(`SELECT * FROM public.ezw_get_battery_general_by_id($1)`, id).Scan(
		&m.BatteryID, &m.SerialNo, &m.StatusID, &m.StatusName, &m.Remark,
	); err != nil {
		if err == sql.ErrNoRows {
			responses.OK(c, nil)
			return
		}
		log.Println("query error:", err)
		responses.InternalServerError(c, "query error")
		return
	}

	resp := dto.EZWGetBatteryGeneralResponse{}
	if m.BatteryID.Valid {
		tmp := int(m.BatteryID.Int64)
		resp.BatteryID = &tmp
	}
	if m.SerialNo.Valid {
		tmp := m.SerialNo.String
		resp.SerialNo = &tmp
	}
	if m.StatusID.Valid {
		tmp := int(m.StatusID.Int64)
		resp.StatusID = &tmp
	}
	if m.StatusName.Valid {
		tmp := m.StatusName.String
		resp.StatusName = &tmp
	}
	if m.Remark.Valid {
		tmp := m.Remark.String
		resp.Remark = &tmp
	}

	responses.OK(c, resp)
}
