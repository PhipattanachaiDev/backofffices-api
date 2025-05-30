package TrakcerControllers

import (
	// "database/sql"
	"database/sql"
	"log"
	"strconv"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/tracker"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/tracker"
	responses "ezview.asia/ezview-web/ezview-lite-back-office/types/responses"

	"github.com/gin-gonic/gin"
)

// SearchTracker รับ payload (EZWSearchTrackerRequest) และเรียกฟังก์ชัน ezw_search_tracker
// @Summary ค้นหา Tracker
// @Description API สำหรับค้นหา Tracker ตามเงื่อนไข
// @Tags Tracker
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param EZWSearchTrackerRequest body dto.EZWSearchTrackerRequest true "Payload สำหรับค้นหา Tracker"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWSearchTrackerModelResponse] "OK"
// @Failure 400 {object} responses.BadRequestResponseSwagger "Bad Request"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "Internal Server Error"
// @Router /v1/tracker/search [post]
func SearchTracker(c *gin.Context) {

	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	var req dto.EZWSearchTrackerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// เรียกฟังก์ชัน ezw_search_tracker
	// อย่าลืม CAST ให้ตรง: tracker_code, brand_id, model_id, status_id, remark
	rows, err := db.Query(`SELECT * FROM public.ezw_search_tracker2($1, $2, $3, $4, $5)`,
		req.TrackerCode,
		req.BrandId,
		req.ModelId,
		req.StatusId,
		req.Remark,
	)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var results []dto.EZWSearchTrackerModelResponse

	for rows.Next() {
		var m models.EZWSearchTrackerModel
		if err := rows.Scan(
			&m.TrackerBomId,
			&m.TrackerId,
			&m.TrackerCode,
			&m.SerialNo,
			&m.TrackerModelId,
			&m.TrackerModelName,
			&m.TrackerBrandId,
			&m.TrackerBrandName,
			&m.StatusId,
			&m.StatusName,
			&m.Remark,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}

		results = append(results, dto.EZWSearchTrackerModelResponse{
			TrackerBomId:     m.TrackerBomId,
			TrackerId:        m.TrackerId,
			TrackerCode:      m.TrackerCode.String,
			SerialNo:         m.SerialNo.String,
			TrackerModelId:   int(m.TrackerModelId.Int64),
			TrackerModelName: m.TrackerModelName.String,
			TrackerBrandId:   int(m.TrackerBrandId.Int64),
			TrackerBrandName: m.TrackerBrandName.String,
			StatusId:         int(m.StatusId.Int64),
			StatusName:       m.StatusName.String,

			Remark: m.Remark.String,
		})
	}

	// เช็ค errors จาก iteration
	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	if len(results) == 0 {
		results = []dto.EZWSearchTrackerModelResponse{}
	}

	responses.OK(c, results)
}

