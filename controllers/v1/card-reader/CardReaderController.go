package CardReaderControllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/card-reader"

	middlewares "ezview.asia/ezview-web/ezview-lite-back-office/middlewares"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/card-reader"
	responses "ezview.asia/ezview-web/ezview-lite-back-office/types/responses"

	"github.com/gin-gonic/gin"
)

// GetCardReaderStatus ดึงข้อมูล Card Reader Status
// @Summary ดึงข้อมูล Card Reader Status
// @Description API นี้ใช้สำหรับดึงข้อมูลสถานะแบตเตอรี่จากฐานข้อมูล
// @Tags Card Reader
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWCardReaderStatusResponse] "ดึงข้อมูลสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/card-reader/status [get]
func GetCardReaderStatus(c *gin.Context) {

	jwtMiddleware := middlewares.JWTMiddleware()
	jwtMiddleware(c)

	// หลังจาก middleware ทำงาน ค่า "user_id" ควรจะถูกตั้งไว้ใน context
	userID, exists := c.Get("user_id")
	if !exists {
		responses.Unauthorized(c, "User not authenticated")
		return
	}

	uid, ok := userID.(int)
	if !ok {
		responses.InternalServerError(c, "User ID format is invalid")
		return
	}

	log.Printf("User %d accessing Card Reader Model", uid)

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	rows, err := db.Query("SELECT * FROM public.ezw_get_card_reader_status()")
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var status []dto.EZWCardReaderStatusResponse

	status = append(status, dto.EZWCardReaderStatusResponse{

		StatusId:   0,
		StatusName: "ทั้งหมด",
	})
	for rows.Next() {
		var cardReader models.EZWCardReaderStatusModel
		if err := rows.Scan(&cardReader.StatusId, &cardReader.StatusName); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		status = append(status, dto.EZWCardReaderStatusResponse{
			StatusId:   cardReader.StatusId,
			StatusName: cardReader.StatusName,
		})
	}

	responses.OK(c, status)
}

// GetCardReaderBrand ดึงข้อมูล Card Reader Brand
// @Summary ดึงข้อมูล Card Reader Brand
// @Description API นี้ใช้สำหรับดึงข้อมูลแบรนด์ของ Card Reader จากฐานข้อมูล
// @Tags Card Reader
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWCardReaderBrandResponse] "ดึงข้อมูลแบรนด์ Card Reader สำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "ข้อมูลที่ส่งมาไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/card-reader/brand [get]
func GetCardReaderBrand(c *gin.Context) {
	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	query := `SELECT * FROM public.ezw_get_card_reader_brand()`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var brands []dto.EZWCardReaderBrandResponse

	brands = append(brands, dto.EZWCardReaderBrandResponse{

		BrandId:   0,
		BrandName: "ทั้งหมด",
	})
	for rows.Next() {
		var cardReaderBrand models.EZWCardReaderBrand
		// ฟังก์ชัน SQL คืนค่า: model_id, type_name
		if err := rows.Scan(&cardReaderBrand.BrandId, &cardReaderBrand.BrandName); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		brands = append(brands, dto.EZWCardReaderBrandResponse{
			BrandId:   cardReaderBrand.BrandId,
			BrandName: cardReaderBrand.BrandName,
		})
	}

	responses.OK(c, brands)
}

