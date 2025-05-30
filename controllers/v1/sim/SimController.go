package SimControllers

import (
	"database/sql"
	"log"
	"strconv"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/sim"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/sim"
	responses "ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	"github.com/gin-gonic/gin"
)

// GetSimStatus ดึงข้อมูล Sim Status
// @Summary ดึงข้อมูล Sim Status
// @Description API นี้ใช้สำหรับดึงข้อมูลสถานะซิมการ์ดจากฐานข้อมูล
// @Tags Sim
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWSimStatusModelResponse] "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/sim/status [get]
func GetSimStatus(c *gin.Context) {
	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM public.ezw_get_sim_status()")
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var status []dto.EZWSimStatusModelResponse

	// เพิ่มแถวพิเศษ operator_id=0 ก่อน หรือหลัง ก็ได้
	status = append(status, dto.EZWSimStatusModelResponse{
		StatusId:   0,
		StatusName: "ทั้งหมด",
	})
	for rows.Next() {
		var sim models.EZWGetSimStatus
		if err := rows.Scan(&sim.StatusId, &sim.StatusName); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		status = append(status, dto.EZWSimStatusModelResponse{
			StatusId:   sim.StatusId,
			StatusName: sim.StatusName,
		})
	}

	responses.OK(c, status)
}

// SearchSim ค้นหา SIM ตามเงื่อนไขพารามิเตอร์
// @Summary ค้นหา SIM
// @Description รับ JSON และส่งพารามิเตอร์ไปยังฟังก์ชัน ezw_search_sim
// @Tags Sim
// @Accept json
// @Produce json
// @Security Bearer
// @Param searchRequest body dto.EZWSearchSimRequest true "Payload สำหรับค้นหา SIM"
// @Success 201 {object} responses.SuccessResponseSwagger[dto.EZWSearchSimResponse] "ผลลัพธ์การค้นหา SIM"
// @Failure 400 {object} responses.BadRequestResponseSwagger "ข้อมูลที่ส่งมาไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.Response "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/sim/search [post]
func SearchSim(c *gin.Context) {
	var reqBody dto.EZWSearchSimRequest
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
		simNoNull      sql.NullString
		operatorIDNull sql.NullInt32
		statusIDNull   sql.NullInt32
		remarkNull     sql.NullString
	)

	// ถ้า string ไม่ว่าง
	if reqBody.SimNo != "" {
		simNoNull = sql.NullString{String: reqBody.SimNo, Valid: true}
	}

	// ถ้าเป็น 0 → ถือว่า NULL, แต่ถ้าเป็น 99 → ส่งเป็น valid เพื่อให้ PostgreSQL มองว่า “เลือกทั้งหมด”
	if reqBody.OperatorId != 0 {
		operatorIDNull = sql.NullInt32{Int32: int32(reqBody.OperatorId), Valid: true}
	}
	if reqBody.StatusId != 0 {
		statusIDNull = sql.NullInt32{Int32: int32(reqBody.StatusId), Valid: true}
	}
	if reqBody.Remark != "" {
		remarkNull = sql.NullString{String: reqBody.Remark, Valid: true}
	}

	query := `
		SELECT * FROM public.ezw_search_sim(
			$1::varchar,
			$2::int,
			$3::int,
			$4::varchar
		)
	`
	rows, err := db.Query(query,
		simNoNull,
		operatorIDNull,
		statusIDNull,
		remarkNull,
	)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var results []dto.EZWSearchSimResponse
	for rows.Next() {
		var s models.EZWSearchSim
		if err := rows.Scan(
			&s.SimId,        // column 1: sim_id
			&s.SimNo,        // column 2: sim_no
			&s.OperatorId,   // column 3: operator_id
			&s.OperatorCode, // column 4: operator_code
			&s.OperatorName, // column 5: operator_name
			&s.StatusId,     // column 6: status_id
			&s.StatusCode,   // column 7: status_code
			&s.StatusName,   // column 8: status_name
			&s.Remark,       // column 9: remark

		); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		results = append(results, dto.EZWSearchSimResponse{
			SimId:        s.SimId,
			SimNo:        s.SimNo,
			OperatorId:   s.OperatorId,
			OperatorCode: s.OperatorCode,
			OperatorName: s.OperatorName,
			StatusId:     s.StatusId,
			StatusCode:   s.StatusCode,
			StatusName:   s.StatusName,
			Remark:       s.Remark,
		})
	}

	if len(results) == 0 {
		responses.OK(c, []dto.EZWSearchSimResponse{})
		return
	}

	responses.OK(c, results)

}

