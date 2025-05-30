package customers

type GetJuristicCustomerByIdResponse struct {
	// รหัสลูกค้า
	CustomerId int `json:"customer_id"`

	// ชื่อนิติบุคคล
	JuristicName string `json:"juristic_name"`

	// ชื่อเรียกขาน
	CallName string `json:"call_name"`

	// เลขทะเบียนนิติบุคคล
	CustomerIdCard string `json:"customer_id_card"`

	// หมายเลขสำนักงานใหญ่
	HeadOfficeNumber string `json:"head_office_number"`

	// เลขทะเบียนสาขา
	BranchNumber string `json:"branch_number"`

	// เบอร์โทร
	CustomerTel string `json:"customer_tel"`

	// รหัสระดับลูกค้า
	LevelId int `json:"level_id"`

	// ที่อยู่
	CustomerAddress string `json:"customer_address"`

	// จังหวัด
	CustomerProvinceId int `json:"customer_province_id"`

	// อำเภอ
	CustomerDistrictId int `json:"customer_district_id"`

	// ตำบล
	CustomerSubDistrictId int `json:"customer_sub_district_id"`

	// รหัสไปรษณีย์
	CustomerZipCode string `json:"customer_zip_code"`

	// ชื่อผู้มีอำนาจลงนาม
	AuthorizedName string `json:"authorized_name"`

	// นามสกุลผู้มีอำนาจลงนาม
	AuthorizedLastName string `json:"authorized_last_name"`

	// เบอร์โทรผู้มีอำนาจลงนาม
	AuthorizedPhoneNumber string `json:"authorized_phone_number"`

	// รายละเอียดลูกค้า
	Detail string `json:"detail"`
}
