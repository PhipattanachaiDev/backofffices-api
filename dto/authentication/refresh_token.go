package authentication

type RefreshTokenRequest struct {
	// โทเค็นที่ต้องการ Refresh เวลาการใช้งาน
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJqb2huX2RvZSIsImV4cCI6MTY2NzQ2OTg4NywiaWF0IjoxNjY3NDY5NzI3fQ.SvJv9_XERK8TSEhK9y7BX34vxIZYz3ndY8bX_MSRZGQ" extensions:"x-order=0" description:"โทเค็นที่ใช้ในการยืนยันตัวตนเพื่อเข้าถึงระบบ"`
}
type RefreshTokenResponse struct {
	// โทเค็นที่ใช้ในการยืนยันตัวตน
	Token string `json:"token" extensions:"x-order=1" description:"โทเค็นที่ใช้ในการยืนยันตัวตนเพื่อเข้าถึงระบบ"`

	// เวลาที่โทเค็นหมดอายุ
	ExpiresAt string `json:"expires_at" example:"2024-07-19T09:43:39Z" extensions:"x-order=2" description:"เวลาที่โทเค็นหมดอายุ (รูปแบบ ISO 8601)"`
}
