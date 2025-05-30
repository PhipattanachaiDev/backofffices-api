package models

type EZWGetJuristicCustomerById struct {
	CustomerId            int    // รหัสลูกค้า
	JuristicName          string // ชื่อนิติบุคคล
	CallName              string // ชื่อเรียกขาน
	CustomerIdCard        string // หมายเลขบัตรประชาชน
	HeadOfficeNumber      string // เลขทะเบียนนิติบุคคล
	BranchNumber          string // เลขทะเบียนสาขา
	CustomerTel           string // เบอร์โทร
	LevelId               int    // รหัสระดับลูกค้า
	CustomerAddress       string // ที่อยู่
	CustomerProvinceId    int    // รหัสจังหวัด
	CustomerDistrictId    int    // รหัสอำเภอ
	CustomerSubDistrictId int    // รหัสตำบล
	CustomerZipCode       string // รหัสไปรษณีย์
	AuthorizedName        string // ชื่อผู้มีอำนาจลงนาม
	AuthorizedLastName    string // นามสกุลผู้มีอำนาจลงนาม
	AuthorizedPhoneNumber string // เบอร์โทรผู้มีอำนาจลงนาม
	Detail                string // รายละเอียดลูกค้า
}
