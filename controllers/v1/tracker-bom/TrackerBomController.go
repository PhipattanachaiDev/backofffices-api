package TrakcerBomControllers

import (
	"database/sql"
	"log"
	"strconv"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/tracker-bom"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/tracker-bom"
	responses "ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	"github.com/gin-gonic/gin"
)

// GetTrackerBom ดึงข้อมูล Tracker BOM
// @Summary ดึงข้อมูล Tracker BOM
// @Description API นี้ใช้สำหรับดึงข้อมูล Tracker BOM จากฐานข้อมูล
// @Tags Tracker Bom
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWGetTrackerBomResponse] "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/tracker-bom/master [get]
func GetTrackerBom(c *gin.Context) {
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM public.ezw_get_tracker_bom()")
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var trackerBoms []dto.EZWGetTrackerBomResponse

	for rows.Next() {
		var bom models.EZWGetTrackerBom
		if err := rows.Scan(
			&bom.TrackerBomId,
			&bom.TrackerId,
			&bom.TrackerCode,
			&bom.GpsId,
			&bom.GpsBrandName,
			&bom.GpsModelName,
			&bom.GpsAntennaId,
			&bom.GpsAntennaSerialNo,
			&bom.GsmAntennaId,
			&bom.GsmAntennaSerialNo,
			&bom.CardReaderId,
			&bom.CardReaderSerialNo,
			&bom.CardReaderBrandName,
			&bom.CardReaderModelName,
			&bom.SimId,
			&bom.SimNo,
			&bom.BatteryId,
			&bom.BatterySerialNo,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}

		trackerBoms = append(trackerBoms, dto.EZWGetTrackerBomResponse{
			TrackerBomId:        bom.TrackerBomId,
			TrackerId:           int(bom.TrackerId.Int64),
			TrackerCode:         bom.TrackerCode.String,
			GpsId:               int(bom.GpsId.Int64),
			GpsBrandName:        bom.GpsBrandName.String,
			GpsModelName:        bom.GpsModelName.String,
			GpsAntennaId:        int(bom.GpsAntennaId.Int64),
			GpsAntennaSerialNo:  bom.GpsAntennaSerialNo.String,
			GsmAntennaId:        int(bom.GsmAntennaId.Int64),
			GsmAntennaSerialNo:  bom.GsmAntennaSerialNo.String,
			CardReaderId:        int(bom.CardReaderId.Int64),
			CardReaderSerialNo:  bom.CardReaderSerialNo.String,
			CardReaderBrandName: bom.CardReaderBrandName.String,
			CardReaderModelName: bom.CardReaderModelName.String,
			SimId:               int(bom.SimId.Int64),
			SimNo:               bom.SimNo.String,
			BatteryId:           int(bom.BatteryId.Int64),
			BatterySerialNo:     bom.BatterySerialNo.String,
		})
	}

	if len(trackerBoms) == 0 {
		trackerBoms = []dto.EZWGetTrackerBomResponse{}
	}

	responses.OK(c, trackerBoms)
}

// InsertTrackerBom บันทึกข้อมูล Tracker BOM
// @Summary บันทึกข้อมูล Tracker BOM
// @Description API นี้ใช้สำหรับ Insert ข้อมูล Tracker BOM
// @Tags Tracker Bom
// @Accept json
// @Produce json
// @Security Bearer
// @Param payload body dto.EZWInsertTrackerBomRequest true "Payload สำหรับบันทึก Tracker BOM"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWInsertTrackerBomResponse] "บันทึกข้อมูลสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/tracker-bom/insert [post]
func InsertTrackerBom(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	var req dto.EZWInsertTrackerBomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON payload")
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
		responses.Unauthorized(c, "Invalid User ID format")
		return
	}

	var resultMessage string
	query := `SELECT public.ezw_insert_tracker_bom($1, $2, $3, $4, $5, $6, $7)`

	err := db.QueryRow(query,
		req.GpsId,
		req.GpsAntennaId,
		req.GsmAntennaId,
		req.CardReaderId,
		req.SimId,
		req.BatteryId,
		userID,
	).Scan(&resultMessage)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	if resultMessage[:5] == "Error" {
		responses.BadRequest(c, resultMessage)
		return
	}

	responses.OK(c, dto.EZWInsertTrackerBomResponse{
		Message: resultMessage,
	})
}

