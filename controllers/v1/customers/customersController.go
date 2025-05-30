package customersController

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/customers"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/customers"
	"ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// GetCustomers ใช้สำหรับสร้างข้อมูลลูกค้าบุคคลทั่วไป
// @Summary สร้างข้อมูลลูกค้าบุคคลทั่วไป
// @Description API นี้ใช้สำหรับสร้างข้อมูลลูกค้าบุคคลทั่วไป
// @Tags Customers
// @Accept mpfd
// @Produce json
// @Param first_name formData string true "First Name"
// @Param last_name formData string true "Last Name"
// @Param id_card formData string true "ID Card"
// @Param phone_number formData string true "Phone Number"
// @Param address formData string true "Address"
// @Param sub_district_id formData string true "Sub District ID"
// @Param district_id formData string true "District ID"
// @Param province_id formData string true "Province ID"
// @Param zip_code formData string true "Zip Code"
// @Param images formData file true "Images"
// @Success 201 {object} responses.CreateDataResponseSwagger[dto.CreateIndividualCustomerResponse] "สร้างข้อมูลลูกค้าบุคคลทั่วไปสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/customers/individual [post]
func CreateIndividualCustomer(c *gin.Context) {
	var reqBody dto.CreateIndividualCustomerRequest
	if err := c.ShouldBind(&reqBody); err != nil {
		responses.BadRequest(c, "Invalid form data")
		return
	}

	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
		return
	}

	// Validate latitude and longitude format
	if reqBody.Latitude != "" {
		if _, err := strconv.ParseFloat(reqBody.Latitude, 64); err != nil {
			responses.BadRequest(c, "Invalid latitude format")
			return
		}
	}

	if reqBody.Longitude != "" {
		if _, err := strconv.ParseFloat(reqBody.Longitude, 64); err != nil {
			responses.BadRequest(c, "Invalid longitude format")
			return
		}
	}

	// Handle multiple file uploads
	form, err := c.MultipartForm()
	if err != nil {
		responses.BadRequest(c, "File upload failed")
		return
	}
	files := form.File["files"]

	// Check file extensions and save files
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	var filePaths []string
	var savedFilePaths []string
	for _, file := range files {
		ext := strings.ToLower(filepath.Ext(file.Filename))
		valid := false
		for _, allowedExt := range allowedExtensions {
			if ext == allowedExt {
				valid = true
				break
			}
		}
		if !valid {
			responses.BadRequest(c, "Invalid file type")
			return
		}

		// Hash the phone number to ensure uniqueness
		hashedPhone := fmt.Sprintf("%x", sha256.Sum256([]byte(reqBody.PhoneNumber)))

		// Save the file with hashed phone number and index as the filename
		filePath := fmt.Sprintf("uploads/%s_%d%s", hashedPhone, time.Now().UnixNano(), ext)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			responses.InternalServerError(c, "Failed to save file")
			return
		}

		// Remove "uploads/" prefix before storing in database
		filePaths = append(filePaths, filePath[len("uploads/"):])
		savedFilePaths = append(savedFilePaths, filePath)
	}

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.ConnectDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// Call the function
	query := `SELECT ezw_register_individual_customer($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)` // 21 parameters
	var customerID int
	err = db.QueryRow(
		query,
		reqBody.FirstName,
		reqBody.LastName,
		reqBody.CallName,
		reqBody.IdCard,
		reqBody.PhoneNumber,
		reqBody.LevelId,
		reqBody.Address,
		reqBody.SubDistrictId,
		reqBody.DistrictId,
		reqBody.ProvinceId,
		reqBody.ZipCode,
		reqBody.Latitude,
		reqBody.Longitude,
		reqBody.ContactName,
		reqBody.ContactLastName,
		reqBody.ContactPhoneNumber,
		reqBody.ContactEmail,
		reqBody.Detail,
		reqBody.DealerId,
		pq.Array(filePaths),
		userID,
	).Scan(&customerID)
	if err != nil {
		log.Printf("Error executing stored procedure: %v", err)

		// Remove uploaded files if the database operation fails
		for _, filePath := range savedFilePaths {
			if err := os.Remove(filePath); err != nil {
				log.Printf("Failed to remove file: %v", err)
			}
		}

		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.CreatedData(c, gin.H{"customer_id": customerID})
}

