package GpsControllers

import (
	"database/sql"
	"log"
	"reflect"
	"strconv"
	"strings"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/gps"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/gps"
	"ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	"github.com/gin-gonic/gin"
)

// GetGpsStatus ดึงข้อมูลสถานะ GPS
// @Summary ดึงข้อมูลสถานะ GPS ทั้งหมด
// @Description API นี้ใช้สำหรับดึงข้อมูลสถานะ GPS ทั้งหมด
// @Accept json
// @Tags GPS
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.GPSStatusResponse] "การร้องขอข้อมูลสถานะ GPS สำเร็จ"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "Token ไม่ถูกต้องหรือหมดอายุ"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/gps-status [get]
func GetGpsStatus(c *gin.Context) {
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	query := `SELECT * FROM ezw_get_gpsstatus()`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch GPS data")
		return
	}
	defer rows.Close()

	var gpsstatus []dto.GPSStatusResponse

	gpsstatus = append(gpsstatus, dto.GPSStatusResponse{
		TypeID:   0,
		TypeName: "ทั้งหมด",
	})

	for rows.Next() {
		var gpsStatus models.GPSStatus
		if err := rows.Scan(&gpsStatus.TypeID, &gpsStatus.TypeName); err != nil {
			log.Println("Row scan error:", err)
			responses.InternalServerError(c, "Error processing data")
			return
		}

		gpsstatus = append(gpsstatus, dto.GPSStatusResponse{
			TypeID:   gpsStatus.TypeID,
			TypeName: gpsStatus.TypeName,
		})
	}

	if err := rows.Err(); err != nil {
		log.Println("Rows iteration error:", err)
		responses.InternalServerError(c, "Error retrieving GPS data")
		return
	}

	if len(gpsstatus) == 0 {
		responses.OK(c, []dto.GPSStatusResponse{})
		return
	}

	responses.OK(c, gpsstatus)
}

// GetGpsBrands ดึงข้อมูลยี่ห้อ GPS
// @Summary ดึงข้อมูลยี่ห้อ GPS ทั้งหมด
// @Description API นี้ใช้สำหรับดึงข้อมูลยี่ห้อ GPS ทั้งหมด
// @Accept json
// @Tags DropdownsGPS
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.GPSBrandResponse] "การร้องขอข้อมูลยี่ห้อ GPS สำเร็จ"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "Token ไม่ถูกต้องหรือหมดอายุ"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/gps-brands [get]
func GetGpsBrands(c *gin.Context) {
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	categoryID := 136

	query := `SELECT * FROM ezw_get_gpsbrands($1)`
	rows, err := db.Query(query, categoryID)
	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch GPS data")
		return
	}
	defer rows.Close()

	var gpsbrans []dto.GPSBrandResponse

	gpsbrans = append(gpsbrans, dto.GPSBrandResponse{
		TypeID:   0,
		TypeName: "ทั้งหมด",
	})

	for rows.Next() {
		var gpsBrand models.GPSBrand
		if err := rows.Scan(&gpsBrand.TypeID, &gpsBrand.TypeName); err != nil {
			log.Println("Row scan error:", err)
			responses.InternalServerError(c, "Error processing data")
			return
		}

		gpsbrans = append(gpsbrans, dto.GPSBrandResponse{
			TypeID:   gpsBrand.TypeID,
			TypeName: gpsBrand.TypeName,
		})
	}

	if len(gpsbrans) == 0 {
		responses.OK(c, []dto.GPSBrandResponse{})
		return
	}

	responses.OK(c, gpsbrans)

}

// GetGpsModels ดึงข้อมูลรุ่น GPS
// @Summary ดึงข้อมูลรุ่น GPS ตามยี่ห้อที่เลือก
// @Description API นี้ใช้สำหรับดึงข้อมูลรุ่น GPS ตามยี่ห้อที่กำหนด
// @Accept json
// @Tags DropdownsGPS
// @Produce json
// @Param type_id query int true "รหัสยี่ห้อ GPS"
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.GPSModelResponse] "การร้องขอข้อมูลรุ่น GPS สำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "คำขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "Token ไม่ถูกต้องหรือหมดอายุ"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/gps-models [get]
func GetGpsModels(c *gin.Context) {
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}
	TypeIDStr := c.DefaultQuery("type_id", "")
	if TypeIDStr == "" {
		responses.BadRequest(c, "Missing type_id parameter")
		return
	}

	TypeID, err := strconv.Atoi(TypeIDStr)
	if err != nil {
		responses.BadRequest(c, "Invalid type_id parameter. It must be a number")
		return
	}

	query := `SELECT * FROM ezw_get_gpsmodels($1)`
	rows, err := db.Query(query, TypeID)
	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch GPS data")
		return
	}
	defer rows.Close()

	var gpsmodels []dto.GPSModelResponse

	gpsmodels = append(gpsmodels, dto.GPSModelResponse{
		TypeID:   0,
		TypeName: "ทั้งหมด",
	})

	for rows.Next() {
		var gpsModel models.GPSModel
		if err := rows.Scan(&gpsModel.TypeID, &gpsModel.TypeName); err != nil {
			log.Println("Row scan error:", err)
			responses.InternalServerError(c, "Error processing data")
			return
		}

		gpsmodels = append(gpsmodels, dto.GPSModelResponse{
			TypeID:   gpsModel.TypeID,
			TypeName: gpsModel.TypeName,
		})
	}

	if len(gpsmodels) == 0 {
		responses.OK(c, []dto.GPSModelResponse{})
		return
	}

	responses.OK(c, gpsmodels)

}

