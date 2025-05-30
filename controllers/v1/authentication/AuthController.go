package AuthControllers

import (
	"database/sql"
	"strconv"

	"log"
	"os"

	"time"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	dto "ezview.asia/ezview-web/ezview-lite-back-office/dto/authentication"
	"ezview.asia/ezview-web/ezview-lite-back-office/middlewares"
	models "ezview.asia/ezview-web/ezview-lite-back-office/models/authentication"
	lineNotificationService "ezview.asia/ezview-web/ezview-lite-back-office/services/lineNotificationService"
	responses "ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Load environment variables from .env file
func init() {
	if os.Getenv("GIN_MODE") != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}
}

// Retrieve JWT key from environment variable
var jwtKey = []byte(os.Getenv("JWT_KEY"))

// var encryptionKey string = os.Getenv("ENCRYPTION_KEY")

// Login ใช้สำหรับเข้าสู่ระบบ
// @Summary เข้าสู่ระบบ
// @Description API นี้ใช้สำหรับการเข้าสู่ระบบของผู้ใช้และคืนค่า JWT token ซึ่ง Response จะมีข้อมูลดังนี้:
// @Description
// @Description - **token**: JWT token ที่ได้รับจากการเข้าสู่ระบบ
// @Description - **expires_at**: เวลาที่ token หมดอายุ
// @Description - **expires_after**: ระยะเวลาที่ token หมดอายุหลังจากนี้ (60 นาที)
// @Description
// @Description **หมายเหตุ:** กรุณาใช้ข้อมูลของ User Name และ Password ที่ถูกต้องเพื่อรับ Token ในการใช้งาน API อื่น ๆ
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginRequest body dto.LoginRequest true "Login Request"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.LoginResponse] "การเข้าสู่ระบบสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/login [post]
func Login(c *gin.Context) {
	var reqBody dto.LoginRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	username := reqBody.Username
	password := reqBody.Password

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	// ใช้ Prepared Statement (แบบไม่ต้องเรียก `Prepare()` เอง)
	query := `SELECT * FROM ezw_authentication_backoffice($1, $2)`
	row := db.QueryRow(query, username, password)

	var user models.EZWAuthenticationModel
	err := row.Scan(&user.UserId, &user.UserName, &user.RoleId, &user.RoleName, &user.Access)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.Unauthorized(c, "Invalid credentials")
		} else {
			log.Printf("Error executing query: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
		}
		return
	}

	// สร้าง JWT Token
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &middlewares.Claims{
		UserId:   user.UserId,
		Username: user.UserName,
		RoleId:   user.RoleId,
		RoleName: user.RoleName,
		Access:   user.Access,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		responses.InternalServerError(c, "Failed to generate token")
		return
	}

	// ตอบกลับข้อมูล
	loginResponse := dto.LoginResponse{
		Token:        "Bearer " + tokenString,
		ExpiresAt:    expirationTime.Format(time.RFC3339),
		ExpiresAfter: "60 minutes",
	}

	responses.OK(c, loginResponse)
}

// RefreshToken ใช้สำหรับรีเฟรช JWT token
// @Summary รีเฟรช JWT token
// @Description API นี้ใช้สำหรับรีเฟรช JWT token โดยใช้ Refresh token ที่ได้รับจากการเข้าสู่ระบบเพื่อสร้าง Access token ใหม่ ซึ่ง Response จะมีข้อมูลดังนี้:
// @Description
// @Description - **token**: JWT token ที่ได้รับจากการรีเฟรช
// @Description - **expires_at**: เวลาที่ token หมดอายุ
// @Description
// @Tags Authentication
// @Accept json
// @Produce json
// @Param refreshToken body dto.RefreshTokenRequest true "Refresh Token Request"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.RefreshTokenResponse] "การรีเฟรช Token สำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 401 {object} responses.UnauthorizedResponseSwagger "ข้อมูลเข้าสู่ระบบไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/refresh-token [post]
func RefreshToken(c *gin.Context) {
	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
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
	query := `SELECT * FROM ezw_authentication_refresh_backoffice($1)`
	row := db.QueryRow(query, userID)

	var user models.EZWAuthenticationModel
	err := row.Scan(&user.UserId, &user.UserName, &user.RoleId, &user.RoleName, &user.Access)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.Unauthorized(c, "Invalid credentials")
		} else {
			log.Printf("Error executing query: %v", err)
			responses.InternalServerError(c, "Internal Server Error")
		}
		return
	}

	// สร้าง JWT Token
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &middlewares.Claims{
		UserId:   user.UserId,
		Username: user.UserName,
		RoleId:   user.RoleId,
		RoleName: user.RoleName,
		Access:   user.Access,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		responses.InternalServerError(c, "Failed to generate token")
		return
	}

	// ตอบกลับข้อมูล
	loginResponse := dto.LoginResponse{
		Token:        "Bearer " + tokenString,
		ExpiresAt:    expirationTime.Format(time.RFC3339),
		ExpiresAfter: "60 minutes",
	}

	responses.OK(c, loginResponse)
}