// CreateJuristicCustomer ใช้สำหรับสร้างข้อมูลลูกค้านิติบุคคล
// @Summary สร้างข้อมูลลูกค้านิติบุคคล
// @Description API นี้ใช้สำหรับสร้างข้อมูลลูกค้านิติบุคคล
// @Tags Customers
// @Accept mpfd
// @Produce json
// @Param juristic_name formData string true "Juristic Name"
// @Param contact_name formData string true "Contact Name"
// @Param contact_lastname formData string true "Contact Lastname"
// @Param contact_position formData string true "Contact Position"
// @Param tax_id formData string true "Tax ID"
// @Param phone_number formData string true "Phone Number"
// @Param address formData string true "Address"
// @Param sub_district_id formData string true "Sub District ID"
// @Param district_id formData string true "District ID"
// @Param province_id formData string true "Province ID"
// @Param zip_code formData string true "Zip Code"
// @Param images formData file true "Images"
// @Success 201 {object} responses.CreateDataResponseSwagger[dto.CreateJuristicCustomerResponse] "สร้างข้อมูลลูกค้านิติบุคคลสำเร็จ
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/customers/juristic [post]
func CreateJuristicCustomer(c *gin.Context) {
	var reqBody dto.CreateJuristicCustomerRequest
	if err := c.ShouldBind(&reqBody); err != nil {
		responses.BadRequest(c, "Invalid form data")
		return
	}

	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
		return
	}

	// Handle multiple file uploads
	form, err := c.MultipartForm()
	if err != nil {
		responses.BadRequest(c, "File upload failed")
		return
	}
	files := form.File["images"]

	// Check file extensions and save files
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	var filePaths []string
	var savedFilePaths []string
	for _, file := range files {
		ext := strings.ToLower(filepath.Ext(file.Filename))
		valid := false
		for _, allowedExt := range allowedExtensions {
			if ext == allowedExt {
				valid = true
				break
			}
		}
		if !valid {
			responses.BadRequest(c, "Invalid file type")
			return
		}

		// Hash the company name to ensure uniqueness
		hashedJuristic := fmt.Sprintf("%x", sha256.Sum256([]byte(reqBody.JuristicName)))

		// Save the file with hashed company name and index as the filename
		filePath := fmt.Sprintf("uploads/%s_%d%s", hashedJuristic, time.Now().UnixNano(), ext)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			responses.InternalServerError(c, "Failed to save file")
			return
		}

		// Remove "uploads/" prefix before storing in database
		filePaths = append(filePaths, filePath[len("uploads/"):])
		savedFilePaths = append(savedFilePaths, filePath)
	}

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.ConnectDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// Call the stored procedure
	query := `SELECT public.ezw_register_juristic_customer($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)` // 25 parameters
	var customerID int
	err = db.QueryRow(
		query,
		reqBody.JuristicName,
		reqBody.CallName,
		reqBody.TaxID,
		reqBody.HeadOfficeNumber,
		reqBody.BranchNumber,
		reqBody.PhoneNumber,
		reqBody.LevelId,
		reqBody.Address,
		reqBody.SubDistrictId,
		reqBody.DistrictId,
		reqBody.ProvinceId,
		reqBody.ZipCode,
		reqBody.Latitude,
		reqBody.Longitude,
		reqBody.ContactName,
		reqBody.ContactLastName,
		reqBody.ContactPhoneNumber,
		reqBody.ContactEmail,
		reqBody.AuthorizedName,
		reqBody.AuthorizedLastName,
		reqBody.AuthorizedPhoneNumber,
		reqBody.Detail,
		reqBody.DealerId,
		pq.Array(filePaths),
		userID,
	).Scan(&customerID)
	if err != nil {
		log.Printf("Error executing stored procedure: %v", err)

		// Remove uploaded files if the database operation fails
		for _, filePath := range savedFilePaths {
			if err := os.Remove(filePath); err != nil {
				log.Printf("Failed to remove file: %v", err)
			}
		}

		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.CreatedData(c, gin.H{"customer_id": customerID})
}