// SearchGPS ค้นหาข้อมูล GPS
// @Summary ค้นหาข้อมูล GPS ตามเงื่อนไขที่กำหนด
// @Description API นี้ใช้สำหรับค้นหาข้อมูล GPS ตามเงื่อนไขที่กำหนด
// @Description
// @Description - **s_gps_imei**: ชื่อสินค้า ของ GPS
// @Description - **s_serial_no**: หมายเลขรหัสสินค้า ของ GPS
// @Description - **s_brand_id**:  ID ยี่ห้อ GPS
// @Description - **s_model_id**: ID รุ่น GPS
// @Description - **s_status_id**: ID สถานะ GPS
// @Description - **s_remark**: รายละเอียด
// @Accept json
// @Tags GPS
// @Produce json
// @Security Bearer
// @Param searchRequest body dto.GPSRequest true "Search Request"
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.GPSResponse] "การร้องขอข้อมูล GPS สำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "คำขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "Token ไม่ถูกต้องหรือหมดอายุ"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/search-gps [post]
func SearchGPS(c *gin.Context) {
	var reqBody dto.GPSRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Println("JSON Bind Error:", err)
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// ตรวจสอบค่า GpsIMEI และ SerialNumber (string)
	if reflect.TypeOf(reqBody.GpsIMEI).Kind() != reflect.String {
		responses.BadRequest(c, "GpsIMEI must be a string")
		return
	}
	if reflect.TypeOf(reqBody.SerialNumber).Kind() != reflect.String {
		responses.BadRequest(c, "SerialNumber must be a string")
		return
	}

	// ตรวจสอบค่า BrandID, ModelID, StatusID (int)
	if reflect.TypeOf(reqBody.BrandID).Kind() != reflect.Int {
		responses.BadRequest(c, "BrandID must be an integer")
		return
	}
	if reflect.TypeOf(reqBody.ModelID).Kind() != reflect.Int {
		responses.BadRequest(c, "ModelID must be an integer")
		return
	}
	if reflect.TypeOf(reqBody.StatusID).Kind() != reflect.Int {
		responses.BadRequest(c, "StatusID must be an integer")
		return
	}

	// ตรวจสอบค่า Remark (string)
	if reflect.TypeOf(reqBody.Remark).Kind() != reflect.String {
		responses.BadRequest(c, "Remark must be a string")
		return
	}

	Gpsimei := reqBody.GpsIMEI
	SerialNumber := reqBody.SerialNumber
	BrandID := reqBody.BrandID
	ModelID := reqBody.ModelID
	StatusID := reqBody.StatusID
	Remark := reqBody.Remark

	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	query := `SELECT * FROM ezw_search_gps($1, $2, $3, $4, $5, $6)`
	rows, err := db.Query(query, Gpsimei, SerialNumber, BrandID, ModelID, StatusID, Remark)
	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch Search GPS data")
		return
	}
	defer rows.Close()

	var gpsList []dto.GPSResponse

	for rows.Next() {
		var gpsData models.SearchGPS
		if err := rows.Scan(
			&gpsData.GpsID,
			&gpsData.GpsIMEI,
			&gpsData.SerialNumber,
			&gpsData.StatusID,
			&gpsData.StatusName,
			&gpsData.Remark,
			&gpsData.BrandName,
			&gpsData.BrandCode,
			&gpsData.ModelName,
			&gpsData.ModelCode,
		); err != nil {
			log.Println("Row scan error:", err)
			responses.InternalServerError(c, "Error processing Search data")
			return
		}

		gpsList = append(gpsList, dto.GPSResponse{
			GpsID:        gpsData.GpsID,
			GpsIMEI:      gpsData.GpsIMEI,
			SerialNumber: gpsData.SerialNumber,
			StatusID:     gpsData.StatusID,
			StatusName:   gpsData.StatusName,
			Remark:       gpsData.Remark,
			BrandName:    gpsData.BrandName,
			BrandCode:    gpsData.BrandCode,
			ModelName:    gpsData.ModelName,
			ModelCode:    gpsData.ModelCode,
		})
	}

	if len(gpsList) == 0 {
		responses.OK(c, []dto.GPSResponse{})
		return
	}

	responses.OK(c, gpsList)
}