// GetTrackerBrand ดึงข้อมูล Tracker Brand
// @Summary ดึง Tracker Brand
// @Description ดึงข้อมูล Brand (category_id=146) จากระบบ
// @Tags Tracker
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWTrackerBrandResponse] "OK"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "Internal Server Error"
// @Router /v1/tracker/brand [get]
func GetTrackerBrand(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM public.ezw_get_tracker_brand()")
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var brands []dto.EZWTrackerBrandResponse
	for rows.Next() {
		var m models.EZWTrackerBrand
		if err := rows.Scan(&m.BrandId, &m.BrandName); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		brands = append(brands, dto.EZWTrackerBrandResponse{
			BrandId:   m.BrandId,
			BrandName: (m.BrandName.String),
		})
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	if len(brands) == 0 {
		brands = []dto.EZWTrackerBrandResponse{}
	}

	responses.OK(c, brands)
}

// GetTrackerStatus ดึงข้อมูล Tracker Status
// @Summary ดึง Tracker Status
// @Description ดึงข้อมูลสถานะ Tracker (category_id=128) จากระบบ
// @Tags Tracker
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWTrackerStatusResponse] "OK"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "Internal Server Error"
// @Router /v1/tracker/status [get]
func GetTrackerStatus(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM public.ezw_get_tracker_status()")
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var statuses []dto.EZWTrackerStatusResponse
	for rows.Next() {
		var m models.EZWTrackerStatus
		if err := rows.Scan(&m.StatusId, &m.StatusName); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		statuses = append(statuses, dto.EZWTrackerStatusResponse{
			StatusId:   m.StatusId,
			StatusName: (m.StatusName.String),
		})
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	if len(statuses) == 0 {
		statuses = []dto.EZWTrackerStatusResponse{}
	}

	responses.OK(c, statuses)
}

// GetTrackerStatus ดึงข้อมูล Tracker Model โดยรับ Id ของ Tracker Brand
// @Summary ดึง Tracker Model ตาม brand (parent_id)
// @Description ดึงข้อมูล Model ของ Tracker ที่มี parent เป็น brand_id
// @Tags Tracker
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param parent_id query int true "Brand ID"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWTrackerModelResponse] "OK"
// @Failure 400 {object} responses.BadRequestResponseSwagger "Bad Request"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "Internal Server Error"
// @Router /v1/tracker/model [get]
func GetTrackerModel(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// ดึง parent_id จาก query string
	parentIdStr := c.Query("parent_id")
	if parentIdStr == "" {
		responses.BadRequest(c, "parent_id is required")
		return
	}

	parentId, err := strconv.Atoi(parentIdStr)
	if err != nil {
		responses.BadRequest(c, "Invalid parent_id")
		return
	}

	rows, err := db.Query("SELECT * FROM public.ezw_get_tracker_model($1)", parentId)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var modelsResp []dto.EZWTrackerModelResponse
	for rows.Next() {
		var m models.EZWTrackerModel
		if err := rows.Scan(&m.ModelId, &m.ModelName); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		modelsResp = append(modelsResp, dto.EZWTrackerModelResponse{
			ModelId:   m.ModelId,
			ModelName: (m.ModelName.String),
		})
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	if len(modelsResp) == 0 {
		modelsResp = []dto.EZWTrackerModelResponse{}
	}

	responses.OK(c, modelsResp)
}

// UpdateTracker แก้ไขข้อมูล Tracker
// @Summary แก้ไขข้อมูล Tracker
// @Description API นี้ใช้สำหรับแก้ไขข้อมูล Tracker ในระบบ
// @Tags Tracker
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body dto.EZWUpdateTrackerRequest true "Payload สำหรับแก้ไขข้อมูล Tracker"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWUpdateTrackerResponse] "อัปเดตข้อมูล Tracker สำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/tracker/update [put]
func UpdateTracker(c *gin.Context) {
	var req dto.EZWUpdateTrackerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// 🔐 เช็คค่า payload fields ก่อน
	if req.TrackerId <= 0 {
		responses.BadRequest(c, "TrackerId must be a positive integer")
		return
	}
	if req.TrackerCode == "" {
		responses.BadRequest(c, "TrackerCode is required")
		return
	}
	if req.SerialNo == "" {
		responses.BadRequest(c, "SerialNo is required")
		return
	}
	if req.TrackerBomId <= 0 {
		responses.BadRequest(c, "TrackerBomId must be a positive integer")
		return
	}
	if req.BrandId <= 0 {
		responses.BadRequest(c, "BrandId must be a positive integer")
		return
	}
	if req.ModelId <= 0 {
		responses.BadRequest(c, "ModelId must be a positive integer")
		return
	}
	if req.StatusId <= 0 {
		responses.BadRequest(c, "StatusId must be a positive integer")
		return
	}

	// 🔐 ดึง user_id จาก context
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

	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// // ✅ Handle กรณีที่ Remark อาจเป็น NULL หรือว่างเปล่า
	// remark := sql.NullString{String: req.Remark, Valid: req.Remark != ""}

	// ✅ Execute PostgreSQL function
	var resultMessage string
	err := db.QueryRow(`
        SELECT public.ezw_update_tracker(
            $1::smallint, $2::varchar, $3::varchar, $4::integer, 
            $5::integer, $6::integer, $7::integer, $8::integer, $9::varchar
        )
    `,
		req.TrackerId,
		req.TrackerCode,
		req.SerialNo,
		req.TrackerBomId,
		req.BrandId,
		req.ModelId,
		req.StatusId,
		userID,
		req.Remark,
	).Scan(&resultMessage)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	// 🚨 ตรวจสอบ response จากฟังก์ชันของฐานข้อมูล
	if resultMessage != "Update successful" {
		responses.BadRequest(c, resultMessage)
		return
	}

	// ✅ ตอบกลับผลลัพธ์
	responses.OK(c, dto.EZWUpdateTrackerResponse{Message: resultMessage})
}

