package authentication

type LoginRequest struct {

	// ชื่อผู้ใช้ที่ใช้ในการเข้าสู่ระบบ
	// required: true
	Username string `json:"username" binding:"required" example:"johndoe" extensions:"x-order=0"`

	// รหัสผ่านที่ใช้ในการเข้าสู่ระบบ
	// required: true
	Password string `json:"password" binding:"required" example:"password123" extensions:"x-order=1"`
}

type LoginResponse struct {
	// โทเค็นที่ใช้ในการยืนยันตัวตน
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJqb2huX2RvZSIsImV4cCI6MTY2NzQ2OTg4NywiaWF0IjoxNjY3NDY5NzI3fQ.SvJv9_XERK8TSEhK9y7BX34vxIZYz3ndY8bX_MSRZGQ" extensions:"x-order=2" description:"โทเค็นที่ใช้ในการยืนยันตัวตนเพื่อเข้าถึงระบบ"`

	// เวลาที่โทเค็นหมดอายุ
	ExpiresAt string `json:"expires_at" example:"2024-07-19T09:43:39Z" extensions:"x-order=3" description:"เวลาที่โทเค็นหมดอายุ (รูปแบบ ISO 8601)"`

	// ระยะเวลาที่โทเค็นจะหมดอายุหลังจากการสร้าง
	ExpiresAfter string `json:"expires_after" example:"5 minutes" extensions:"x-order=4" description:"ระยะเวลาที่โทเค็นจะหมดอายุหลังจากการสร้าง (เช่น 5 นาที, 1 ชั่วโมง)"`
}
