package GpsControllers

import (
	"log"
	"strconv"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/gps"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/gps"
	"ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	"github.com/gin-gonic/gin"
)

// GetProvince ดึงข้อมูลจังหวัด
// @Summary ดึงข้อมูลจังหวัดตามภาษาที่เลือก
// @Description API นี้ใช้สำหรับดึงรายชื่อจังหวัดตามภาษาที่กำหนด (th หรือ en)
// @Tags Dropdowns
// @Accept json
// @Produce json
// @Param lang query string false "เลือกภาษา (th หรือ en)" Enums(th, en) default(th)
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.ProvinceResponse] "การร้องขอข้อมูลจังหวัดสำเร็จ"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "Token ไม่ถูกต้องหรือหมดอายุ"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/provinces [get]
func GetProvince(c *gin.Context) {

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	lang := c.DefaultQuery("lang", "th")
	var column string
	switch lang {
	case "th":
		column = "local_name"
	case "en":
		column = "english_name"
	default:
		responses.BadRequest(c, "Invalid or missing lang parameter. Expected 'th' or 'en'")
		return
	}

	query := `SELECT * FROM ezw_get_provinces($1)`
	rows, err := db.Query(query, column)
	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch provinces")
		return
	}

	defer rows.Close()

	var provinces []dto.ProvinceResponse
	for rows.Next() {
		var province models.Province
		if err := rows.Scan(&province.ProvinceID, &province.ProvinceName); err != nil {
			responses.BadRequest(c, "Error processing data")
			return
		}

		provinces = append(provinces, dto.ProvinceResponse{
			ProvinceID:   province.ProvinceID,
			ProvinceName: province.ProvinceName,
		})
	}

	if len(provinces) == 0 {
		responses.OK(c, []dto.ProvinceResponse{})
		return
	}

	responses.OK(c, provinces)

}

// GetDistrict ดึงข้อมูลเขตหรืออำเภอ
// @Summary ดึงข้อมูลเขตหรืออำเภอตามจังหวัดที่เลือก
// @Description API นี้ใช้สำหรับดึงรายชื่อเขตหรืออำเภอตามจังหวัดที่กำหนด
// @Description
// @Description - **provinceId**: รหัสจังหวัด
// @Tags Dropdowns
// @Accept json
// @Produce json
// @Param lang query string false "เลือกภาษา (th หรือ en)" Enums(th, en) default(th)
// @Param province_id query string true "รหัสจังหวัด"
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.DistrictResponse] "การร้องขอข้อมูลเขตหรืออำเภอสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "คำขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "Token ไม่ถูกต้องหรือหมดอายุ"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/districts [get]
func GetDistrict(c *gin.Context) {

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	lang := c.DefaultQuery("lang", "th")
	var column string
	switch lang {
	case "th":
		column = "local_name"
	case "en":
		column = "english_name"
	default:
		responses.BadRequest(c, "Invalid or missing lang parameter. Expected 'th' or 'en'")
		return
	}

	provinceID := c.DefaultQuery("province_id", "")
	if provinceID == "" {
		responses.BadRequest(c, "Missing province_id parameter")
		return
	}

	if _, err := strconv.Atoi(provinceID); err != nil {
		responses.BadRequest(c, "Invalid province_id parameter. It must be a number")
		return
	}

	query := `SELECT * FROM ezw_get_districts($1, $2)`
	rows, err := db.Query(query, provinceID, column)
	if err != nil {
		responses.InternalServerError(c, "Failed to fetch districts")
		return
	}

	defer rows.Close()

	var districts []dto.DistrictResponse
	for rows.Next() {
		var district models.District
		if err := rows.Scan(&district.DistrictID, &district.DistrictName); err != nil {
			log.Println("Error scanning row:", err)
			responses.InternalServerError(c, "Error processing data")
			return
		}

		districts = append(districts, dto.DistrictResponse{
			DistrictID:   district.DistrictID,
			DistrictName: district.DistrictName,
		})
	}
	if len(districts) == 0 {
		responses.OK(c, []dto.DistrictResponse{})
		return
	}

	responses.OK(c, districts)

}

// GetSubDistrict ดึงข้อมูลตำบลหรือแขวง
// @Summary ดึงข้อมูลตำบลหรือแขวงตามเขตหรืออำเภอที่เลือก
// @Description API นี้ใช้สำหรับดึงรายชื่อตำบลหรือแขวงตามเขตหรืออำเภอที่กำหนด
// @Description
// @Description - **district_id**: รหัสเขตหรืออำเภอ
// @Tags Dropdowns
// @Accept json
// @Produce json
// @Param lang query string false "เลือกภาษา (th หรือ en)" Enums(th, en) default(th)
// @Param district_id query string true "รหัสเขตหรืออำเภอ"
// @Security Bearer
// @Success 200 {object} responses.SuccessResponseSwagger[[]dto.SubDistrictResponse] "การร้องขอข้อมูลตำบลหรือแขวงสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "คำขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "Token ไม่ถูกต้องหรือหมดอายุ"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/subdistricts [get]
func GetSubDistrict(c *gin.Context) {
	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	lang := c.DefaultQuery("lang", "th")
	var column string
	switch lang {
	case "th":
		column = "local_name"
	case "en":
		column = "english_name"
	default:
		responses.BadRequest(c, "Invalid or missing lang parameter. Expected 'th' or 'en'")
		return
	}

	districtID := c.DefaultQuery("district_id", "")
	if districtID == "" {
		responses.BadRequest(c, "Missing district_id parameter")
		return
	}

	if _, err := strconv.Atoi(districtID); err != nil {
		responses.BadRequest(c, "Invalid district_id parameter. It must be a number")
		return
	}

	query := `SELECT * FROM ezw_get_subdistricts($1, $2)`
	rows, err := db.Query(query, districtID, column)
	if err != nil {
		log.Println("Database query error:", err)
		responses.InternalServerError(c, "Failed to fetch sub-districts")
		return
	}
	defer rows.Close()

	// Process results
	var subdistricts []dto.SubDistrictResponse
	for rows.Next() {
		var subDistrict models.SubDistrict
		if err := rows.Scan(&subDistrict.SubDistrictID, &subDistrict.SubDistrictName, &subDistrict.ZipCode); err != nil {
			log.Println("Error scanning row:", err)
			responses.InternalServerError(c, "Error processing data")
			return
		}

		subdistricts = append(subdistricts, dto.SubDistrictResponse{
			SubDistrictID:   subDistrict.SubDistrictID,
			SubDistrictName: subDistrict.SubDistrictName,
			ZipCode:         subDistrict.ZipCode,
		})
	}

	if len(subdistricts) == 0 {
		responses.OK(c, []dto.SubDistrictResponse{})
		return
	}

	// Return response
	responses.OK(c, subdistricts)
}