func GetCustomersByConditions(c *gin.Context) {

	var reqBody dto.GetCustomersByConditionsRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	customerId := sql.NullString{}
	if reqBody.CustomerId != "" {
		customerId = sql.NullString{String: reqBody.CustomerId, Valid: true}
	}

	customerName := sql.NullString{}
	if reqBody.CustomerName != "" {
		customerName = sql.NullString{String: reqBody.CustomerName, Valid: true}
	}

	customerGroup := sql.NullInt64{}
	if reqBody.CustomerGroup != 0 {
		customerGroup = sql.NullInt64{Int64: reqBody.CustomerGroup, Valid: true}
	}

	customerDetail := sql.NullString{}
	if reqBody.CustomerDetail != "" {
		customerDetail = sql.NullString{String: reqBody.CustomerDetail, Valid: true}
	}

	customerStatus := sql.NullInt64{}
	if reqBody.CustomerStatus != 0 {
		customerStatus = sql.NullInt64{Int64: reqBody.CustomerStatus, Valid: true}
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
	query := `SELECT * FROM ezw_search_customers($1, $2, $3, $4, $5)`
	rows, err := db.Query(query, customerId, customerName, customerGroup, customerDetail, customerStatus)

	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch customers")
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลลูกค้า
	var response []dto.GetCustomersByConditionsResponse

	// วนลูปเพื่อดึงข้อมูลลูกค้า
	for rows.Next() {
		var customer models.EZWSearchCustomers
		err := rows.Scan(&customer.CustomerId,
			&customer.CustomerName,
			&customer.CustomerGroupId,
			&customer.CustomerGroup,
			&customer.CustomerDetail,
			&customer.CustomerStatus)
		if err != nil {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch customers")
			return
		}

		response = append(response, dto.GetCustomersByConditionsResponse{
			CustomerId:      customer.CustomerId,
			CustomerName:    customer.CustomerName,
			CustomerGroupId: customer.CustomerGroupId,
			CustomerGroup:   customer.CustomerGroup,
			CustomerDetail:  customer.CustomerDetail,
			CustomerStatus:  customer.CustomerStatus,
		})
	}

	if len(response) == 0 {
		responses.OK(c, []dto.GetCustomersByConditionsResponse{})
		return
	}

	responses.OK(c, response)
}

func GetCustomerGroups(c *gin.Context) {

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}
	// ใช้ Prepared Statement (แบบไม่ต้องเรียก `Prepare()` เอง)
	query := `select type_id, type_name from system_master_types smt where smt.category_id = 105`
	rows, err := db.Query(query)

	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch customers")
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลลูกค้า
	var response []dto.GetCustomerGroupsResponse

	// วนลูปเพื่อดึงข้อมูลลูกค้า
	for rows.Next() {
		var customerGroup models.EZWGetCustomerGroup
		err := rows.Scan(&customerGroup.CustomerGroupId,
			&customerGroup.CustomerGroupName)
		if err != nil {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch customers")
			return
		}

		response = append(response, dto.GetCustomerGroupsResponse{
			CustomerGroupId:   customerGroup.CustomerGroupId,
			CustomerGroupName: customerGroup.CustomerGroupName,
		})
	}

	if len(response) == 0 {
		responses.OK(c, []dto.GetCustomerGroupsResponse{})
		return
	}

	responses.OK(c, response)
}

