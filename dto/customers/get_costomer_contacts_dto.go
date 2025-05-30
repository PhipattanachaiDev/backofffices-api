package customers

type GetCustomerContactsResponse struct {
	// รหัสผู้ติดต่อ
	ContactId int `json:"contact_id"`

	// ชื่อผู้ติดต่อ
	ContactName string `json:"contact_name"`

	// นามสกุลผู้ติดต่อ
	ContactLastName string `json:"contact_last_name"`

	// เบอร์โทรศัพท์
	ContactPhone string `json:"contact_phone"`

	// อีเมล
	ContactEmail string `json:"contact_email"`

	// ประเภทผู้ติดต่อ
	ContactType string `json:"contact_type"`
}