// GetCardReaderModel ดึงข้อมูล Card Reader Model ตาม parent_id
// @Summary ดึงข้อมูล Card Reader Model
// @Description API นี้ใช้สำหรับดึงข้อมูลรุ่นของ Card Reader โดยใช้ parent_id (ไอดีของแบรนด์)
// @Tags Card Reader
// @Accept json
// @Produce json
// @Security Bearer
// @Param parent_id query int true "ไอดีของแบรนด์ (parent_id)"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWCardReaderModelResponse] "ดึงข้อมูลรุ่น Card Reader สำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "ข้อมูลที่ส่งมาไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/card-reader/model [get]
func GetCardReaderModel(c *gin.Context) {
	// ดึงค่า user_id ที่ถูกตั้งไว้ใน context จาก JWTMiddleware
	userID, exists := c.Get("user_id")

	userId2, _ := c.Get("user_id")
	log.Printf("userId2", userId2)

	log.Printf("userID", userID)
	if !exists {
		responses.Unauthorized(c, "User not authenticated")
		return
	}

	uid, ok := userID.(int)
	if !ok {
		responses.InternalServerError(c, "User ID format is invalid")
		return
	}

	// แสดง log ด้วยค่า user_id ที่ได้
	log.Printf("User %d accessing Card Reader Model", uid)

	// ส่วนที่เหลือของ handler (เช่น รับ parent_id, query ฐานข้อมูล ฯลฯ)
	parentIDStr := c.Query("parent_id")
	if parentIDStr == "" {
		responses.BadRequest(c, "Missing parent_id parameter")
		return
	}
	parentID, err := strconv.Atoi(parentIDStr)
	if err != nil {
		responses.BadRequest(c, "Invalid parent_id parameter")
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

	query := `SELECT * FROM public.ezw_get_card_readers_model($1::int)`
	rows, err := db.Query(query, parentID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	var models []dto.EZWCardReaderModelResponse

	models = append(models, dto.EZWCardReaderModelResponse{

		ModelId:   0,
		ModelName: "ทั้งหมด",
	})
	for rows.Next() {
		var cardReaderModel dto.EZWCardReaderModelResponse
		if err := rows.Scan(&cardReaderModel.ModelId, &cardReaderModel.ModelName); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}
		models = append(models, cardReaderModel)
	}

	responses.OK(c, models)
}

// SearchCardReader ค้นหา Card Reader ตามเงื่อนไขพารามิเตอร์
// @Summary ค้นหา Card Reader
// @Description รับ JSON และส่งพารามิเตอร์ไปยังฟังก์ชัน ezw_search_card_readers
// @Tags Card Reader
// @Accept json
// @Produce json
// @Security Bearer
// @Param searchRequest body dto.EZWSearchCardReaderModelRequest true "Payload สำหรับค้นหา Card Reader"
// @Success 201 {object} responses.SuccessResponseSwagger[dto.EZWSearchCardReaderModelResponse] "ผลลัพธ์การค้นหา Card Reader"
// @Failure 400 {object} responses.BadRequestResponseSwagger "ข้อมูลที่ส่งมาไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/card-reader/search [post]
func SearchCardReader(c *gin.Context) {
	// ผูกข้อมูล JSON จาก Request Body กับโครงสร้าง DTO ที่เตรียมไว้
	var reqBody dto.EZWSearchCardReaderModelRequest

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&reqBody); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// เชื่อมต่อฐานข้อมูล
	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// เตรียมตัวแปรแบบ sql.Nullxxx สำหรับแต่ละพารามิเตอร์
	var (
		serialNoNull sql.NullString
		brandIDNull  sql.NullInt32
		modelIDNull  sql.NullInt32
		statusIDNull sql.NullInt32
		remarkNull   sql.NullString
	)

	if reqBody.SerialNo != "" {
		serialNoNull = sql.NullString{String: reqBody.SerialNo, Valid: true}
	}
	if reqBody.BrandId != 0 {
		brandIDNull = sql.NullInt32{Int32: int32(reqBody.BrandId), Valid: true}
	}
	if reqBody.ModelId != 0 {
		modelIDNull = sql.NullInt32{Int32: int32(reqBody.ModelId), Valid: true}
	}
	if reqBody.StatusId != 0 {
		statusIDNull = sql.NullInt32{Int32: int32(reqBody.StatusId), Valid: true}
	}
	if reqBody.Remark != "" {
		remarkNull = sql.NullString{String: reqBody.Remark, Valid: true}
	}

	// คำสั่ง SQL สำหรับเรียกฟังก์ชัน ezw_search_card_readers
	query := `
		SELECT * FROM public.ezw_search_card_readers(
			$1::varchar,
			$2::int,
			$3::int,
			$4::int,
			$5::varchar
		)
	`

	// เรียกใช้ Query ด้วยพารามิเตอร์ที่เตรียมไว้
	rows, err := db.Query(query,
		serialNoNull,
		brandIDNull,
		modelIDNull,
		statusIDNull,
		remarkNull,
	)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}
	defer rows.Close()

	// ผลลัพธ์ที่ได้จาก SQL จะถูกเก็บในโครงสร้างของ models (ตามที่กำหนดไว้)
	var results []dto.EZWSearchCardReaderModelResponse
	for rows.Next() {
		var cardReader models.EZWSearchCardReader
		// ลำดับการสแกนต้องตรงกับลำดับที่ฟังก์ชัน SQL ส่งกลับ
		if err := rows.Scan(
			&cardReader.CardReaderId, // column 1: card_reader_id
			&cardReader.SerialNo,     // column 2: serial_no
			&cardReader.BrandId,      // column 3: brand_id
			&cardReader.BrandName,    // column 4: brand_name
			&cardReader.ModelId,      // column 5: model_id
			&cardReader.ModelName,    // column 6: model_name
			&cardReader.StatusId,     // column 7: status_id
			&cardReader.StatusName,   // column 8: status_name
			&cardReader.Remark,       // column 9: remark

		); err != nil {
			log.Printf("Error scanning row: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
			return
		}

		// แปลงข้อมูลจาก models ไปเป็น DTO สำหรับ response
		results = append(results, dto.EZWSearchCardReaderModelResponse{
			CardReaderId: cardReader.CardReaderId,
			SerialNo:     cardReader.SerialNo,
			BrandId:      cardReader.BrandId,
			BrandName:    cardReader.BrandName,
			ModelId: func() *int {
				if cardReader.ModelId.Valid {
					v := int(cardReader.ModelId.Int32)
					return &v
				}
				return nil
			}(),
			ModelName: func() *string {
				if cardReader.ModelName.Valid {
					v := cardReader.ModelName.String
					return &v
				}
				return nil
			}(),
			StatusId:   cardReader.StatusId,
			StatusName: cardReader.StatusName,
			Remark: func() *string {
				if cardReader.Remark.Valid {
					v := cardReader.Remark.String
					return &v
				}
				return nil
			}(),
		})
	}

	// หากไม่มีผลลัพธ์ ให้ส่งกลับ array ว่าง
	if len(results) == 0 {
		responses.OK(c, []dto.EZWSearchCardReaderModelResponse{})
		return
	}

	responses.OK(c, results)
}