// UpdateTrackerBom อัปเดต Tracker BOM
// @Summary แก้ไขข้อมูล Tracker BOM
// @Description API สำหรับ Update ข้อมูล Tracker BOM
// @Tags Tracker Bom
// @Accept json
// @Produce json
// @Security Bearer
// @Param TrackerBom body dto.EZWUpdateTrackerBomRequest true "Tracker BOM Payload"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWUpdateTrackerBomResponse]
// @Failure 400 {object} responses.BadRequestResponseSwagger
// @Failure 401 {object} responses.UnauthorizedResponseSwagger
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/tracker-bom/update [put]
func UpdateTrackerBom(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("❌ Database connection error")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	var req dto.EZWUpdateTrackerBomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body: "+err.Error())
		return
	}

	// 🔍 ตรวจสอบค่าที่เป็น pointer ว่า nil หรือไม่
	if req.TrackerBomId == nil || *req.TrackerBomId <= 0 {
		responses.BadRequest(c, "tracker_bom_id is required and must be a positive integer")
		return
	}
	if req.GpsId == nil || *req.GpsId <= 0 {
		responses.BadRequest(c, "gps_id is required and must be a positive integer")
		return
	}
	if req.GpsAntennaId == nil || *req.GpsAntennaId <= 0 {
		responses.BadRequest(c, "gps_antenna_id is required and must be a positive integer")
		return
	}
	if req.GsmAntennaId == nil || *req.GsmAntennaId <= 0 {
		responses.BadRequest(c, "gsm_antenna_id is required and must be a positive integer")
		return
	}
	if req.SimId == nil || *req.SimId <= 0 {
		responses.BadRequest(c, "sim_id is required and must be a positive integer")
		return
	}
	if req.BatteryId == nil || *req.BatteryId <= 0 {
		responses.BadRequest(c, "battery_id is required and must be a positive integer")
		return
	}

	// ✅ ตรวจสอบ JWT context เพื่อดึง user_id
	userIdVal, exists := c.Get("user_id")
	if !exists {
		responses.Unauthorized(c, "User ID not found in context")
		return
	}
	userId, ok := userIdVal.(int)
	if !ok {
		responses.Unauthorized(c, "Invalid User ID format")
		return
	}

	// 🔁 ดึงผลลัพธ์จากฟังก์ชันใน PostgreSQL
	var result string
	err := db.QueryRow(`
		SELECT public.ezw_update_tracker_bom($1, $2, $3, $4, $5, $6, $7)
	`,
		*req.TrackerBomId,
		*req.GpsId,
		*req.GpsAntennaId,
		*req.GsmAntennaId,
		*req.SimId,
		*req.BatteryId,
		userId,
	).Scan(&result)

	if err != nil {
		log.Printf("❌ Error executing update: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	// 📤 ตอบกลับผลลัพธ์
	responses.OK(c, dto.EZWUpdateTrackerBomResponse{Message: result})
}

// GetGPS ดึงข้อมูล GPS ทั้งหมด
// @Summary ดึงข้อมูล GPS
// @Description เรียกฟังก์ชัน ezw_get_gps()
// @Tags Tracker Bom
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.GPSResponse]
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/tracker-bom/gps [get]
func GetGPS(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM ezw_get_gps()")
	if err != nil {
		log.Println("Query error:", err)
		responses.InternalServerError(c, "Query error")
		return
	}
	defer rows.Close()

	var result []dto.GPSResponse
	for rows.Next() {
		var m models.GPSModel
		if err := rows.Scan(&m.GpsID, &m.GpsIMEI, &m.SerialNo); err != nil {
			log.Println("Scan error:", err)
			responses.InternalServerError(c, "Scan error")
			return
		}

		resItem := dto.GPSResponse{}
		// map Null → pointer
		if m.GpsID.Valid {
			val := int(m.GpsID.Int64)
			resItem.GpsID = &val
		}
		if m.GpsIMEI.Valid {
			val := m.GpsIMEI.String
			resItem.GpsIMEI = &val
		}
		if m.SerialNo.Valid {
			val := m.SerialNo.String
			resItem.SerialNo = &val
		}

		result = append(result, resItem)
	}

	responses.OK(c, result)
}

// GetBattery ดึงข้อมูล Battery
// @Summary ดึงข้อมูล Battery
// @Description เรียกฟังก์ชัน ezw_get_battery()
// @Tags Tracker Bom
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.BatteryResponse]
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/tracker-bom/battery [get]
func GetBattery(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM ezw_get_battery()")
	if err != nil {
		log.Println("Query error:", err)
		responses.InternalServerError(c, "Query error")
		return
	}
	defer rows.Close()

	var result []dto.BatteryResponse
	for rows.Next() {
		var m models.BatteryModel
		if err := rows.Scan(&m.BatteryID, &m.SerialNo); err != nil {
			log.Println("Scan error:", err)
			responses.InternalServerError(c, "Scan error")
			return
		}

		item := dto.BatteryResponse{}
		if m.BatteryID.Valid {
			val := int(m.BatteryID.Int64)
			item.BatteryID = &val
		}
		if m.SerialNo.Valid {
			val := m.SerialNo.String
			item.SerialNo = &val
		}

		result = append(result, item)
	}

	responses.OK(c, result)
}

