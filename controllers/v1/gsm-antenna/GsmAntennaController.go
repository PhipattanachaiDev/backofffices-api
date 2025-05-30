package GsmAntennaControllers

import (
	"database/sql"
	"log"
	"strconv"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/gsm-antenna"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/gsm-antenna"

	responses "ezview.asia/ezview-web/ezview-lite-back-office/types/responses"

	"github.com/gin-gonic/gin"
)

// GetGsmAntennaStatus ดึงข้อมูล Gsm Antenna Status
// @Summary ดึงข้อมูล Gsm Antenna Status
// @Description API นี้ใช้สำหรับดึงข้อมูลสถานะแบตเตอรี่จากฐานข้อมูล
// @Tags Gsm Antenna
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWGsmAntennaStatusModelResponse] "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/gsm-antenna/status [get]
func GetGsmAntennaStatus(c *gin.Context) {
	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM public.ezw_get_gsm_antennas_status()")
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var status []dto.EZWGsmAntennaStatusModelResponse

	status = append(status, dto.EZWGsmAntennaStatusModelResponse{
		StatusId:   0,
		StatusName: "ทั้งหมด",
	})
	for rows.Next() {
		var gsmAntenna models.EZWGetGsmAntennaStatus
		if err := rows.Scan(&gsmAntenna.StatusId, &gsmAntenna.StatusName); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		status = append(status, dto.EZWGsmAntennaStatusModelResponse{
			StatusId:   gsmAntenna.StatusId,
			StatusName: gsmAntenna.StatusName,
		})
	}

	responses.OK(c, status)
}