// UpdateGPS อัพเดทข้อมูล GPS
// @Summary อัพเดทข้อมูล GPS ตามเงื่อนไขที่กำหนด
// @Description API นี้ใช้สำหรับอัพเดทข้อมูล GPS ตามเงื่อนไขที่กำหนด
// @Description
// @Description - **gps_id**: รหัส GPS
// @Description - **gps_imei**: หมายเลข IMEI ของ GPS
// @Description - **serial_no**: หมายเลข Serial ของ GPS
// @Description - **brand_id**: รหัสยี่ห้อ GPS
// @Description - **model_id**: รหัสรุ่น GPS
// @Description - **status_id**: รหัสสถานะ GPS
// @Description - **remark**: รายละเอียด
// @Accept json
// @Tags GPS
// @Produce json
// @Security Bearer
// @Param updateRequest body dto.UpdateGPSRequest true "Update Request"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.UpdateGPSResponse] "การร้องขออัพเดทข้อมูล GPS สำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "คำขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "Token ไม่ถูกต้องหรือหมดอายุ"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/update-gps [put]
func UpdateGPS(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		responses.Unauthorized(c, "User not authenticated")
		return
	}

	var req dto.UpdateGPSRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("JSON Bind Error:", err)
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// ตรวจสอบค่า GpsID
	if req.GpsID <= 0 {
		responses.BadRequest(c, "GpsID must be a valid integer greater than 0")
		return
	}

	// ตรวจสอบค่า GpsIMEI และ SerialNumber ว่าห้ามเป็นค่าว่าง
	if strings.TrimSpace(req.GpsIMEI) == "" {
		responses.BadRequest(c, "GpsIMEI is required and cannot be empty")
		return
	}
	if strings.TrimSpace(req.SerialNumber) == "" {
		responses.BadRequest(c, "SerialNumber is required and cannot be empty")
		return
	}

	if req.BrandID <= 0 {
		responses.BadRequest(c, "Invalid BrandID")
		return
	}
	if req.ModelID <= 0 {
		responses.BadRequest(c, "Invalid ModelID")
		return
	}
	if req.StatusID <= 0 {
		responses.BadRequest(c, "Invalid StatusID")
		return
	}

	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	var result string
	query := "SELECT * FROM public.ezw_update_gps($1, $2, $3, $4, $5, $6, $7, $8)"
	err := db.QueryRow(query, req.GpsID, req.GpsIMEI, req.SerialNumber, req.BrandID, req.ModelID, req.StatusID, req.Remark, userID).Scan(&result)
	if err != nil {
		log.Println("Database query error:", err)
		log.Println("Query result:", result)
		responses.InternalServerError(c, "Failed to Update GPS data")
		return
	}

	log.Println("Query result:", result)

	if strings.Contains(result, "Error: ") {
		errorMessage := strings.TrimPrefix(result, "Error: ")
		responses.BadRequest(c, errorMessage)
		return
	}

	// Alternatively, if you want to catch any errors from RAISE EXCEPTION messages:
	if strings.Contains(result, "not valid") {
		responses.BadRequest(c, result)
		return
	}

	// Response for success
	responseData := dto.UpdateGPSResponse{
		Message: "Update GPS data successfully",
	}

	responses.OK(c, responseData)
}