// GetCardReader ดึงข้อมูล Card Reader ทั้งหมด
// @Summary ดึงข้อมูล Card Reader
// @Description เรียกฟังก์ชัน ezw_get_card_reader()
// @Tags Tracker Bom
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.CardReaderResponse]
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/tracker-bom/card-reader [get]
func GetCardReader(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM ezw_get_card_reader()")
	if err != nil {
		log.Println("Query error:", err)
		responses.InternalServerError(c, "Query error")
		return
	}
	defer rows.Close()

	var result []dto.CardReaderResponse
	for rows.Next() {
		var m models.CardReaderModel
		if err := rows.Scan(&m.CardReaderID, &m.SerialNo); err != nil {
			log.Println("Scan error:", err)
			responses.InternalServerError(c, "Scan error")
			return
		}

		item := dto.CardReaderResponse{}
		if m.CardReaderID.Valid {
			val := int(m.CardReaderID.Int64)
			item.CardReaderID = &val
		}
		if m.SerialNo.Valid {
			val := m.SerialNo.String
			item.SerialNo = &val
		}

		result = append(result, item)
	}

	responses.OK(c, result)
}

// GetGpsAntenna ดึงข้อมูล Gps Antenna ทั้งหมด
// @Summary ดึงข้อมูล Gps Antenna
// @Description เรียกฟังก์ชัน ezw_get_gps_antenna()
// @Tags Tracker Bom
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.GpsAntennaResponse]
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/tracker-bom/gps-antenna [get]
func GetGpsAntenna(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM ezw_get_gps_antenna()")
	if err != nil {
		log.Println("Query error:", err)
		responses.InternalServerError(c, "Query error")
		return
	}
	defer rows.Close()

	var result []dto.GpsAntennaResponse
	for rows.Next() {
		var m models.GpsAntennaModel
		if err := rows.Scan(&m.GpsAntennaID, &m.SerialNo); err != nil {
			log.Println("Scan error:", err)
			responses.InternalServerError(c, "Scan error")
			return
		}

		item := dto.GpsAntennaResponse{}
		if m.GpsAntennaID.Valid {
			val := int(m.GpsAntennaID.Int64)
			item.GpsAntennaID = &val
		}
		if m.SerialNo.Valid {
			val := m.SerialNo.String
			item.SerialNo = &val
		}

		result = append(result, item)
	}

	responses.OK(c, result)
}

// GetSim ดึงข้อมูล SIM
// @Summary ดึงข้อมูล SIM
// @Description เรียกฟังก์ชัน ezw_get_sim()
// @Tags Tracker Bom
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.SimResponse] "OK"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "Internal Server Error"
// @Router /v1/tracker-bom/sim [get]
func GetSim(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM ezw_get_sim()")
	if err != nil {
		log.Println("Query error:", err)
		responses.InternalServerError(c, "Query error")
		return
	}
	defer rows.Close()

	var result []dto.SimResponse
	for rows.Next() {
		var m models.SimModel
		if err := rows.Scan(&m.SimID, &m.SimNo); err != nil {
			log.Println("Scan error:", err)
			responses.InternalServerError(c, "Scan error")
			return
		}

		item := dto.SimResponse{}
		if m.SimID.Valid {
			val := int(m.SimID.Int64)
			item.SimID = &val
		}
		if m.SimNo.Valid {
			val := m.SimNo.String
			item.SimNo = &val
		}

		result = append(result, item)
	}

	responses.OK(c, result)
}

// GetGsmAntenna ดึงข้อมูล GSM Antenna
// @Summary ดึงข้อมูล GSM Antenna
// @Description เรียกฟังก์ชัน ezw_get_gsm_antenna()
// @Tags Tracker Bom
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.GsmAntennaResponse] "OK"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "Internal Server Error"
// @Router /v1/tracker-bom/gsm-antenna [get]
func GetGsmAntenna(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM ezw_get_gsm_antenna()")
	if err != nil {
		log.Println("Query error:", err)
		responses.InternalServerError(c, "Query error")
		return
	}
	defer rows.Close()

	var result []dto.GsmAntennaResponse
	for rows.Next() {
		var m models.GsmAntennaModel
		if err := rows.Scan(&m.GsmAntennaID, &m.SerialNo); err != nil {
			log.Println("Scan error:", err)
			responses.InternalServerError(c, "Scan error")
			return
		}

		item := dto.GsmAntennaResponse{}
		if m.GsmAntennaID.Valid {
			val := int(m.GsmAntennaID.Int64)
			item.GsmAntennaID = &val
		}
		if m.SerialNo.Valid {
			val := m.SerialNo.String
			item.SerialNo = &val
		}

		result = append(result, item)
	}

	responses.OK(c, result)
}

