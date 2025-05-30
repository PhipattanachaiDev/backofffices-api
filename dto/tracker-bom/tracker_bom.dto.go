package trackerbom

type EZWGetTrackerBomResponse struct {

	// ไอดีของ Tracker Bom
	TrackerBomId int `json:"tracker_bom_id" extensions:"x-order=0" description:"ไอดี Tracker BOM"`

	// ไอดีของล่องติดตาม
	TrackerId int `json:"tracker_id" extensions:"x-order=1" description:"ไอดีของล่องติดตาม"`

	// หมายเลขอีมี่ของล่องติดตามก
	TrackerCode string `json:"tracker_code" extensions:"x-order=2" description:"หมายเลขอีมี่ของกล่องติดตาม"`

	// ไอดีของ Gps
	GpsId int `json:"gps_id" extensions:"x-order=3" description:"ไอดี GPS"`

	// ชื่อยี่ห้อ Gps
	GpsBrandName string `json:"gps_brand_name" extensions:"x-order=4" description:"ชื่อแบรนด์ของ GPS"`

	// ชื่อรุ่น Gps
	GpsModelName string `json:"gps_model_name" extensions:"x-order=5" description:"ชื่อรุ่นของ GPS"`

	// ไอดีของ Gps Antenna
	GpsAntennaId int `json:"gps_antenna_id" extensions:"x-order=6" description:"ไอดีของ Gps Antenna"`

	// หมายเลข S/N ของ Gps Antenna
	GpsAntennaSerialNo string `json:"gps_antenna_serial_no" extensions:"x-order=7" description:"หมายเลข S/N ของ Gps Antenna"`

	// ไอดีของ GSM Antenna
	GsmAntennaId int `json:"gsm_antenna_id" extensions:"x-order=8" description:"ไอดีของ GSM Antenna"`

	// หมายเลข S/N ของ Gsm Antenna
	GsmAntennaSerialNo string `json:"gsm_antenna_serial_no" extensions:"x-order=9" description:"หมายเลข S/N ของ Gsm Antenna"`

	// ไอดีของเครื่องอ่านบัตร
	CardReaderId int `json:"card_reader_id" extensions:"x-order=10" description:"ไอดีของเครื่องอ่านบัตร"`

	// หมายเลข S/N ของ เครื่องอ่านบัตร
	CardReaderSerialNo string `json:"card_reader_serial_no" extensions:"x-order=11" description:"หมายเลข S/N ของเครื่องอ่านบัตร"`

	// ชื่อแบรนด์ของ เครื่องอ่านบัตร
	CardReaderBrandName string `json:"card_reader_brand_name" extensions:"x-order=12" description:"ชื่อแบรนด์ของ เครื่องอ่านบัตร"`

	// ชื่อรุ่นของ เครื่องอ่านบัตร
	CardReaderModelName string `json:"card_reader_model_name" extensions:"x-order=13" description:"ชื่อรุ่นของ เครื่องอ่านบัตร"`

	// ไอดีของซิมการ์ด
	SimId int `json:"sim_id" extensions:"x-order=14" description:"ไอดีของซิมการ์ด"`

	// หมายเลขของซิมการ์ด
	SimNo string `json:"sim_no" extensions:"x-order=15" description:"หมายเลขของซิมการ์ด"`

	// ไอดีของแบตเตอรี่
	BatteryId int `json:"battery_id" extensions:"x-order=16" description:"ไอดีของแบตเตอรี่"`

	// หมายเลข S/N ของแบตเตอรี่
	BatterySerialNo string `json:"battery_serial_no" extensions:"x-order=17" description:"หมายเลข S/N ของแบตเตอรี่"`
}

type EZWInsertTrackerBomRequest struct {

	// ไอดีของ Gps
	GpsId *int `json:"gps_id" example:"1" extensions:"x-order=0" description:"ไอดีของ Gps"`

	// ไอดีของ Gps Antenna
	GpsAntennaId *int `json:"gps_antenna_id" example:"2" extensions:"x-order=1" description:"ไอดีของ Gps Antenna"`

	// ไอดีของ Gsm Antenna
	GsmAntennaId *int `json:"gsm_antenna_id" example:"3" extensions:"x-order=2" description:"ไอดีของ Gsm Antenna"`

	// ไอดีของ Card Reader
	CardReaderId *int `json:"card_reader_id" example:"4" extensions:"x-order=3" description:"ไอดีของ Card Reader"`

	// ไอดีของซิมการ์ด
	SimId *int `json:"sim_id" example:"5" extensions:"x-order=4" description:"ไอดีของซิมการ์ด"`

	// ไอดีของแบตเตอรี่
	BatteryId *int `json:"battery_id" example:"6" extensions:"x-order=5" description:"ไอดีของแบตเตอรี่"`
}

