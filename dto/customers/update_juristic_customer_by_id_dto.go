package customers

type UpdateJuristicCustomerRequest struct {

	// ชื่อบริษัท
	// required: true
	JuristicName string `json:"juristic_name" binding:"required" example:"JD Co., Ltd." extensions:"x-order=001"`

	// ชื่อเรียกขาน
	// required: true
	CallName string `json:"call_name" binding:"required" example:"JD" extensions:"x-order=002"`

	// เลขประจำตัวประชาชน
	// required: true
	IdCard string `json:"id_card" binding:"required" example:"1234567890123" extensions:"x-order=003"`

	// หมาเลขสำนักงานใหญ่
	HeadOfficeNo string `json:"head_office_no" example:"123/45" extensions:"x-order=004"`

	// หมายเลขสาขา
	BranchNo string `json:"branch_no" example:"123/45" extensions:"x-order=005"`

	// เบอร์โทรศัพท์
	// required: true
	PhoneNumber string `json:"phone_number" binding:"required" example:"0812345678" extensions:"x-order=006"`

	// ระดับลูกค้า
	// required: true
	LevelId int `json:"level_id" binding:"required" example:"1" extensions:"x-order=007"`

	// ที่อยู่
	// required: true
	Address string `json:"address" binding:"required" example:"123 Main St" extensions:"x-order=008"`

	// จังหวัด
	// required: true
	ProvinceId int `json:"province_id" binding:"required" example:"1" extensions:"x-order=009"`

	// อำเภอ
	// required: true
	DistrictId int `json:"district_id" binding:"required" example:"1" extensions:"x-order=010"`

	// ตำบล
	// required: true
	SubDistrictId int `json:"sub_district_id" binding:"required" example:"1" extensions:"x-order=011"`

	// รหัสไปรษณีย์
	// required: true
	ZipCode string `json:"zip_code" binding:"required" example:"12345" extensions:"x-order=012"`

	// ชื่อผู้ลงนาม
	AuthorizedName string `json:"authorized_name" example:"John Doe" extensions:"x-order=013"`

	// นามสกุลผู้ลงนาม
	AuthorizedLastName string `json:"authorized_last_name" example:"Doe" extensions:"x-order=014"`

	// เบอร์โทรศัพท์ผู้ลงนาม
	AuthorizedPhoneNumber string `json:"authorized_phone_number" example:"0812345678" extensions:"x-order=015"`

	// ละติจูด
	Latitude string `json:"latitude" example:"13.7563" extensions:"x-order=016"`

	// ลองจิจูด
	Longitude string `json:"longitude" example:"100.5018" extensions:"x-order=017"`

	// รายละเอียดลูกค้า
	Detail string `json:"detail" example:"Customer details" extensions:"x-order=018"`
}