// InsertCardReader เพิ่มข้อมูล Card Reader
// @Summary เพิ่มข้อมูล Card Reader
// @Description เพิ่ม Card Reader ใหม่ลงในระบบ
// @Tags Card Reader
// @Accept json
// @Produce json
// @Security Bearer
// @Param cardReader body dto.EZWInsertCardReaderRequest true "ข้อมูล Card Reader ที่ต้องการเพิ่ม"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWInsertCardReaderResponse]
// @Failure 400 {object} responses.BadRequestResponseSwagger
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/card-reader/insert [post]
func InsertCardReader(c *gin.Context) {
	var req dto.EZWInsertCardReaderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// ตรวจสอบ Validation เบื้องต้น
	if req.SerialNo == "" {
		responses.BadRequest(c, "serial_no is required and cannot be empty")
		return
	}
	if req.BrandId <= 0 {
		responses.BadRequest(c, "brand_id must be positive integer")
		return
	}
	if req.ModelId <= 0 {
		responses.BadRequest(c, "model_id must be positive integer")
		return
	}
	if req.StatusId <= 0 {
		responses.BadRequest(c, "status_id must be positive integer")
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
	err := db.QueryRow(`
		SELECT public.ezw_insert_card_reader($1, $2, $3, $4, $5, $6)
	`,
		req.SerialNo, // p_serial_no
		req.BrandId,  // p_brand_id
		req.ModelId,  // p_model_id
		req.StatusId, // p_status_id
		req.Remark,   // p_remark
		userID,       // p_user_id
	).Scan(&result)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.OK(c, dto.EZWInsertCardReaderResponse{Message: result})
}

// UpdateCardReader แก้ไขข้อมูล Card Reader
// @Summary แก้ไขข้อมูล Card Reader
// @Description แก้ไขข้อมูล Card Reader ที่มีอยู่ในระบบ
// @Tags Card Reader
// @Accept json
// @Produce json
// @Security Bearer
// @Param cardReader body dto.EZWUpdateCardReaderRequest true "ข้อมูล Card Reader ที่ต้องการอัปเดต"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWUpdateCardReaderResponse]
// @Failure 400 {object} responses.BadRequestResponseSwagger
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/card-reader/update [put]
func UpdateCardReader(c *gin.Context) {
	var req dto.EZWUpdateCardReaderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// ตรวจสอบ Validation เบื้องต้น
	if req.CardReaderId <= 0 {
		responses.BadRequest(c, "card_reader_id must be positive integer")
		return
	}
	if req.BrandId <= 0 {
		responses.BadRequest(c, "brand_id must be positive integer")
		return
	}
	if req.ModelId <= 0 {
		responses.BadRequest(c, "model_id must be positive integer")
		return
	}
	if req.StatusId <= 0 {
		responses.BadRequest(c, "status_id must be positive integer")
		return
	}
	if req.SerialNo == "" {
		responses.BadRequest(c, "serial_no is required and cannot be empty")
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
	err := db.QueryRow(`
		SELECT public.ezw_update_card_reader($1, $2, $3, $4, $5, $6, $7)
	`,
		req.CardReaderId, // p_card_reader_id
		req.BrandId,      // p_brand_id
		req.ModelId,      // p_model_id
		req.StatusId,     // p_status_id
		userID,           // p_user_id
		req.SerialNo,     // p_serial_no
		req.Remark,       // p_remark
	).Scan(&result)

	if err != nil {
		if err == sql.ErrNoRows {
			responses.NotFound(c, "Card Reader not found")
			return
		}
		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.OK(c, dto.EZWUpdateCardReaderResponse{Message: result})
}

// @Summary Get Card Reader by ID
// @Description call ezw_get_card_reader_general_by_id($1)
// @Tags Card Reader
// @Accept json
// @Produce json
// @Security Bearer
// @Param card_reader_id query int true "Card Reader ID"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.EZWCardReaderGeneralResponse]
// @Failure 400 {object} responses.BadRequestResponseSwagger
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger
// @Router /v1/card-reader/general [get]
func GetCardReaderGeneral(c *gin.Context) {
	idStr := c.Query("card_reader_id")
	if idStr == "" {
		responses.BadRequest(c, "card_reader_id is required")
		return
	}
	idVal, err := strconv.Atoi(idStr)
	if err != nil || idVal <= 0 {
		responses.BadRequest(c, "invalid card_reader_id")
		return
	}

	db := database.GetDB()
	if db == nil {
		responses.InternalServerError(c, "Database connection error")
		return
	}

	row := db.QueryRow(`SELECT * FROM ezw_get_card_reader_general_by_id($1::smallint)`, idVal)

	var m models.CardReaderGeneralModel
	if err := row.Scan(
		&m.CardReaderID,
		&m.BrandID,
		&m.BrandName,
		&m.ModelID,
		&m.ModelName,
		&m.SerialNo,
		&m.StatusID,
		&m.StatusName,
		&m.Remark,
	); err != nil {
		if err == sql.ErrNoRows {
			responses.OK(c, nil)
			return
		}
		log.Printf("Error scanning card_reader: %v\n", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	resp := dto.EZWCardReaderGeneralResponse{}
	if m.CardReaderID.Valid {
		tmp := int(m.CardReaderID.Int64)
		resp.CardReaderID = &tmp
	}
	if m.BrandID.Valid {
		tmp := int(m.BrandID.Int64)
		resp.BrandID = &tmp
	}
	if m.BrandName.Valid {
		tmp := m.BrandName.String
		resp.BrandName = &tmp
	}
	if m.ModelID.Valid {
		tmp := int(m.ModelID.Int64)
		resp.ModelID = &tmp
	}
	if m.ModelName.Valid {
		tmp := m.ModelName.String
		resp.ModelName = &tmp
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
