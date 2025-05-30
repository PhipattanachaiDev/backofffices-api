package battery

type EZWBatteryStatusModelResponse struct {

	// สถานะไอดีของ Battery
	StatusId int `json:"status_id" example:"100001" extensions:"x-order=0" description:"สถานะไอดีของ Battery"`

	// ชื่อสถานะของ Battery
	StatusName string `json:"status_name" example:"Active" extensions:"x-order=1" description:"ชื่อสถานะของ Battery"`
}

type EZWSearchBatteryModelRequest struct {

	// รหัสซีเรียลนัมเบอร์ Battery
	// required: true
	SerialNo string `json:"serial_no" example:"1" extensions:"x-order=0"  description:"รหัสซีเรียลนัมเบอร์ Battery"`

	// ไอดีสถานะ Battery
	// required: true
	StatusId int `json:"status_id" example:"100001" extensions:"x-order=1" description:"ไอดีสถานะ Battery"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"ทดสอบ" extensions:"x-order=2" description:"หมายเหตุ"`
}

type EZWSearchBatteryModelResponse struct {

	// ไอดีของแบตเตอรี่
	BatteryId int `json:"battery_id" example:"1" extensions:"x-order=0" description:"ไอดีของแบตเตอรี่"`

	// หมายเลขซีเรียลนัมเบอร์แบตเตอรี่
	SerialNo string `json:"serial_no" example:"123456" extensions:"x-order=1" description:"หมายเลขซีเรียลนัมเบอร์แบตเตอรี่"`

	// ไอดีสถานะแบตเตอรี่
	StatusId int `json:"status_id" example:"100001" extensions:"x-order=2" description:"ไอดีสถานะแบตเตอรี่"`

	// ชื่อสถานะแบตเตอรี่
	StatusName string `json:"status_name" example:"Active" extensions:"x-order=3" description:"ชื่อสถานะแบตเตอรี่"`

	// หมายเหตุ
	Remark string `json:"remark" example:"ทดสอบ" extensions:"x-order=4" description:"หมายเหตุ"`
}

type EZWUpdateBatteryRequest struct {

	// ไอดี Battery
	// required: true
	BatteryId int `json:"battery_id" example:"1" extensions:"x-order=0"  description:"ไอดี Battery"`

	// หมายเลขอุปกรณ์ Battery
	// required: true
	SerialNo string `json:"serial_no" example:"0" extensions:"x-order=1" description:"หมายเลขอุปกรณ์ Battery"`

	// ไอดีสถานะ Battery
	// required: true
	StatusId int `json:"status_id" example:"0" extensions:"x-order=2" description:"ไอดีสถานะ Battery"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"สำหรับทดสอบ" extensions:"x-order=3" description:"หมายเหตุ"`
}

type EZWUpdateBatteryResponse struct {

	// รายละเอียดตอบกลับ
	// required: true
	Message string `json:"remark" example:"eZView PLUS" extensions:"x-order=0" description:"รายละเอียดตอบกลับ"`
}

type EZWInsertBatteryRequest struct {

	// หมายเลขอุปกรณ์ Battery
	// required: true
	SerialNo string `json:"serial_no" example:"0" extensions:"x-order=1" description:"หมายเลขอุปกรณ์ Battery"`

	// ไอดีสถานะ Battery
	// required: true
	StatusId int `json:"status_id" example:"0" extensions:"x-order=2" description:"ไอดีสถานะ Battery"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"สำหรับทดสอบ" extensions:"x-order=3" description:"หมายเหตุ"`
}

type EZWInsertBatteryResponse struct {

	// รายละเอียดตอบกลับ
	// required: true
	Message string `json:"message" example:"eZView PLUS" extensions:"x-order=0" description:"รายละเอียดตอบกลับ"`
}

type EZWGetBatteryGeneralRequest struct {
	// ไอดีของแบตเตอรี่
	BatteryId *int `json:"battery_id" extensions:"x-order=0" description:"ไอดีของแบตเตอรี่"`
}

type EZWGetBatteryGeneralResponse struct {
	// ไอดีของแบตเตอรี่
	BatteryID *int `json:"battery_id" extensions:"x-order=0" description:"ไอดีของแบตเตอรี่"`
	// หมายเลขซีเรียลนัมเบอร์แบตเตอรี่
	SerialNo *string `json:"serial_no" extensions:"x-order=1" description:"หมายเลขซีเรียลนัมเบอร์แบตเตอรี่"`
	// ไอดีสถานะแบตเตอรี่
	StatusID *int `json:"status_id" extensions:"x-order=2" description:"ไอดีสถานะแบตเตอรี่"`
	// ชื่อสถานะแบตเตอรี่
	StatusName *string `json:"status_name" extensions:"x-order=3" description:"ชื่อสถานะแบตเตอรี่"`
	// หมายเหตุ
	Remark *string `json:"remark" extensions:"x-order=4" description:"หมายเหตุ"`
}
