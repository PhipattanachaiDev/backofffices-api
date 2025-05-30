package models

type EZWGetCustomerContactByCustomerId struct {
	ContactId       int    // รหัสผู้ติดต่อ
	ContactName     string // ชื่อผู้ติดต่อ
	ContactLastName string // นามสกุลผู้ติดต่อ
	ContactPhone    string // เบอร์โทรศัพท์
	ContactEmail    string // อีเมล
	ContactType     string // ประเภทผู้ติดต่อ
}
