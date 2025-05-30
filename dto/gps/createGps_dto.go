package gps

type CreateGPSRequest struct {
	GpsIMEI      string `json:"gps_imei" example:"00000001" extensions:"x-order=0"`
	SerialNumber string `json:"serial_no" example:"10000001" extensions:"x-order=1"`
	BrandID      int    `json:"brand_id" example:"136001" extensions:"x-order=2"`
	ModelID      int    `json:"model_id" example:"137001" extensions:"x-order=3"`
	StatusID     int    `json:"status_id" example:"119002" extensions:"x-order=4"`
	Remark       string `json:"remark" example:"ทดสอบ" extensions:"x-order=5"`
}

type CreateGPSResponse struct {
	Message string `json:"message" example:"Create GPS data successfully"`
}
