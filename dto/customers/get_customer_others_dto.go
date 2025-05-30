package customers

type GetCustomerOthersByCustomerIdResponse struct {
	// สร้างโดย
	CreatedBy string `json:"created_by"`

	// วันที่สร้าง
	CreatedAt string `json:"created_at"`

	// อัพเดทโดย
	UpdatedBy string `json:"updated_by"`

	// วันที่อัพเดท
	UpdatedAt string `json:"updated_at"`
}
