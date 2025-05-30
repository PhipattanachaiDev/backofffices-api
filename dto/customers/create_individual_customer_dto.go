package customers

type CreateIndividualCustomerRequest struct {
	// ชื่อ
	// required: true
	FirstName string `form:"first_name" binding:"required" example:"John" extensions:"x-order=x001"`

	// นามสกุล
	// required: true
	LastName string `form:"last_name" binding:"required" example:"Doe" extensions:"x-order=x002"`

	// ชื่อเรียกขาน
	// required: true
	CallName string `form:"call_name" binding:"required" example:"JD" extensions:"x-order=x003"`

	// เลขประจำตัวประชาชน
	// required: true
	IdCard string `form:"id_card" binding:"required" example:"1234567890123" extensions:"x-order=x004"`

	// เบอร์โทรศัพท์
	// required: true
	PhoneNumber string `form:"phone_number" binding:"required" example:"0812345678" extensions:"x-order=x005"`

	// ระดับลูกค้า
	// required: true
	LevelId int `form:"level_id" binding:"required" example:"1" extensions:"x-order=x006"`

	// ที่อยู่
	// required: true
	Address string `form:"address" binding:"required" example:"123 Main St" extensions:"x-order=x007"`

	// จังหวัด
	// required: true
	ProvinceId int `form:"province_id" binding:"required" example:"1" extensions:"x-order=x008"`

	// อำเภอ
	// required: true
	DistrictId int `form:"district_id" binding:"required" example:"1" extensions:"x-order=x009"`

	// ตำบล
	// required: true
	SubDistrictId int `form:"sub_district_id" binding:"required" example:"1" extensions:"x-order=x010"`

	// รหัสไปรษณีย์
	// required: true
	ZipCode string `form:"zip_code" binding:"required" example:"12345" extensions:"x-order=x011"`

	// ชื่อผู้ติดต่อ
	// required: true
	ContactName string `form:"contact_name" binding:"required" example:"John" extensions:"x-order=x012"`

	// นามสกุลผู้ติดต่อ
	// required: true
	ContactLastName string `form:"contact_lastname" binding:"required" example:"Doe" extensions:"x-order=x013"`

	// เบอร์โทรศัพท์ผู้ติดต่อ
	// required: true
	ContactPhoneNumber string `form:"contact_phone" binding:"required" example:"0812345678" extensions:"x-order=x014"`

	// อีเมลผู้ติดต่อ
	// required: true
	ContactEmail string `form:"contact_email" binding:"required" example:"test@mail.com" extensions:"x-order=x015"`

	// รายละเอียด
	Detail string `form:"detail" example:"รายละเอียดการลงทะเบียน" extensions:"x-order=x016"`

	// ละติจูด
	Latitude string `form:"latitude" example:"13.7563" extensions:"x-order=x017"`

	// ลองจิจูด
	Longitude string `form:"longitude" example:"100.5018" extensions:"x-order=x018"`

	// รหัสตัวแทนจำหน่าย
	DealerId int `form:"dealer_id" example:"1" extensions:"x-order=x019"`
}

type CreateIndividualCustomerResponse struct {
	// รหัสลูกค้า
	// example: 1
	CustomerId int `json:"customer_id" example:"1" extensions:"x-order=0"`
}
