package customers

type UpdateIndividualCustomerRequest struct {
	// ชื่อ
	// required: true
	FirstName string `json:"first_name" binding:"required" example:"John" extensions:"x-order=x001"`

	// นามสกุล
	// required: true
	LastName string `json:"last_name" binding:"required" example:"Doe" extensions:"x-order=x002"`

	// ชื่อเรียกขาน
	// required: true
	CallName string `json:"call_name" binding:"required" example:"JD" extensions:"x-order=x003"`

	// เลขประจำตัวประชาชน
	// required: true
	IdCard string `json:"id_card" binding:"required" example:"1234567890123" extensions:"x-order=x004"`

	// เบอร์โทรศัพท์
	// required: true
	PhoneNumber string `json:"phone_number" binding:"required" example:"0812345678" extensions:"x-order=x005"`

	// ระดับลูกค้า
	// required: true
	LevelId int `json:"level_id" binding:"required" example:"1" extensions:"x-order=x006"`

	// ที่อยู่
	// required: true
	Address string `json:"address" binding:"required" example:"123 Main St" extensions:"x-order=x007"`

	// จังหวัด
	// required: true
	ProvinceId int `json:"province_id" binding:"required" example:"1" extensions:"x-order=x008"`

	// อำเภอ
	// required: true
	DistrictId int `json:"district_id" binding:"required" example:"1" extensions:"x-order=x009"`

	// ตำบล
	// required: true
	SubDistrictId int `json:"sub_district_id" binding:"required" example:"1" extensions:"x-order=x010"`

	// รหัสไปรษณีย์
	// required: true
	ZipCode string `json:"zip_code" binding:"required" example:"12345" extensions:"x-order=x011"`

	// ละติจูด
	Latitude string `json:"latitude" example:"13.7563" extensions:"x-order=x012"`

	// ลองจิจูด
	Longitude string `json:"longitude" example:"100.5018" extensions:"x-order=x013"`

	// รายละเอียดลูกค้า
	Detail string `json:"detail" example:"Customer details" extensions:"x-order=x014"`
}
