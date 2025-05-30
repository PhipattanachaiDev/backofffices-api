package gps_antenna

type EZWGpsAntennaStatusModelResponse struct {

	// สถานะไอดีของ Gps Antenna
	StatusId int `json:"status_id" example:"117001" extensions:"x-order=0" description:"สถานะไอดีของ Gps Antenna"`

	// ชื่อสถานะของ Gps Antenna
	StatusName string `json:"status_name" example:"คงอยู่" extensions:"x-order=1" description:"ชื่อสถานะของ Gps Antenna"`
}

type EZWSearchGpsAntennaModelRequest struct {

	// รหัสซีเรียลนัมเบอร์ Gps Antenna
	// required: true
	SerialNo string `json:"serial_no" example:"1" extensions:"x-order=0"  description:"รหัสซีเรียลนัมเบอร์ Gps Antenna"`

	// ไอดีสถานะ Gps Antenna
	// required: true
	StatusId int `json:"status_id" example:"117001" extensions:"x-order=1" description:"ไอดีสถานะ Gps Antenna"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"ทดสอบสร้าง" extensions:"x-order=2" description:"หมายเหตุ"`
}

type EZWSearchGpsAntennaModelResponse struct {

	// ไอดีของ Gps Antenna
	GpsAntennaId int `json:"gps_antenna_id" example:"1" extensions:"x-order=0" description:"ไอดีของ Gps Antenna"`

	// หมายเลขซีเรียลนัมเบอร์ Gps Antenna
	SerialNo string `json:"serial_no" example:"1234567890" extensions:"x-order=1" description:"หมายเลข Gps Antenna"`

	// ไอดีสถานะ Gps Antenna
	StatusId int `json:"status_id" example:"117001" extensions:"x-order=2" description:"ไอดีสถานะ Gps Antenna"`

	// รหัสสถานะ Gps Antenna
	StatusCode string `json:"status_code" example:"" extensions:"x-order=3" description:"รหัสสถานะ Gps Antenna"`

	// ชื่อสถานะ Gps Antenna
	StatusName string `json:"status_name" example:"คงอยู่" extensions:"x-order=4" description:"ชื่อสถานะ Gps Antenna"`

	// หมายเหตุ
	Remark string `json:"remark" example:"ทดสอบสร้าง 1" extensions:"x-order=5" description:"หมายเหตุ"`
}

type EZWUpdateGpsAntennaRequest struct {

	// ไอดี Gps Antenna
	// required: true
	GpsAntennaId int `json:"gps_antenna_id" example:"66672455949" extensions:"x-order=0"  description:"ไอดี Gps Antenna"`

	// หมายเลขอุปกรณ์ Gps Antenna
	// required: true
	SerialNo string `json:"serial_no" example:"0" extensions:"x-order=1" description:"หมายเลขอุปกรณ์ Gps Antenna"`

	// ไอดีสถานะ Gps Antenna
	// required: true
	StatusId int `json:"status_id" example:"0" extensions:"x-order=2" description:"ไอดีสถานะ Gps Antenna"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"สำหรับทดสอบ" extensions:"x-order=3" description:"หมายเหตุ"`
}

type EZWUpdateGpsAntennaResponse struct {

	// รายละเอียดตอบกลับ
	// required: true
	Message string `json:"remark" example:"eZView PLUS" extensions:"x-order=0" description:"รายละเอียดตอบกลับ"`
}

type EZWInsertGpsAntennaRequest struct {

	// หมายเลขอุปกรณ์ Gps Antenna
	// required: true
	SerialNo string `json:"serial_no" example:"0" extensions:"x-order=1" description:"หมายเลขอุปกรณ์ Gps Antenna"`

	// ไอดีสถานะ Gps Antenna
	// required: true
	StatusId int `json:"status_id" example:"0" extensions:"x-order=2" description:"ไอดีสถานะ Gps Antenna"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"สำหรับทดสอบ" extensions:"x-order=3" description:"หมายเหตุ"`
}

type EZWInsertGpsAntennaResponse struct {

	// รายละเอียดตอบกลับ
	// required: true
	Message string `json:"remark" example:"eZView PLUS" extensions:"x-order=0" description:"รายละเอียดตอบกลับ"`
}

type EZWInsertGpsAntennaGeneralRequest struct {
	// หมายเลขอุปกรณ์ Gps Antenna
	// required: true
	GpsAntennaID string `json:"gps_antenna_id" example:"0" extensions:"x-order=1" description:"ไอดี Gps Antenna"`
}

type GpsAntennaGeneralResponse struct {
	// ไอดีของ Gps Antenna
	GpsAntennaID *int `json:"gps_antenna_id" example:"1" extensions:"x-order=0" description:"ไอดีของ Gps Antenna"`

	// หมายเลขซีเรียลนัมเบอร์ Gps Antenna
	SerialNo *string `json:"serial_no" example:"123456789012345678901234567890" extensions:"x-order=1" description:"หมายเลข Gps Antenna"`

	// ไอดีสถานะ Gps Antenna
	StatusID *int `json:"status_id" example:"117001" extensions:"x-order=2" description:"ไอดีสถานะ Gps Antenna"`

	// รหัสสถานะ Gps Antenna
	StatusName *string `json:"status_name" example:"คงอยู่" extensions:"x-order=3" description:"ชื่อสถานะ Gps Antenna"`

	// หมายเหตุ
	Remark *string `json:"remark" example:"สำหรับทดสอบ" extensions:"x-order=4" description:"หมายเหตุ"`
}