// GetTrackerBomGeneralByID รับ tracker_bom_id จาก query
// @Summary ดึงข้อมูล Tracker BOM
// @Description เรียกฟังก์ชัน ezw_get_tracker_bom_general_by_tracker_bom_id(p_tracker_bom_id)
// @Tags Tracker Bom
// @Accept json
// @Produce json
// @Security Bearer
// @Param tracker_bom_id query int true "Tracker BOM ID"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWGetTrackerBomGeneralResponse]
// @Failure 400 {object} responses.BadRequestResponseSwagger
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/tracker-bom/general [get]
func GetTrackerBomGeneralByID(c *gin.Context) {
	// 1) รับ tracker_bom_id จาก query
	trackerBomIdStr := c.Query("tracker_bom_id")
	if trackerBomIdStr == "" {
		responses.BadRequest(c, "tracker_bom_id is required in query")
		return
	}
	trackerBomId, err := strconv.Atoi(trackerBomIdStr)
	if err != nil || trackerBomId <= 0 {
		responses.BadRequest(c, "invalid tracker_bom_id")
		return
	}

	db := database.GetDB()
	if db == nil {
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// 2) Query function
	row := db.QueryRow(`
		SELECT *
		FROM public.ezw_get_tracker_bom_general_by_tracker_bom_id($1)
	`, trackerBomId)

	var m models.EZWTrackerBomGeneralModel
	err = row.Scan(
		&m.TrackerBomID,
		&m.TrackerBomCode,
		&m.GpsID,
		&m.GpsIMEI,
		&m.GpsSerialNo,
		&m.GpsAntennaID,
		&m.GpsAntennaSerialNo,
		&m.GsmAntennaID,
		&m.GsmAntennaSerialNo,
		&m.SimID,
		&m.SimNo,
		&m.BatteryID,
		&m.BatterySerialNo,
		&m.CardReaderId,
		&m.CardReaderSerialNo,
		&m.IsActive,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.OK(c, nil) // no data
			return
		}
		log.Printf("Scan error: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	// 3) map sql.Null -> pointer
	resDto := dto.EZWGetTrackerBomGeneralResponse{}

	if m.TrackerBomID.Valid {
		tmp := int(m.TrackerBomID.Int64)
		resDto.TrackerBomID = &tmp
	}
	if m.TrackerBomCode.Valid {
		tmp := m.TrackerBomCode.String
		resDto.TrackerBomCode = &tmp
	}
	if m.GpsID.Valid {
		tmp := int(m.GpsID.Int64)
		resDto.GpsID = &tmp
	}
	if m.GpsIMEI.Valid {
		tmp := m.GpsIMEI.String
		resDto.GpsIMEI = &tmp
	}
	if m.GpsSerialNo.Valid {
		tmp := m.GpsSerialNo.String
		resDto.GpsSerialNo = &tmp
	}
	if m.GpsAntennaID.Valid {
		tmp := int(m.GpsAntennaID.Int64)
		resDto.GpsAntennaID = &tmp
	}
	if m.GpsAntennaSerialNo.Valid {
		tmp := m.GpsAntennaSerialNo.String
		resDto.GpsAntennaSerialNo = &tmp
	}
	if m.GsmAntennaID.Valid {
		tmp := int(m.GsmAntennaID.Int64)
		resDto.GsmAntennaID = &tmp
	}
	if m.GsmAntennaSerialNo.Valid {
		tmp := m.GsmAntennaSerialNo.String
		resDto.GsmAntennaSerialNo = &tmp
	}
	if m.SimID.Valid {
		tmp := int(m.SimID.Int64)
		resDto.SimID = &tmp
	}
	if m.SimNo.Valid {
		tmp := m.SimNo.String
		resDto.SimNo = &tmp
	}
	if m.BatteryID.Valid {
		tmp := int(m.BatteryID.Int64)
		resDto.BatteryID = &tmp
	}
	if m.BatterySerialNo.Valid {
		tmp := m.BatterySerialNo.String
		resDto.BatterySerialNo = &tmp
	}
	if m.CardReaderId.Valid {
		tmp := int(m.CardReaderId.Int64)
		resDto.CardReaderID = &tmp
	}
	if m.CardReaderSerialNo.Valid {
		tmp := m.CardReaderSerialNo.String
		resDto.CardReaderSerialNo = &tmp
	}
	if m.IsActive.Valid {
		tmp := m.IsActive.Bool
		resDto.IsActive = &tmp
	}

	responses.OK(c, resDto)
}