func GetCustomerStatus(c *gin.Context) {

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}
	// ใช้ Prepared Statement (แบบไม่ต้องเรียก `Prepare()` เอง)
	query := `select type_id, type_name from system_master_types smt where smt.category_id = 106`
	rows, err := db.Query(query)

	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch customers")
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลลูกค้า
	var response []dto.GetCustomerStatusResponse

	// วนลูปเพื่อดึงข้อมูลลูกค้า
	for rows.Next() {
		var customerStatus models.EZWCustomerStatus
		err := rows.Scan(&customerStatus.StatusId,
			&customerStatus.StatusName)
		if err != nil {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch customers")
			return
		}

		response = append(response, dto.GetCustomerStatusResponse{
			StatusId:   customerStatus.StatusId,
			StatusName: customerStatus.StatusName,
		})
	}

	if len(response) == 0 {
		responses.OK(c, []dto.GetCustomerStatusResponse{})
		return
	}

	responses.OK(c, response)
}

func GetIndividualCustomerById(c *gin.Context) {

	customerId := c.Param("id")
	if customerId == "" {
		responses.BadRequest(c, "Missing required path parameter: id")
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
	query := `SELECT customer_id, customer_name, customer_lastname, customer_callname, id_card, phone, level, address, province_id, district_id, subdistrict_id, zipcode, COALESCE(description, '') FROM customers WHERE customer_type_id = 105002 AND customer_id = $1 AND is_active = true`
	rows, err := db.Query(query, customerId)

	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch customers")
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลลูกค้า
	var response dto.GetCustomerByIdResponse

	// วนลูปเพื่อดึงข้อมูลลูกค้า
	for rows.Next() {
		var customer models.EZWGetCustomerById
		err := rows.Scan(&customer.CustomerId,
			&customer.CustomerName,
			&customer.CustomerLast,
			&customer.CallName,
			&customer.CustomerIdCard,
			&customer.CustomerTel,
			&customer.LevelId,
			&customer.CustomerAddress,
			&customer.CustomerProvinceId,
			&customer.CustomerDistrictId,
			&customer.CustomerSubDistrictId,
			&customer.CustomerZipCode,
			&customer.Detail)

		if err != nil {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch customers")
			return
		}

		response = dto.GetCustomerByIdResponse{
			CustomerId:            customer.CustomerId,
			CustomerName:          customer.CustomerName,
			CustomerLast:          customer.CustomerLast,
			CallName:              customer.CallName,
			CustomerIdCard:        customer.CustomerIdCard,
			CustomerTel:           customer.CustomerTel,
			LevelId:               customer.LevelId,
			CustomerAddress:       customer.CustomerAddress,
			CustomerProvinceId:    customer.CustomerProvinceId,
			CustomerDistrictId:    customer.CustomerDistrictId,
			CustomerSubDistrictId: customer.CustomerSubDistrictId,
			CustomerZipCode:       customer.CustomerZipCode,
			Detail:                customer.Detail,
		}
	}

	responses.OK(c, response)
}

func UpdateIndividualCustomer(c *gin.Context) {
	customerId := c.Param("id")
	if customerId == "" {
		responses.BadRequest(c, "Missing required path parameter: id")
		return
	}

	var reqBody dto.UpdateIndividualCustomerRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Printf("JSON Binding Error: %v", err)
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
		return
	}

	// Validate latitude and longitude format
	if reqBody.Latitude != "" {
		if _, err := strconv.ParseFloat(reqBody.Latitude, 64); err != nil {
			responses.BadRequest(c, "Invalid latitude format")
			return
		}
	}

	if reqBody.Longitude != "" {
		if _, err := strconv.ParseFloat(reqBody.Longitude, 64); err != nil {
			responses.BadRequest(c, "Invalid longitude format")
			return
		}
	}

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.ConnectDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// Call the function
	query := `SELECT ezw_update_individual_customer($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)` // 16 parameters
	var customerID int
	err := db.QueryRow(
		query,
		customerId,
		reqBody.FirstName,
		reqBody.LastName,
		reqBody.CallName,
		reqBody.IdCard,
		reqBody.PhoneNumber,
		reqBody.LevelId,
		reqBody.Address,
		reqBody.SubDistrictId,
		reqBody.DistrictId,
		reqBody.ProvinceId,
		reqBody.ZipCode,
		reqBody.Latitude,
		reqBody.Longitude,
		reqBody.Detail,
		userID,
	).Scan(&customerID)
	if err != nil {
		log.Printf("Error executing stored procedure: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.OK(c, "Customer updated successfully")
}

func GetJuristicCustomerById(c *gin.Context) {

	customerId := c.Param("id")
	if customerId == "" {
		responses.BadRequest(c, "Missing required path parameter: id")
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
	query := `SELECT customer_id, juristic_name, customer_callname, id_card, head_office_no, branch_no, phone, level, address, province_id, district_id, subdistrict_id, zipcode, authorized_name, authorized_lastname, authorized_phone, COALESCE(description, '') FROM customers WHERE customer_type_id = 105001 AND customer_id = $1 AND is_active = true`
	rows, err := db.Query(query, customerId)

	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch customers")
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลลูกค้า
	var response dto.GetJuristicCustomerByIdResponse

	// วนลูปเพื่อดึงข้อมูลลูกค้า
	for rows.Next() {
		var customer models.EZWGetJuristicCustomerById
		err := rows.Scan(&customer.CustomerId,
			&customer.JuristicName,
			&customer.CallName,
			&customer.CustomerIdCard,
			&customer.HeadOfficeNumber,
			&customer.BranchNumber,
			&customer.CustomerTel,
			&customer.LevelId,
			&customer.CustomerAddress,
			&customer.CustomerProvinceId,
			&customer.CustomerDistrictId,
			&customer.CustomerSubDistrictId,
			&customer.CustomerZipCode,
			&customer.AuthorizedName,
			&customer.AuthorizedLastName,
			&customer.AuthorizedPhoneNumber,
			&customer.Detail)

		if err != nil {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch customers")
			return
		}

		response = dto.GetJuristicCustomerByIdResponse{
			CustomerId:            customer.CustomerId,
			JuristicName:          customer.JuristicName,
			CallName:              customer.CallName,
			CustomerIdCard:        customer.CustomerIdCard,
			HeadOfficeNumber:      customer.HeadOfficeNumber,
			BranchNumber:          customer.BranchNumber,
			CustomerTel:           customer.CustomerTel,
			LevelId:               customer.LevelId,
			CustomerAddress:       customer.CustomerAddress,
			CustomerProvinceId:    customer.CustomerProvinceId,
			CustomerDistrictId:    customer.CustomerDistrictId,
			CustomerSubDistrictId: customer.CustomerSubDistrictId,
			CustomerZipCode:       customer.CustomerZipCode,
			AuthorizedName:        customer.AuthorizedName,
			AuthorizedLastName:    customer.AuthorizedLastName,
			AuthorizedPhoneNumber: customer.AuthorizedPhoneNumber,
			Detail:                customer.Detail,
		}
	}

	responses.OK(c, response)
}

func UpdateJuristicCustomer(c *gin.Context) {
	customerId := c.Param("id")
	if customerId == "" {
		responses.BadRequest(c, "Missing required path parameter: id")
		return
	}

	var reqBody dto.UpdateJuristicCustomerRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Printf("JSON Binding Error: %v", err)
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
		return
	}

	// Validate latitude and longitude format
	if reqBody.Latitude != "" {
		if _, err := strconv.ParseFloat(reqBody.Latitude, 64); err != nil {
			responses.BadRequest(c, "Invalid latitude format")
			return
		}
	}

	if reqBody.Longitude != "" {
		if _, err := strconv.ParseFloat(reqBody.Longitude, 64); err != nil {
			responses.BadRequest(c, "Invalid longitude format")
			return
		}
	}

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.ConnectDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// Call the function
	query := `SELECT ezw_update_juristic_customer($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)` // 20 parameters
	var customerID int
	err := db.QueryRow(
		query,
		customerId,
		reqBody.JuristicName,
		reqBody.CallName,
		reqBody.IdCard,
		reqBody.HeadOfficeNo,
		reqBody.BranchNo,
		reqBody.PhoneNumber,
		reqBody.LevelId,
		reqBody.Address,
		reqBody.SubDistrictId,
		reqBody.DistrictId,
		reqBody.ProvinceId,
		reqBody.ZipCode,
		reqBody.AuthorizedName,
		reqBody.AuthorizedLastName,
		reqBody.AuthorizedPhoneNumber,
		reqBody.Latitude,
		reqBody.Longitude,
		reqBody.Detail,
		userID,
	).Scan(&customerID)
	if err != nil {
		log.Printf("Error executing stored procedure: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.OK(c, "Customer updated successfully")
}

func GetCustomerContactByCustomerId(c *gin.Context) {
	customerId := c.Param("id")
	if customerId == "" {
		responses.BadRequest(c, "Missing required path parameter: id")
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
	query := `SELECT * FROM ezw_get_customer_contacts($1)`
	rows, err := db.Query(query, customerId)

	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch customer contacts")
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลลูกค้า
	var response []dto.GetCustomerContactsResponse

	// วนลูปเพื่อดึงข้อมูลลูกค้า
	for rows.Next() {
		var contact models.EZWGetCustomerContactByCustomerId
		err := rows.Scan(&contact.ContactId,
			&contact.ContactName,
			&contact.ContactLastName,
			&contact.ContactEmail,
			&contact.ContactPhone,
			&contact.ContactType)

		if err != nil {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch customers")
			return
		}

		response = append(response, dto.GetCustomerContactsResponse{
			ContactId:       contact.ContactId,
			ContactName:     contact.ContactName,
			ContactLastName: contact.ContactLastName,
			ContactPhone:    contact.ContactPhone,
			ContactEmail:    contact.ContactEmail,
			ContactType:     contact.ContactType,
		})
	}

	if len(response) == 0 {
		responses.OK(c, []dto.GetCustomerContactsResponse{})
		return
	}

	responses.OK(c, response)
}

func CreateCustomerContact(c *gin.Context) {

	var reqBody dto.CreateCustomerContactRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Printf("JSON Binding Error: %v", err)
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
		return
	}

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.ConnectDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// Call the function
	query := `SELECT ezw_create_customer_contact($1, $2, $3, $4, $5, $6, $7)` // 7 parameters
	var contactID int
	err := db.QueryRow(
		query,
		reqBody.CustomerId,
		reqBody.ContactName,
		reqBody.ContactLastName,
		reqBody.ContactPhone,
		reqBody.ContactEmail,
		reqBody.ContactType,
		userID,
	).Scan(&contactID)
	if err != nil {
		log.Printf("Error executing stored procedure: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.CreatedData(c, gin.H{"contact_id": contactID})
}

func UpdateCustomerContact(c *gin.Context) {

	var reqBody dto.UpdateCustomerContactRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Printf("JSON Binding Error: %v", err)
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
		return
	}

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.ConnectDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// Call the function
	query := `SELECT ezw_update_customer_contact($1, $2, $3, $4, $5, $6, $7)` // 6 parameters
	var contactID bool
	err := db.QueryRow(
		query,
		reqBody.ContactId,
		reqBody.ContactName,
		reqBody.ContactLastName,
		reqBody.ContactPhone,
		reqBody.ContactEmail,
		reqBody.ContactType,
		userID,
	).Scan(&contactID)
	if err != nil {
		log.Printf("Error executing stored procedure: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.OK(c, "Customer contact updated successfully")
}

func DeleteCustomerContact(c *gin.Context) {
	contactId := c.Param("id")
	if contactId == "" {
		responses.BadRequest(c, "Missing required path parameter: id")
		return
	}

	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
		return
	}

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.ConnectDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// Call the function
	query := `SELECT ezw_delete_customer_contact($1, $2)` // 2 parameters
	var deletedContactID bool
	err := db.QueryRow(
		query,
		contactId,
		userID,
	).Scan(&deletedContactID)
	if err != nil {
		log.Printf("Error executing stored procedure: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.OK(c, "Customer contact deleted successfully")
}

func GetCustomerOthersByCustomerId(c *gin.Context) {
	customerId := c.Param("id")
	if customerId == "" {
		responses.BadRequest(c, "Missing required path parameter: id")
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
	query := `SELECT 
    		u1.user_name AS created_by, 
    		TO_CHAR(c.created_at, 'DD/MM/YYYY HH24:MI:SS') AS created_at, 
    		u2.user_name AS modified_by, 
    		TO_CHAR(c.modified_at, 'DD/MM/YYYY HH24:MI:SS') AS modified_at
		FROM customers c
		LEFT JOIN users u1 ON u1.user_id = c.created_by
		LEFT JOIN users u2 ON u2.user_id = c.modified_by
		WHERE customer_id = $1`
	row := db.QueryRow(query, customerId)

	// สร้าง response สำหรับข้อมูลลูกค้า
	var response dto.GetCustomerOthersByCustomerIdResponse

	var other models.EZWGetCustomerOthersByCustomerId

	// ดึงข้อมูลลูกค้า
	err := row.Scan(
		&other.CreatedBy,
		&other.CreatedAt,
		&other.UpdatedBy,
		&other.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.NotFound(c, "Customer not found")
		} else {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch customer")
		}
		return
	}

	response = dto.GetCustomerOthersByCustomerIdResponse{
		CreatedBy: other.CreatedBy,
		CreatedAt: other.CreatedAt,
		UpdatedBy: other.UpdatedBy,
		UpdatedAt: other.UpdatedAt,
	}

	responses.OK(c, response)

}

func DeleteCustomerById(c *gin.Context) {
	customerId := c.Param("id")
	if customerId == "" {
		responses.BadRequest(c, "Missing required path parameter: id")
		return
	}

	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
		return
	}

	// Convert customerId to integer
	customerIDInt, err := strconv.Atoi(customerId)
	if err != nil {
		responses.BadRequest(c, "Invalid customer ID")
		return
	}

	// Connect to the database
	db := database.ConnectDB()

	// Check database connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// Call the function
	query := `SELECT ezw_delete_customer($1, $2)` // 2 parameters
	var deletedCustomerID int
	err = db.QueryRow(
		query,
		customerIDInt,
		userID,
	).Scan(&deletedCustomerID)
	if err != nil {
		log.Printf("Error executing stored procedure: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.OK(c, "Customer and related records deactivated successfully")
}

func GetUsersCustomerById(c *gin.Context) {
	customerId := c.Param("id")
	if customerId == "" {
		responses.BadRequest(c, "Missing required path parameter: id")
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
	query := `SELECT * FROM ezw_get_user_info($1)`
	rows, err := db.Query(query, customerId)

	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch users")
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลลูกค้า
	var response []dto.GetUsersCustomerByIdResponse

	// วนลูปเพื่อดึงข้อมูลลูกค้า
	for rows.Next() {
		var user models.EZWGetUsersCustomerById
		err := rows.Scan(&user.UserId,
			&user.UserName,
			&user.Name,
			&user.Description)

		if err != nil {
			log.Println("Database scan error:", err)
			responses.InternalServerError(c, "Failed to fetch customers")
			return
		}

		response = append(response, dto.GetUsersCustomerByIdResponse{
			UserId:      user.UserId,
			UserName:    user.UserName,
			Name:        user.Name,
			Description: user.Description,
		})
	}

	if len(response) == 0 {
		responses.OK(c, []dto.GetUsersCustomerByIdResponse{})
		return
	}

	responses.OK(c, response)
}
