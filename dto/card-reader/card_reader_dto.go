package card_reader

type EZWCardReaderStatusResponse struct {

	// ไอดีของสถานะเครื่องอ่านบัตร
	StatusId int `json:"status_id" example:"1" extensions:"x-order=0" description:"โมเดลไอดีของ GPS Antenna"`

	// ชื่อของสถานะเครื่องอ่านบัตร
	StatusName string `json:"status_name" example:"Active" extensions:"x-order=1" description:"ชื่อประเภทสถานะของ Gps Antenna"`
}

type EZWSearchCardReaderModelRequest struct {
	// หมายเลขเครื่องอ่านบัตร (Serial No)
	SerialNo string `json:"serial_no" example:"CR123456" extensions:"x-order=0" description:"หมายเลขเครื่องอ่านบัตร"`

	// ไอดีของแบรนด์เครื่องอ่านบัตร
	BrandId int `json:"brand_id" example:"1" extensions:"x-order=1" description:"ไอดีของแบรนด์เครื่องอ่านบัตร"`

	// ไอดีของรุ่นเครื่องอ่านบัตร
	ModelId int `json:"model_id" example:"1" extensions:"x-order=2" description:"ไอดีของรุ่นเครื่องอ่านบัตร"`

	// ไอดีสถานะของเครื่องอ่านบัตร (0, 99, หรือค่าอื่นๆ ตามเงื่อนไข)
	StatusId int `json:"status_id" example:"1" extensions:"x-order=3" description:"ไอดีสถานะของเครื่องอ่านบัตร"`

	// หมายเหตุเพิ่มเติมสำหรับการค้นหา
	Remark string `json:"remark" example:"Remark" extensions:"x-order=4" description:"หมายเหตุเพิ่มเติมสำหรับการค้นหา"`
}

// EZWSearchCardReaderModelResponse ใช้สำหรับส่งข้อมูลผลลัพธ์การค้นหา Card Reader ไปยัง client
type EZWSearchCardReaderModelResponse struct {
	// ไอดีของเครื่องอ่านบัตร
	CardReaderId int `json:"card_reader_id" example:"1" extensions:"x-order=0" description:"ไอดีของเครื่องอ่านบัตร"`

	// หมายเลขเครื่องอ่านบัตร
	SerialNo string `json:"serial_no" example:"CR123456" extensions:"x-order=1" description:"หมายเลขเครื่องอ่านบัตร"`

	// ไอดีของแบรนด์เครื่องอ่านบัตร
	BrandId int `json:"brand_id" example:"1" extensions:"x-order=2" description:"ไอดีของแบรนด์เครื่องอ่านบัตร"`

	// ชื่อแบรนด์ของเครื่องอ่านบัตร
	BrandName string `json:"brand_name" example:"BrandX" extensions:"x-order=3" description:"ชื่อแบรนด์ของเครื่องอ่านบัตร"`

	// ไอดีของรุ่นเครื่องอ่านบัตร
	ModelId *int `json:"model_id" example:"1" extensions:"x-order=4" description:"ไอดีของรุ่นเครื่องอ่านบัตร"`

	// ชื่อรุ่นของเครื่องอ่านบัตร
	ModelName *string `json:"model_name" example:"ModelY" extensions:"x-order=5" description:"ชื่อรุ่นของเครื่องอ่านบัตร"`

	// ไอดีสถานะของเครื่องอ่านบัตร
	StatusId int `json:"status_id" example:"1" extensions:"x-order=6" description:"ไอดีสถานะของเครื่องอ่านบัตร"`

	// ชื่อสถานะของเครื่องอ่านบัตร
	StatusName string `json:"status_name" example:"Active" extensions:"x-order=7" description:"ชื่อสถานะของเครื่องอ่านบัตร"`

	// หมายเหตุเพิ่มเติม
	Remark *string `json:"remark" example:"Remark" extensions:"x-order=7" description:"หมายเหตุเพิ่มเติม"`
}

type EZWCardReaderBrandResponse struct {
	// ไอดีของแบรนด์เครื่องอ่านบัตร
	BrandId int `json:"brand_id" example:"1" extensions:"x-order=0" description:"ไอดีของแบรนด์เครื่องอ่านบัตร"`

	// ชื่อแบรนด์เครื่องอ่านบัตร
	BrandName string `json:"brand_name" example:"BrandX" extensions:"x-order=1" description:"ชื่อแบรนด์ของเครื่องอ่านบัตร"`
}

type EZWCardReaderModelResponse struct {

	// ไอดีของรุ่นเครื่องอ่านบัตร
	ModelId int `json:"model_id" example:"1" extensions:"x-order=0" description:"ไอดีของรุ่นเครื่องอ่านบัตร"`

	// ชื่อรุ่นเครื่องอ่านบัตร
	ModelName string `json:"model_name" example:"ModelY" extensions:"x-order=1" description:"ชื่อรุ่นของเครื่องอ่านบัตร"`
}

