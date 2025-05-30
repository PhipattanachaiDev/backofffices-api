package GpsAntennaControllers

import (
	"database/sql"
	"log"
	"strconv"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/gps-antenna"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/gps-antenna"
	responses "ezview.asia/ezview-web/ezview-lite-back-office/types/responses"

	"github.com/gin-gonic/gin"
)

// GetGpsAntennaStatus ดึงข้อมูล GPS Antenna Status
// @Summary ดึงข้อมูล GPS Antenna Status
// @Description API นี้ใช้สำหรับดึงข้อมูลสถานะแบตเตอรี่จากฐานข้อมูล
// @Tags Gps Antenna
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWGpsAntennaStatusModelResponse] "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/gps-antenna/status [get]
func GetGpsAntennaStatus(c *gin.Context) {
	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM public.ezw_get_gps_antennas_status()")
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var status []dto.EZWGpsAntennaStatusModelResponse

	status = append(status, dto.EZWGpsAntennaStatusModelResponse{
		StatusId:   0,
		StatusName: "ทั้งหมด",
	})
	for rows.Next() {
		var gpsAntenna models.EZWGpsAntennaStatus
		if err := rows.Scan(&gpsAntenna.StatusId, &gpsAntenna.StatusName); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		status = append(status, dto.EZWGpsAntennaStatusModelResponse{
			StatusId:   gpsAntenna.StatusId,
			StatusName: gpsAntenna.StatusName,
		})
	}

	responses.OK(c, status)
}