// CreateGPS สร้างข้อมูล GPS
// @Summary สร้างข้อมูล GPS ตามเงื่อนไขที่กำหนด
// @Description API นี้ใช้สำหรับสร้างข้อมูล GPS ตามเงื่อนไขที่กำหนด
// @Description
// @Description - **gps_imei**: หมายเลข IMEI ของ GPS
// @Description - **serial_no**: หมายเลข Serial ของ GPS
// @Description - **brand_id**: รหัสยี่ห้อ GPS
// @Description - **model_id**: รหัสรุ่น GPS
// @Description - **status_id**: รหัสสถานะ GPS
// @Description - **remark**: รายละเอียด
// @Accept json
// @Tags GPS
// @Produce json
// @Security Bearer
// @Param createRequest body dto.CreateGPSRequest true "Create GPS Request"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.CreateGPSResponse] "การร้องขอสร้างข้อมูล GPS สำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "คำขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "Token ไม่ถูกต้องหรือหมดอายุ"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/create-gps [post]
func CreateGPS(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		responses.Unauthorized(c, "User not authenticated")
		return
	}

	var req dto.CreateGPSRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// ตรวจสอบค่า GpsIMEI และ SerialNumber ว่าห้ามเป็นค่าว่าง
	if strings.TrimSpace(req.GpsIMEI) == "" {
		responses.BadRequest(c, "GpsIMEI is required and cannot be empty")
		return
	}
	if strings.TrimSpace(req.SerialNumber) == "" {
		responses.BadRequest(c, "SerialNumber is required and cannot be empty")
		return
	}

	if req.BrandID <= 0 {
		responses.BadRequest(c, "Invalid BrandID")
		return
	}
	if req.ModelID <= 0 {
		responses.BadRequest(c, "Invalid ModelID")
		return
	}
	if req.StatusID <= 0 {
		responses.BadRequest(c, "Invalid StatusID")
		return
	}

	// ตรวจสอบค่า Remark (string)
	if reflect.TypeOf(req.Remark).Kind() != reflect.String {
		responses.BadRequest(c, "Remark must be a string")
		return
	}

	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	var result string
	query := "SELECT * FROM public.ezw_create_gps($1, $2, $3, $4, $5, $6, $7)"
	err := db.QueryRow(query, req.GpsIMEI, req.SerialNumber, req.BrandID, req.ModelID, req.StatusID, req.Remark, userID).Scan(&result)
	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to create GPS data")
		return
	}

	// ตรวจสอบผลลัพธ์ที่ได้จากฟังก์ชัน sql
	if strings.HasPrefix(result, "Error: ") {
		errorMessage := strings.TrimPrefix(result, "Error: ")
		responses.BadRequest(c, errorMessage)
		return
	}

	responsesData := dto.CreateGPSResponse{
		Message: "Create GPS data successfully",
	}

	responses.OK(c, responsesData)

}

// GetGpsGeneralByGpsId ดึงข้อมูลเบื้องต้นของ GPS ตาม gps_id
// @Summary ดึงข้อมูล GPS ตาม gps_id
// @Description เรียกฟังก์ชัน ezw_get_gps_general_by_gps_id
// @Tags GPS
// @Accept json
// @Produce json
// @Security Bearer
// @Param gps_id query int true "GPS ID"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWGetGpsGeneralByGpsIdResponse] "OK"
// @Failure 400 {object} responses.BadRequestResponseSwagger "Bad Request"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "Internal Server Error"
// @Router /v1/gps/general [get]
func GetGpsGeneralByGpsId(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// รับค่า gps_id จาก Query Param
	gpsIdStr := c.Query("gps_id")
	if gpsIdStr == "" {
		responses.BadRequest(c, "gps_id is required")
		return
	}
	gpsIdInt, err := strconv.Atoi(gpsIdStr)
	if err != nil || gpsIdInt <= 0 {
		responses.BadRequest(c, "invalid gps_id")
		return
	}

	// เรียกฟังก์ชัน PostgreSQL
	row := db.QueryRow(`SELECT * FROM public.ezw_get_gps_general_by_gps_id($1)`, gpsIdInt)

	var m models.EZWGetGpsGeneralModel
	err = row.Scan(
		&m.GpsImei,
		&m.SerialNo,
		&m.StatusId,
		&m.StatusName,
		&m.BrandId,
		&m.BrandName,
		&m.ModelId,
		&m.ModelName,
		&m.Remark,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.OK(c, nil) // หรือส่ง [] ก็ได้
			return
		}
		log.Printf("Error scanning row: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	// แปลงเป็น DTO (pointer) เพื่อส่งค่า null -> null
	res := dto.EZWGetGpsGeneralByGpsIdResponse{}

	if m.GpsImei.Valid {
		temp := m.GpsImei.String
		res.GpsImei = &temp
	}
	if m.SerialNo.Valid {
		temp := m.SerialNo.String
		res.SerialNo = &temp
	}
	if m.StatusId.Valid {
		temp := int(m.StatusId.Int64)
		res.StatusId = &temp
	}
	if m.StatusName.Valid {
		temp := m.StatusName.String
		res.StatusName = &temp
	}
	if m.BrandId.Valid {
		temp := int(m.BrandId.Int64)
		res.BrandId = &temp
	}
	if m.BrandName.Valid {
		temp := m.BrandName.String
		res.BrandName = &temp
	}
	if m.ModelId.Valid {
		temp := int(m.ModelId.Int64)
		res.ModelId = &temp
	}
	if m.ModelName.Valid {
		temp := m.ModelName.String
		res.ModelName = &temp
	}
	if m.Remark.Valid {
		temp := m.Remark.String
		res.Remark = &temp
	}

	responses.OK(c, res)
}
