package customers

type GetCustomerByIdResponse struct {
	// รหัสลูกค้า
	CustomerId int `json:"customer_id"`

	// ชื่อลูกค้า
	CustomerName string `json:"customer_name"`

	// นามสกุลลูกค้า
	CustomerLast string `json:"customer_lastname"`

	// ชื่อเรียกขาน
	CallName string `json:"call_name"`

	// หมายเลขบัตรประชาชน
	CustomerIdCard string `json:"customer_id_card"`

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

	// รายละเอียดลูกค้า
	Detail string `json:"detail"`
}