// SearchGpsAntenna ค้นหา Gps Antenna ตามเงื่อนไขพารามิเตอร์
// @Summary ค้นหา Gps Antenna
// @Description รับ JSON และส่งพารามิเตอร์ไปยังฟังก์ชัน ezw_search_gps_antenna
// @Tags Gps Antenna
// @Accept json
// @Produce json
// @Security Bearer
// @Param searchRequest body dto.EZWSearchGpsAntennaModelRequest true "Payload สำหรับค้นหา Gps Antenna"
// @Success 201 {object} responses.SuccessResponseSwagger[dto.EZWSearchGpsAntennaModelResponse] "ผลลัพธ์การค้นหา Gps Antenna"
// @Failure 400 {object} responses.BadRequestResponseSwagger "ข้อมูลที่ส่งมาไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.Response "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/gps-antenna/search [post]
func SearchGpsAntenna(c *gin.Context) {
	var reqBody dto.EZWSearchGpsAntennaModelRequest
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
		SELECT * FROM public.ezw_search_gps_antenna(
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

	var results []dto.EZWSearchGpsAntennaModelResponse
	for rows.Next() {
		var gpsAntenna models.EZWSearchGpsAntenna
		if err := rows.Scan(
			&gpsAntenna.GpsAntennaId, // column 1: gsm_antenna_id
			&gpsAntenna.SerialNo,     // column 2: serial_no
			&gpsAntenna.StatusId,     // column 3: status_id
			&gpsAntenna.StatusCode,   // column 4: status_code
			&gpsAntenna.StatusName,   // column 5: status_name
			&gpsAntenna.Remark,       // column 6: remark

		); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		results = append(results, dto.EZWSearchGpsAntennaModelResponse{
			GpsAntennaId: gpsAntenna.GpsAntennaId,
			SerialNo:     gpsAntenna.SerialNo,
			StatusId:     gpsAntenna.StatusId,
			StatusCode:   gpsAntenna.StatusCode,
			StatusName:   gpsAntenna.StatusName,
			Remark:       gpsAntenna.Remark,
		})
	}

	if len(results) == 0 {
		responses.OK(c, []dto.EZWSearchGpsAntennaModelResponse{})
		return
	}

	responses.OK(c, results)

}

// InsertGpsAntenna เพิ่มข้อมูล Gps Antenna
// @Summary เพิ่มข้อมูล Gps Antenna
// @Description เพิ่ม Gps Antenna ใหม่ลงในระบบ
// @Tags Gps Antenna
// @Accept json
// @Produce json
// @Security Bearer
// @Param gsmAntenna body dto.EZWInsertGpsAntennaRequest true "ข้อมูล Gps Antenna ที่ต้องการเพิ่ม"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWInsertGpsAntennaResponse]
// @Failure 400 {object} responses.BadRequestResponseSwagger
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/gps-antenna/insert [post]
func InsertGpsAntenna(c *gin.Context) {
	var req dto.EZWInsertGpsAntennaRequest
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
	// เรียกใช้งานฟังก์ชัน ezw_insert_gsm_antenna
	err := db.QueryRow(`
        SELECT public.ezw_insert_gsm_antenna($1, $2, $3, $4)
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
	responses.OK(c, dto.EZWInsertGpsAntennaResponse{Message: result})
}

// UpdateGpsAntenna แก้ไขข้อมูล Gps Antenna
// @Summary แก้ไขข้อมูล Gps Antenna
// @Description แก้ไขข้อมูล Gps Antenna ที่มีอยู่ในระบบ
// @Tags Gps Antenna
// @Accept json
// @Produce json
// @Security Bearer
// @Param gsmAntenna body dto.EZWUpdateGpsAntennaRequest true "ข้อมูล Gps Antenna ที่ต้องการอัปเดต"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWUpdateGpsAntennaResponse]
// @Failure 400 {object} responses.BadRequestResponseSwagger
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/gps-antenna/update [put]
func UpdateGpsAntenna(c *gin.Context) {
	var req dto.EZWUpdateGpsAntennaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// ตรวจสอบข้อมูลแบบง่าย ๆ
	if req.GpsAntennaId <= 0 {
		responses.BadRequest(c, "gps_antenna_id must be a positive integer")
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
	// เรียกใช้งานฟังก์ชัน ezw_update_gsm_antenna
	err := db.QueryRow(`
        SELECT public.ezw_update_gsm_antenna($1, $2, $3, $4, $5)
    `,
		req.GpsAntennaId, // p_gps_antenna_id
		req.SerialNo,     // p_serial_no
		req.StatusId,     // p_status_id
		req.Remark,       // p_remark
		userID,           // p_user_id
	).Scan(&result)

	if err != nil {
		// กรณี record ไม่พบ หรือ error อื่นๆ
		if err == sql.ErrNoRows {
			responses.NotFound(c, "Gps Antenna not found")
			return
		}
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.OK(c, dto.EZWUpdateGpsAntennaResponse{Message: result})
}

// GetGpsAntennaGeneralById godoc
// @Summary      Get GPS‑Antenna general information
// @Tags 		 Gps Antenna
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        gps_antenna_id query int true "GPS‑Antenna ID"
// @Success      200 {object} responses.SuccessResponseSwagger[dto.GpsAntennaGeneralResponse]
// @Failure      400 {object} responses.BadRequestResponseSwagger
// @Failure      500 {object} responses.InternalServerErrorResponseSwagger
// @Router       /v1/gps-antenna/general [get]
func GetGpsAntennaGeneralById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("gps_antenna_id"))
	if err != nil || id <= 0 {
		responses.BadRequest(c, "invalid gps_antenna_id")
		return
	}

	db := database.GetDB()
	if db == nil {
		responses.InternalServerError(c, "db error")
		return
	}

	var m models.GpsAntennaGeneralModel
	if err := db.QueryRow(`SELECT * FROM public.ezw_get_gps_antenna_general_by_id($1)`, id).Scan(
		&m.GpsAntennaID, &m.SerialNo, &m.StatusID, &m.StatusName, &m.Remark,
	); err != nil {
		if err == sql.ErrNoRows {
			responses.OK(c, nil)
			return
		}
		log.Println(err)
		responses.InternalServerError(c, "query error")
		return
	}

	resp := dto.GpsAntennaGeneralResponse{}
	if m.GpsAntennaID.Valid {
		tmp := int(m.GpsAntennaID.Int64)
		resp.GpsAntennaID = &tmp
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