// Logout invalidates JWT tokens
func Logout(c *gin.Context) {
	// Optionally implement token invalidation logic here
	responses.OK(c, gin.H{"message": "Logged out successfully"})
	lineNotificationService.SendNotification("User logged out successfully")
}

// Register ใช้สำหรับลงทะเบียนผู้ใช้ใหม่สำหรับลูกค้า
// @Summary ลงทะเบียนผู้ใช้ใหม่สำหรับลูกค้า
// @Description API นี้ใช้สำหรับการลงทะเบียนผู้ใช้ใหม่สำหรับลูกค้าและคืนค่า JWT token ซึ่ง Response จะมีข้อมูลดังนี้:
// @Description
// @Description - **token**: JWT token ที่ได้รับจากการลงทะเบียน
// @Description - **expires_at**: เวลาที่ token หมดอายุ
// @Description - **expires_after**: ระยะเวลาที่ token หมดอายุหลังจากนี้ (60 นาที)
// @Description
// @Description Customer ID ได้จาก API /v1/customers/individual หรือ /v1/customers/company
// @Description
// @Description **หมายเหตุ:** กรุณาใช้ข้อมูลของ Customer ID, User Name และ Password ที่ถูกต้องเพื่อลงทะเบียนผู้ใช้ใหม่
// @Tags Authentication
// @Accept json
// @Produce json
// @Param registerRequest body dto.RegisterRequest true "Register Request"
// @Success 200 {object} responses.SuccessResponseSwagger[dto.LoginResponse] "การลงทะเบียนสำเร็จ"
// @Failure 400 {object} responses.BadRequestResponseSwagger "การร้องขอไม่ถูกต้อง"
// @Failure 500 {object} responses.InternalServerErrorResponseSwagger "ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /v1/auth/register/user/customer [post]
func RegisterUserCustomer(c *gin.Context) {
	// Retrieve user ID from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		responses.BadRequest(c, "User ID not found in context")
		return
	}

	var reqBody dto.RegisterRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		responses.BadRequest(c, "Invalid JSON body")
		return
	}

	customerId := reqBody.CustomerId
	username := reqBody.Username
	password := reqBody.Password
	name := reqBody.Name
	description := reqBody.Description

	// encrypted := cryptography.Encryption(password, encryptionKey)

	// ดึง instance ของ database (ไม่เปิด connection ใหม่)
	db := database.GetDB()

	// ตรวจสอบ connection
	if db == nil {
		log.Println("Database connection not available")
		responses.InternalServerError(c, "Database connection error")
		return
	}

	query := `SELECT * FROM ezw_register_user_customer($1, $2, $3, $4, $5, $6)`
	row := db.QueryRow(query, username, password, name, description, customerId, userID)

	var user_id int
	err := row.Scan(&user_id)

	if err != nil {
		// Check for specific error message
		if err.Error() == "pq: Login name "+username+" is already taken." {
			responses.BadRequest(c, "Login name already exists")
			return
		}

		if err.Error() == "pq: Customer ID "+strconv.Itoa(customerId)+" is not valid." {
			responses.BadRequest(c, "Customer ID does not exist")
			return
		}

		log.Printf("Error executing query: %v", err)
		responses.InternalServerError(c, "Internal Server Error")
		return
	}

	responses.Created(c)
}
