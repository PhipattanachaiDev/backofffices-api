package gps

type GPSRequest struct {
	// ชื่อสินค้า ของ GPS
	GpsIMEI string `json:"gps_imei" example:"" extensions:"x-order=0"`
	// หมายเลขรหัสสินค้า ของ GPS
	SerialNumber string `json:"serial_no" example:"" extensions:"x-order=1"`
	// ID ยี่ห้อ GPS
	BrandID int `json:"brand_id" example:"0" extensions:"x-order=2"`
	// ID รุ่น GPS
	ModelID int `json:"model_id" example:"0" extensions:"x-order=3"`
	// ID สถานะ GPS
	StatusID int `json:"status_id" example:"0" extensions:"x-order=4"`
	// รายละเอียด
	Remark string `json:"remark" example:"" extensions:"x-order=5"`
}

type GPSResponse struct {
	GpsID        int    `json:"gps_id" example:"1" extensions:"x-order=0"`
	GpsIMEI      string `json:"gps_imei" example:"863141055492266" extensions:"x-order=1"`
	SerialNumber string `json:"serial_no" example:"59624920078" extensions:"x-order=2"`
	StatusID     int    `json:"status_id" example:"119002" extensions:"x-order=3"`
	StatusName   string `json:"status_name" example:"Active" extensions:"x-order=4"`
	Remark       string `json:"remark" example:"สำหรับทดสอบ" extensions:"x-order=5"`
	BrandName    string `json:"brand_name" example:"Meitrack" extensions:"x-order=6"`
	BrandCode    string `json:"brand_code" example:"GPS-02" extensions:"x-order=7"`
	ModelName    string `json:"model_name" example:"T399L" extensions:"x-order=8"`
	ModelCode    string `json:"model_code" example:"GPS-02-07" extensions:"x-order=9"`
}
