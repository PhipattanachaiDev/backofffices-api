package gsm_antenna

type EZWGsmAntennaStatusModelResponse struct {

	// สถานะไอดีของ Gsm Antenna
	StatusId int `json:"status_id" example:"120001" extensions:"x-order=0" description:"สถานะไอดีของ Gsm Antenna"`

	// ชื่อสถานะของ Gsm Antenna
	StatusName string `json:"status_name" example:"คงอยู่" extensions:"x-order=1" description:"ชื่อสถานะของ Gsm Antenna"`
}

type EZWSearchGsmAntennaModelRequest struct {

	// รหัสซีเรียลนัมเบอร์ Gsm Antenna
	// required: true
	SerialNo string `json:"serial_no" example:"1421412421" extensions:"x-order=0"  description:"รหัสซีเรียลนัมเบอร์ Gsm Antenna"`

	// ไอดีสถานะ Gsm Antenna
	// required: true
	StatusId int `json:"status_id" example:"120002" extensions:"x-order=1" description:"ไอดีสถานะ Gsm Antenna"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"ยกเลิก" extensions:"x-order=2" description:"หมายเหตุ"`
}

type EZWSearchGsmAntennaModelResponse struct {

	// ไอดีของ Gsm Antenna
	GsmAntennaId int `json:"gsm_antenna_id" example:"1" extensions:"x-order=0" description:"ไอดีของ Gsm Antenna"`

	// หมายเลขซีเรียลนัมเบอร์ Gsm Antenna
	SerialNo string `json:"serial_no" example:"123456789012345678901234567890" extensions:"x-order=1" description:"หมายเลข Gsm Antenna"`

	// ไอดีสถานะ Gsm Antenna
	StatusId int `json:"status_id" example:"120001" extensions:"x-order=2" description:"ไอดีสถานะ Gsm Antenna"`

	// รหัสสถานะ Gsm Antenna
	StatusCode string `json:"status_code" example:"คงอยู่" extensions:"x-order=3" description:"รหัสสถานะ Gsm Antenna"`

	// ชื่อสถานะ Gsm Antenna
	StatusName string `json:"status_name" example:"Active" extensions:"x-order=4" description:"ชื่อสถานะ Gsm Antenna"`

	// หมายเหตุ
	Remark string `json:"remark" example:"eZView PLUS" extensions:"x-order=5" description:"หมายเหตุ"`
}

type EZWUpdateGsmAntennaRequest struct {

	// ไอดี Gsm Antenna
	// required: true
	GsmAntennaId int `json:"gsm_antenna_id" example:"66672455949" extensions:"x-order=0"  description:"ไอดี Gsm Antenna"`

	// หมายเลขอุปกรณ์ Gsm Antenna
	// required: true
	SerialNo string `json:"serial_no" example:"0" extensions:"x-order=1" description:"หมายเลขอุปกรณ์ Gsm Antenna"`

	// ไอดีสถานะ Gsm Antenna
	// required: true
	StatusId int `json:"status_id" example:"0" extensions:"x-order=2" description:"ไอดีสถานะ Gsm Antenna"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"สำหรับทดสอบ" extensions:"x-order=3" description:"หมายเหตุ"`
}

type EZWUpdateGsmAntennaResponse struct {

	// รายละเอียดตอบกลับ
	// required: true
	Message string `json:"remark" example:"eZView PLUS" extensions:"x-order=0" description:"รายละเอียดตอบกลับ"`
}

type EZWInsertGsmAntennaRequest struct {

	// หมายเลขอุปกรณ์ Gsm Antenna
	// required: true
	SerialNo string `json:"serial_no" example:"0" extensions:"x-order=1" description:"หมายเลขอุปกรณ์ Gsm Antenna"`

	// ไอดีสถานะ Gsm Antenna
	// required: true
	StatusId int `json:"status_id" example:"0" extensions:"x-order=2" description:"ไอดีสถานะ Gsm Antenna"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"สำหรับทดสอบ" extensions:"x-order=3" description:"หมายเหตุ"`
}

type EZWInsertGsmAntennaResponse struct {

	// รายละเอียดตอบกลับ
	// required: true
	Message string `json:"remark" example:"eZView PLUS" extensions:"x-order=0" description:"รายละเอียดตอบกลับ"`
}

type EZWSearchGsmAntennaRequest struct {

	// ไอดี Gsm Antenna
	// required: true
	GsmAntennaId *int `json:"gsm_antenna_id" description:"ไอดี Gsm Antenna"`
}

type EZGsmAntennaGeneralRequest struct {
	// ไอดี Gsm Antenna
	// required: true
	GsmAntennaID *int `json:"gsm_antenna_id" extensions:"x-order=0" description:"ไอดี Gsm Antenna"`
}

type EZGsmAntennaGeneralResponse struct {
	// ไอดี Gsm Antenna
	GsmAntennaID *int `json:"gsm_antenna_id" extensions:"x-order=0" description:"ไอดี Gsm Antenna"`

	// ซีเรียลนัมเบอร์ Gsm Antenna
	SerialNo *string `json:"serial_no" extensions:"x-order=1" description:"ซีเรียลนัมเบอร์ Gsm Antenna"`

	// ไอดีสถานะ Gsm Antenna
	StatusID *int `json:"status_id" extensions:"x-order=2" description:"ไอดีสถานะ Gsm Antenna"`

	// ชื่อสถานะ Gsm Antenna
	StatusName *string `json:"status_name" extensions:"x-order=3" description:"ชื่อสถานะ Gsm Antenna"`

	// หมายเหตุ
	Remark *string `json:"remark" extensions:"x-order=4" description:"หมายเหตุ"`
}