type EZWInsertTrackerBomResponse struct {

	// หมายเหตุ
	Message string `json:"battery_id" example:"สำเร็จ" extensions:"x-order=0" description:"หมายเหตุ"`
}

type EZWUpdateTrackerBomRequest struct {
	TrackerBomId *int `json:"tracker_bom_id" binding:"required" extensions:"x-order=0"`
	GpsId        *int `json:"gps_id" binding:"required" extensions:"x-order=1"`
	GpsAntennaId *int `json:"gps_antenna_id" binding:"required" extensions:"x-order=2"`
	GsmAntennaId *int `json:"gsm_antenna_id" binding:"required" extensions:"x-order=3"`
	SimId        *int `json:"sim_id" binding:"required" extensions:"x-order=4"`
	BatteryId    *int `json:"battery_id" binding:"required" extensions:"x-order=5"`
}

// Response Payload
type EZWUpdateTrackerBomResponse struct {
	Message string `json:"message" extensions:"x-order=0"`
}

type EZWGetGps struct {
	Message string `json:"message" extensions:"x-order=0"`
}

type GPSResponse struct {
	GpsID    *int    `json:"gps_id" extensions:"x-order=0"`
	GpsIMEI  *string `json:"gps_imei" extensions:"x-order=1"`
	SerialNo *string `json:"serial_no" extensions:"x-order=2"`
}

// Battery Response
type BatteryResponse struct {
	BatteryID *int    `json:"battery_id" extensions:"x-order=0"`
	SerialNo  *string `json:"serial_no" extensions:"x-order=1"`
}

// CardReader Response
type CardReaderResponse struct {
	CardReaderID *int    `json:"card_reader_id" extensions:"x-order=0"`
	SerialNo     *string `json:"serial_no" extensions:"x-order=1"`
}

// GpsAntenna Response
type GpsAntennaResponse struct {
	GpsAntennaID *int    `json:"gps_antenna_id" extensions:"x-order=0"`
	SerialNo     *string `json:"serial_no" extensions:"x-order=1"`
}

// GsmAntenna Response
type GsmAntennaResponse struct {
	GsmAntennaID *int    `json:"gsm_antenna_id" extensions:"x-order=0"`
	SerialNo     *string `json:"serial_no" extensions:"x-order=1"`
}

// Sim Response
type SimResponse struct {
	SimID *int    `json:"sim_id" extensions:"x-order=0"`
	SimNo *string `json:"sim_no" extensions:"x-order=1"`
}

type EZWGetTrackerBomGeneralRequest struct {
	TrackerBomID int `json:"tracker_bom_id" binding:"required" example:"1"`
}

// Response
type EZWGetTrackerBomGeneralResponse struct {
	TrackerBomID       *int    `json:"tracker_bom_id" example:"123" extensions:"x-order=0"`
	TrackerBomCode     *string `json:"tracker_bom_code" example:"BOM-001" extensions:"x-order=1"`
	GpsID              *int    `json:"gps_id" extensions:"x-order=2"`
	GpsIMEI            *string `json:"gps_imei" extensions:"x-order=3"`
	GpsSerialNo        *string `json:"gps_serial_no" extensions:"x-order=4"`
	GpsAntennaID       *int    `json:"gps_antenna_id" extensions:"x-order=5"`
	GpsAntennaSerialNo *string `json:"gps_antenna_serial_no" extensions:"x-order=6"`
	GsmAntennaID       *int    `json:"gsm_antenna_id" extensions:"x-order=7"`
	GsmAntennaSerialNo *string `json:"gsm_antenna_serial_no" extensions:"x-order=8"`
	SimID              *int    `json:"sim_id" extensions:"x-order=9"`
	SimNo              *string `json:"sim_no" extensions:"x-order=10"`
	BatteryID          *int    `json:"battery_id" extensions:"x-order=11"`
	BatterySerialNo    *string `json:"battery_serial_no" extensions:"x-order=12"`
	CardReaderID       *int    `json:"card_reader_id" extensions:"x-order=13"`
	CardReaderSerialNo *string `json:"card_reader_serial_no" extensions:"x-order=14"`
	IsActive           *bool   `json:"is_active" extensions:"x-order=15"`
}
