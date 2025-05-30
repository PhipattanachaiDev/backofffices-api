package customers

type UpdateCustomerContactRequest struct {
	ContactId       int    `json:"contact_id"`        // รหัสผู้ติดต่อ
	ContactName     string `json:"contact_name"`      // ชื่อผู้ติดต่อ
	ContactLastName string `json:"contact_last_name"` // นามสกุลผู้ติดต่อ
	ContactPhone    string `json:"contact_phone"`     // เบอร์โทรศัพท์
	ContactEmail    string `json:"contact_email"`     // อีเมล
	ContactType     string `json:"contact_type"`      // ประเภทผู้ติดต่อ
}
