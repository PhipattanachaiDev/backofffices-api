package customers

type CreateJuristicCustomerRequest struct {
	// ชื่อนิติบุคคล
	// required: true
	JuristicName string `form:"juristic_name" binding:"required" example:"บริษัท อีซีวิว จำกัด" extensions:"x-order=x001"`

	// ชื่อเรียกขาน
	// required: true
	CallName string `form:"call_name" binding:"required" example:"อีซีวิว" extensions:"x-order=x002"`

	// เลขประจำตัวผู้เสียภาษี
	// required: true
	TaxID string `form:"tax_id" binding:"required" example:"1234567890123" extensions:"x-order=x003"`

	// หมายเลขสำนักงานใหญ่
	HeadOfficeNumber string `form:"headoffice_no" example:"1234567890123" extensions:"x-order=x004"`

	// หมายเลขสาขา
	BranchNumber string `form:"branch_no" example:"1234567890123" extensions:"x-order=x005"`

	// เบอร์โทรศัพท์
	// required: true
	PhoneNumber string `form:"phone" binding:"required" example:"0812345678" extensions:"x-order=x006"`

	// ระดับลูกค้า
	// required: true
	LevelId int `form:"level_id" binding:"required" example:"1" extensions:"x-order=x007"`

	// ที่อยู่บริษัท
	// required: true
	Address string `form:"address" binding:"required" example:"123 ถนน สุขุมวิท แขวง คลองเตย เขต คลองเตย กรุงเทพมหานคร 10110" extensions:"x-order=x008"`

	// จังหวัด
	// required: true
	ProvinceId int `form:"province_id" binding:"required" example:"1" extensions:"x-order=x009"`

	// อำเภอ
	// required: true
	DistrictId int `form:"district_id" binding:"required" example:"1" extensions:"x-order=x010"`

	// ตำบล
	// required: true
	SubDistrictId int `form:"sub_district_id" binding:"required" example:"1" extensions:"x-order=x011"`

	// รหัสไปรษณีย์
	// required: true
	ZipCode string `form:"zip_code" binding:"required" example:"10110" extensions:"x-order=x012"`

	// ชื่อผู้ติดต่อ
	// required: true
	ContactName string `form:"contact_name" binding:"required" example:"นายสมชาย" extensions:"x-order=x013"`

	// นามสกุลผู้ติดต่อ
	// required: true
	ContactLastName string `form:"contact_lastname" binding:"required" example:"ใจดี" extensions:"x-order=x014"`

	// เบอร์โทรศัพท์ผู้ติดต่อ
	// required: true
	ContactPhoneNumber string `form:"contact_phone" binding:"required" example:"0812345678" extensions:"x-order=x015"`

	// อีเมลผู้ติดต่อ
	// required: true
	ContactEmail string `form:"contact_email" binding:"required" example:"test@mail.com" extensions:"x-order=x016"`

	// ชื่อผู้ลงนาม
	// required: true
	AuthorizedName string `form:"authorized_name" binding:"required" example:"นายสมชาย" extensions:"x-order=x017"`

	// นามสกุลผู้ลงนาม
	// required: true
	AuthorizedLastName string `form:"authorized_lastname" binding:"required" example:"ใจดี" extensions:"x-order=x018"`

	// เบอร์โทรศัพท์ผู้ลงนาม
	// required: true
	AuthorizedPhoneNumber string `form:"authorized_phone" binding:"required" example:"0812345678" extensions:"x-order=x019"`

	// รายละเอียด
	Detail string `form:"detail" example:"รายละเอียดการลงทะเบียน" extensions:"x-order=x020"`

	// ละติจูด
	Latitude string `form:"latitude" example:"13.7563" extensions:"x-order=x021"`

	// ลองจิจูด
	Longitude string `form:"longitude" example:"100.5018" extensions:"x-order=x022"`

	// รหัสตัวแทนจำหน่าย
	DealerId int `form:"dealer_id" example:"1" extensions:"x-order=x023"`
}

type CreateJuristicCustomerResponse struct {
	// รหัสลูกค้า
	// example: 1
	CustomerId int `json:"customer_id" example:"1" extensions:"x-order=0"`
}