// InsertTracker ใช้สำหรับ Insert ข้อมูล Tracker ใหม่
// @Summary สร้าง Tracker ใหม่
// @Description API สำหรับ Insert ข้อมูล Tracker
// @Tags Tracker
// @Accept json
// @Produce json
// @Security Bearer
// @Param EZWInsertTrackerRequest body dto.EZWInsertTrackerRequest true "Payload สำหรับ Insert Tracker"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWInsertTrackerResponse] "Insert successful"
// @Failure 400 {object} responses.BadRequestResponseSwagger "Bad Request"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "Internal Server Error"
// @Router /v1/tracker/insert [post]
func InsertTracker(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	var req dto.EZWInsertTrackerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// ดึง user_id จาก context
	userIDVal, exists := c.Get("user_id")
	if !exists {
		responses.Unauthorized(c, "User ID not found in context")
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		responses.Unauthorized(c, "Invalid user ID format")
		return
	}

	// ตรวจสอบค่าใน req ถ้าต้องการ
	if req.TrackerCode == "" {
		responses.BadRequest(c, "tracker_code is required")
		return
	}
	if req.SerialNo == "" {
		responses.BadRequest(c, "serial_no is required")
		return
	}
	if req.TrackerBomId <= 0 {
		responses.BadRequest(c, "tracker_bom_id must be positive")
		return
	}
	if req.BrandId <= 0 {
		responses.BadRequest(c, "brand_id must be positive")
		return
	}
	if req.ModelId <= 0 {
		responses.BadRequest(c, "model_id must be positive")
		return
	}
	if req.StatusId <= 0 {
		responses.BadRequest(c, "status_id must be positive")
		return
	}

	var result string
	err := db.QueryRow(`
        SELECT public.ezw_insert_tracker(
            $1,
            $2,
            $3,
            $4,
            $5,
            $6,
            $7,
            $8
        )
    `,
		req.TrackerCode,
		req.SerialNo,
		req.TrackerBomId,
		req.BrandId,
		req.ModelId,
		req.StatusId,
		req.Remark,
		userID,
	).Scan(&result)

	if err != nil {
		log.Printf("Error executing ezw_insert_tracker: %v", err)
		responses.InternalServerError(c, "Insert failed")
		return
	}

	// ตรวจสอบผลลัพธ์
	if result == "" {
		responses.BadRequest(c, "No result from ezw_insert_tracker")
		return
	}

	// เช่นได้ข้อความ "Insert successful: tracker_id = X"
	responses.OK(c, dto.EZWInsertTrackerResponse{Message: result})
}

// GetTrackerGeneralByTrackerId ดึงข้อมูล Tracker ตาม tracker_id
// @Summary ดึงข้อมูล Tracker ตาม tracker_id
// @Description เรียกฟังก์ชัน ezw_get_tracker_general_by_tracker_id
// @Tags Tracker
// @Accept json
// @Produce json
// @Security Bearer
// @Param tracker_id query int true "Tracker ID"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWResponseGetTrackerGeneral] "OK"
// @Failure 400 {object} responses.BadRequestResponseSwagger "Bad Request"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "Internal Server Error"
// @Router /v1/tracker/general [get]
func GetTrackerGeneralByTrackerId(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// ดึง tracker_id จาก query param
	trackerIdStr := c.Query("tracker_id")
	if trackerIdStr == "" {
		responses.BadRequest(c, "tracker_id is required")
		return
	}

	trackerId, err := strconv.Atoi(trackerIdStr)
	if err != nil || trackerId <= 0 {
		responses.BadRequest(c, "invalid tracker_id")
		return
	}

	row := db.QueryRow(`SELECT * FROM public.ezw_get_tracker_general_by_tracker_id($1::smallint)`, trackerId)

	var tg models.EZWGetTrackerGeneral
	err = row.Scan(
		&tg.TrackerBomId,
		&tg.TrackerCode,
		&tg.SerialNo,
		&tg.TrackerBrandId,
		&tg.TrackerBrandName,
		&tg.TrackerModelId,
		&tg.TrackerModelName,
		&tg.StatusId,
		&tg.StatusName,
		&tg.Remark,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.NotFound(c, "Tracker not found")
			return
		}
		log.Printf("Error scanning row: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	response := dto.EZWResponseGetTrackerGeneral{}

	if tg.TrackerBomId.Valid {
		val := int(tg.TrackerBomId.Int64)
		response.TrackerBomId = &val
	}
	if tg.TrackerCode.Valid {
		response.TrackerCode = &tg.TrackerCode.String
	}
	if tg.SerialNo.Valid {
		response.SerialNo = &tg.SerialNo.String
	}
	if tg.TrackerBrandId.Valid {
		val := int(tg.TrackerBrandId.Int64)
		response.TrackerBrandId = &val
	}
	if tg.TrackerBrandName.Valid {
		response.TrackerBrandName = &tg.TrackerBrandName.String
	}
	if tg.TrackerModelId.Valid {
		val := int(tg.TrackerModelId.Int64)
		response.TrackerModelId = &val
	}
	if tg.TrackerModelName.Valid {
		response.TrackerModelName = &tg.TrackerModelName.String
	}
	if tg.StatusId.Valid {
		val := int(tg.StatusId.Int64)
		response.StatusId = &val
	}
	if tg.StatusName.Valid {
		response.StatusName = &tg.StatusName.String
	}
	if tg.Remark.Valid {
		response.Remark = &tg.Remark.String
	}

	responses.OK(c, response)
}
