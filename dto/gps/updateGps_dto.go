package gps

type UpdateGPSRequest struct {
	GpsID        int    `json:"gps_id" example:"27" extensions:"x-order=0"`
	GpsIMEI      string `json:"gps_imei" example:"864714063600838" extensions:"x-order=1"`
	SerialNumber string `json:"serial_no" example:"14063600838" extensions:"x-order=2"`
	BrandID      int    `json:"brand_id" example:"136001" extensions:"x-order=3"`
	ModelID      int    `json:"model_id" example:"137001" extensions:"x-order=4"`
	StatusID     int    `json:"status_id" example:"119002" extensions:"x-order=5"`
	Remark       string `json:"remark" example:"ทดสอบกับAPI" extensions:"x-order=6"`
}

type UpdateGPSResponse struct {
	Message string `json:"message" example:"Update GPS data successfully"`
}