type EZWUpdateCardReaderRequest struct {

	// ไอดี Card Reader
	// required: true
	CardReaderId int `json:"card_reader_id" example:"66672455949" extensions:"x-order=0"  description:"ไอดี Card Reader"`

	// ไอดีสถานะ Card Reader
	// required: true
	BrandId int `json:"brand_id" example:"0" extensions:"x-order=1" description:"ไอดีสถานะ Card Reader"`

	// ไอดีสถานะ Card Reader
	// required: true
	ModelId int `json:"model_id" example:"0" extensions:"x-order=2" description:"ไอดีสถานะ Card Reader"`

	// ไอดีสถานะ Card Reader
	// required: true
	StatusId int `json:"status_id" example:"0" extensions:"x-order=2" description:"ไอดีสถานะ Card Reader"`

	// หมายเลขอุปกรณ์ Card Reader
	// required: true
	SerialNo string `json:"serial_no" example:"0" extensions:"x-order=1" description:"หมายเลขอุปกรณ์ Card Reader"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"สำหรับทดสอบ" extensions:"x-order=3" description:"หมายเหตุ"`
}

type EZWUpdateCardReaderResponse struct {

	// รายละเอียดตอบกลับ
	// required: true
	Message string `json:"remark" example:"eZView PLUS" extensions:"x-order=0" description:"รายละเอียดตอบกลับ"`
}

type EZWInsertCardReaderRequest struct {

	// หมายเลขอุปกรณ์ Card Reader
	// required: true
	SerialNo string `json:"serial_no" example:"0" extensions:"x-order=0" description:"หมายเลขอุปกรณ์ Card Reader"`

	// ไอดีสถานะ Card Reader
	// required: true
	BrandId int `json:"brand_id" example:"0" extensions:"x-order=1" description:"ไอดีสถานะ Card Reader"`

	// ไอดีสถานะ Card Reader
	// required: true
	ModelId int `json:"model_id" example:"0" extensions:"x-order=2" description:"ไอดีสถานะ Card Reader"`

	// ไอดีสถานะ Card Reader
	// required: true
	StatusId int `json:"status_id" example:"0" extensions:"x-order=3" description:"ไอดีสถานะ Card Reader"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"สำหรับทดสอบ" extensions:"x-order=4" description:"หมายเหตุ"`
}

type EZWInsertCardReaderResponse struct {

	// รายละเอียดตอบกลับ
	// required: true
	Message string `json:"remark" example:"eZView PLUS" extensions:"x-order=0" description:"รายละเอียดตอบกลับ"`
}

type EZWCardReaderGeneralRequset struct {

	// ไอดีของเครื่องอ่านบัตร
	CardReaderId int `json:"card_reader_id" example:"1" extensions:"x-order=0" description:"ไอดีของเครื่องอ่านบัตร"`
}

type EZWCardReaderGeneralResponse struct {

	// ไอดีของเครื่องอ่านบัตร
	CardReaderID *int `json:"card_reader_id" extensions:"x-order=0" description:"ไอดีของเครื่องอ่านบัตร"`

	// หมายเลขเครื่องอ่านบัตร
	BrandID *int `json:"brand_id" extensions:"x-order=1" description:"หมายเลขเครื่องอ่านบัตร"`

	// ชื่อแบรนด์เครื่องอ่านบัตร
	BrandName *string `json:"brand_name" extensions:"x-order=2" description:"ชื่อแบรนด์เครื่องอ่านบัตร"`

	// ไอดีของรุ่นเครื่องอ่านบัตร
	ModelID *int `json:"model_id" extensions:"x-order=3" description:"ไอดีของรุ่นเครื่องอ่านบัตร"`

	// ชื่อรุ่นเครื่องอ่านบัตร
	ModelName *string `json:"model_name" extensions:"x-order=4" description:"ชื่อรุ่นเครื่องอ่านบัตร"`

	// หมายเลขเครื่องอ่านบัตร
	SerialNo *string `json:"serial_no" extensions:"x-order=5" description:"หมายเลขเครื่องอ่านบัตร"`

	// ไอดีสถานะของเครื่องอ่านบัตร
	StatusID *int `json:"status_id" extensions:"x-order=6" description:"ไอดีสถานะของเครื่องอ่านบัตร"`

	// ชื่อสถานะของเครื่องอ่านบัตร
	StatusName *string `json:"status_name" extensions:"x-order=7" description:"ชื่อสถานะของเครื่องอ่านบัตร"`

	// หมายเหตุ
	Remark *string `json:"remark" extensions:"x-order=8" description:"หมายเหตุ"`
}