// SearchGsmAntenna ค้นหา Gsm Antenna ตามเงื่อนไขพารามิเตอร์
// @Summary ค้นหา Gsm Antenna
// @Description รับ JSON และส่งพารามิเตอร์ไปยังฟังก์ชัน ezw_search_gsm_antenna
// @Tags Gsm Antenna
// @Accept json
// @Produce json
// @Security Bearer
// @Param searchRequest body dto.EZWSearchGsmAntennaModelRequest true "Payload สำหรับค้นหา Gsm Antenna"
// @Success 201 {object} responses.SuccessResponseSwagger[dto.EZWSearchGsmAntennaModelResponse] "ผลลัพธ์การค้นหา Gsm Antenna"
// @Failure 400 {object} responses.BadRequestResponseSwagger "ข้อมูลที่ส่งมาไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.Response "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/gsm-antenna/search [post]
func SearchGsmAntenna(c *gin.Context) {
	var reqBody dto.EZWSearchGsmAntennaModelRequest
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
		SELECT * FROM public.ezw_search_gsm_antenna(
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

	var results []dto.EZWSearchGsmAntennaModelResponse
	for rows.Next() {
		var gsmAntenna models.EZWSearchGsmAntenna
		if err := rows.Scan(
			&gsmAntenna.GsmAntennaId, // column 1: gsm_antenna_id
			&gsmAntenna.SerialNo,     // column 2: serial_no
			&gsmAntenna.StatusId,     // column 3: status_id
			&gsmAntenna.StatusCode,   // column 4: status_code
			&gsmAntenna.StatusName,   // column 5: status_name
			&gsmAntenna.Remark,       // column 6: remark

		); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		results = append(results, dto.EZWSearchGsmAntennaModelResponse{
			GsmAntennaId: gsmAntenna.GsmAntennaId,
			SerialNo:     gsmAntenna.SerialNo,
			StatusId:     gsmAntenna.StatusId,
			StatusCode:   gsmAntenna.StatusCode,
			StatusName:   gsmAntenna.StatusName,
			Remark:       gsmAntenna.Remark,
		})
	}

	if len(results) == 0 {
		responses.OK(c, []dto.EZWSearchGsmAntennaModelResponse{})
		return
	}

	responses.OK(c, results)

}

// InsertGsmAntenna เพิ่มข้อมูล Gsm Antenna
// @Summary เพิ่มข้อมูล Gsm Antenna
// @Description เพิ่ม Gsm Antenna ใหม่ลงในระบบ
// @Tags Gsm Antenna
// @Accept json
// @Produce json
// @Security Bearer
// @Param gsmAntenna body dto.EZWInsertGsmAntennaRequest true "ข้อมูล Gsm Antenna ที่ต้องการเพิ่ม"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWInsertGsmAntennaResponse]
// @Failure 400 {object} responses.BadRequestResponseSwagger
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/gsm-antenna/insert [post]
func InsertGsmAntenna(c *gin.Context) {
	var req dto.EZWInsertGsmAntennaRequest
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
	responses.OK(c, dto.EZWInsertGsmAntennaResponse{Message: result})
}

// UpdateGsmAntenna แก้ไขข้อมูล Gsm Antenna
// @Summary แก้ไขข้อมูล Gsm Antenna
// @Description แก้ไขข้อมูล Gsm Antenna ที่มีอยู่ในระบบ
// @Tags Gsm Antenna
// @Accept json
// @Produce json
// @Security Bearer
// @Param gsmAntenna body dto.EZWUpdateGsmAntennaRequest true "ข้อมูล Gsm Antenna ที่ต้องการอัปเดต"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWUpdateGsmAntennaResponse]
// @Failure 400 {object} responses.BadRequestResponseSwagger
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/gsm-antenna/update [put]
func UpdateGsmAntenna(c *gin.Context) {
	var req dto.EZWUpdateGsmAntennaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// ตรวจสอบข้อมูลแบบง่าย ๆ
	if req.GsmAntennaId <= 0 {
		responses.BadRequest(c, "gsm_antenna_id must be a positive integer")
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
		req.GsmAntennaId, // p_gsm_antenna_id
		req.SerialNo,     // p_serial_no
		req.StatusId,     // p_status_id
		req.Remark,       // p_remark
		userID,           // p_user_id
	).Scan(&result)

	if err != nil {
		// กรณี record ไม่พบ หรือ error อื่นๆ
		if err == sql.ErrNoRows {
			responses.NotFound(c, "Gsm Antenna not found")
			return
		}
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.OK(c, dto.EZWUpdateGsmAntennaResponse{Message: result})
}

// GetGsmAntennaGeneralById godoc
// @Summary      Get GSM‑Antenna general information
// @Tags         Gsm Antenna
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        gsm_antenna_id query int true "GSM‑Antenna ID"
// @Success      200 {object} responses.SuccessResponseSwagger[dto.EZGsmAntennaGeneralResponse]
// @Failure      400 {object} responses.BadRequestResponseSwagger
// @Failure      500 {object} responses.InternalServerErrorResponseSwagger
// @Router       /v1/gsm-antenna/general [get]
func GetGsmAntennaGeneralById(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("gsm_antenna_id"))
	if err != nil || id <= 0 {
		responses.BadRequest(c, "invalid gsm_antenna_id")
		return
	}

	db := database.GetDB()
	if db == nil {
		responses.InternalServerError(c, "database connection error")
		return
	}

	var m models.GsmAntennaGeneralModel
	if err := db.QueryRow(`SELECT * FROM public.ezw_get_gsm_antenna_general_by_id($1)`, id).Scan(
		&m.GsmAntennaID, &m.SerialNo, &m.StatusID, &m.StatusName, &m.Remark,
	); err != nil {
		if err == sql.ErrNoRows {
			responses.OK(c, nil)
			return
		}
		log.Println("query error:", err)
		responses.InternalServerError(c, "query error")
		return
	}

	resp := dto.EZGsmAntennaGeneralResponse{}
	if m.GsmAntennaID.Valid {
		tmp := int(m.GsmAntennaID.Int64)
		resp.GsmAntennaID = &tmp
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