// InsertSim เพิ่มข้อมูล SIM ใหม่
// @Summary เพิ่มข้อมูล SIM
// @Description เพิ่ม SIM ใหม่ลงฐานข้อมูล
// @Tags Sim
// @Accept json
// @Produce json
// @Security Bearer
// @Param sim body dto.EZWInsertSimRequest true "ข้อมูล SIM ที่ต้องการเพิ่ม"
// @Success 201 {object} responses.SuccessResponseSwagger[dto.EZWInsertSimRequest] "เพิ่มข้อมูล SIM สำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "ข้อมูลที่ส่งมาไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/sim/insert [post]
func InsertSim(c *gin.Context) {
	var req dto.EZWInsertSimRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// 1) Check payload fields manually
	if req.SimNo == "" {
		responses.BadRequest(c, "SimNo is required")
		return
	}
	if req.OperatorId <= 0 {
		responses.BadRequest(c, "OperatorId must be positive integer")
		return
	}
	if req.StatusId <= 0 {
		responses.BadRequest(c, "StatusId must be positive integer")
		return
	}

	// 2) ดึง user_id จาก context
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

	// 3) เรียก db
	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// 4) เรียกฟังก์ชัน SQL
	var result string
	err := db.QueryRow(`
		SELECT public.ezw_insert_sim($1, $2, $3, $4, $5)
	`,
		req.SimNo,
		req.OperatorId,
		req.StatusId,
		req.Remark,
		userID,
	).Scan(&result)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	// 5) ตอบกลับ
	responses.OK(c, dto.EZWInsertSimResponse{Message: result})
}

// UpdateSim แก้ไขข้อมูล SIM
// @Summary แก้ไขข้อมูล SIM
// @Description อัปเดตข้อมูล SIM ที่มีอยู่ในฐานข้อมูล
// @Tags Sim
// @Accept json
// @Produce json
// @Security Bearer
// @Param sim body dto.EZWUpdateSimRequest true "ข้อมูล SIM ที่ต้องการอัปเดต"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWUpdateSimResponse] "อัปเดตข้อมูล SIM สำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "ข้อมูลที่ส่งมาไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/sim/update [put]
func UpdateSim(c *gin.Context) {
	var req dto.EZWUpdateSimRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// 1) Check payload fields
	if req.SimId <= 0 {
		responses.BadRequest(c, "SimId must be positive integer")
		return
	}
	if req.SimNo == "" {
		responses.BadRequest(c, "SimNo is required")
		return
	}
	if req.OperatorId <= 0 {
		responses.BadRequest(c, "OperatorId must be positive integer")
		return
	}
	if req.StatusId <= 0 {
		responses.BadRequest(c, "StatusId must be positive integer")
		return
	}

	// 2) ดึง user_id จาก context
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

	// 3) เรียก db
	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// 4) เรียกฟังก์ชัน SQL
	var result string
	err := db.QueryRow(`
		SELECT public.ezw_update_sim($1, $2, $3, $4, $5, $6)
	`,
		req.SimId,
		req.SimNo,
		req.OperatorId,
		req.StatusId,
		req.Remark,
		userID,
	).Scan(&result)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	// 5) ตอบกลับ
	responses.OK(c, dto.EZWUpdateSimResponse{Message: result})
}

// GetSimOperator ดึงข้อมูล Sim Operator
// @Summary ดึงข้อมูล Sim Operator
// @Description API นี้ใช้สำหรับดึงข้อมูลสถานะซิมการ์ดจากฐานข้อมูล
// @Tags Sim
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWSimOperatorModelResponse] "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/sim/operator [get]
func GetSimOperator(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM public.ezw_get_sim_operator()")
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var operators []dto.EZWSimOperatorModelResponse

	// เพิ่มแถวพิเศษ operator_id=0 ก่อน หรือหลัง ก็ได้
	operators = append(operators, dto.EZWSimOperatorModelResponse{
		OperatorId:   0,
		OperatorName: "ทั้งหมด",
	})

	for rows.Next() {
		var sim models.EZWGetSimOperator
		if err := rows.Scan(&sim.OperatorId, &sim.OperatorName); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		operators = append(operators, dto.EZWSimOperatorModelResponse{
			OperatorId:   sim.OperatorId,
			OperatorName: sim.OperatorName,
		})
	}

	responses.OK(c, operators)
}

// GetSimGeneralById godoc
// @Summary      Get SIM general information
// @Tags         Sim
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        sim_id query int true "SIM ID"
// @Success      200 {object} responses.SuccessResponseSwagger[dto.EZWGetSimGeneralResponse]
// @Failure      400 {object} responses.BadRequestResponseSwagger
// @Failure      500 {object} responses.InternalServerErrorResponseSwagger
// @Router       /v1/sim/general [get]
func GetSimGeneralById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("sim_id"))
	if err != nil || id <= 0 {
		responses.BadRequest(c, "invalid sim_id")
		return
	}

	db := database.GetDB()
	if db == nil {
		responses.InternalServerError(c, "database connection error")
		return
	}

	var m models.SimGeneralModel
	if err := db.QueryRow(`SELECT * FROM public.ezw_get_sim_general_by_id($1)`, id).Scan(
		&m.SimID, &m.SimNo, &m.OperatorID, &m.OperatorName,
		&m.StatusID, &m.StatusName, &m.Remark,
	); err != nil {
		if err == sql.ErrNoRows {
			responses.OK(c, nil)
			return
		}
		log.Println("query error:", err)
		responses.InternalServerError(c, "query error")
		return
	}

	resp := dto.EZWGetSimGeneralResponse{}
	if m.SimID.Valid {
		tmp := int(m.SimID.Int64)
		resp.SimID = &tmp
	}
	if m.SimNo.Valid {
		tmp := m.SimNo.String
		resp.SimNo = &tmp
	}
	if m.OperatorID.Valid {
		tmp := int(m.OperatorID.Int64)
		resp.OperatorID = &tmp
	}
	if m.OperatorName.Valid {
		tmp := m.OperatorName.String
		resp.OperatorName = &tmp
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
