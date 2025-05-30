package models

type EZWGetCustomerById struct {
	CustomerId            int    // รหัสลูกค้า
	CustomerName          string // ชื่อลูกค้า
	CustomerLast          string // นามสกุลลูกค้า
	CallName              string // ชื่อเรียกขาน
	CustomerIdCard        string // หมายเลขบัตรประชาชน
	CustomerTel           string // เบอร์โทร
	LevelId               int    // รหัสระดับลูกค้า
	CustomerAddress       string // ที่อยู่
	CustomerProvinceId    int    // รหัสจังหวัด
	CustomerDistrictId    int    // รหัสอำเภอ
	CustomerSubDistrictId int    // รหัสตำบล
	CustomerZipCode       string // รหัสไปรษณีย์
	Detail                string // รายละเอียดลูกค้า
}
