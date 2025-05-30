package gps

type EZWGetGpsGeneralByGpsIdRequest struct {
	GpsId int `json:"gps_id" form:"gps_id" binding:"required" example:"1"`
}

type EZWGetGpsGeneralByGpsIdResponse struct {
	GpsImei    *string `json:"gps_imei,omitempty"`
	SerialNo   *string `json:"serial_no,omitempty"`
	StatusId   *int    `json:"status_id,omitempty"`
	StatusName *string `json:"status_name,omitempty"`
	BrandId    *int    `json:"brand_id,omitempty"`
	BrandName  *string `json:"brand_name,omitempty"`
	ModelId    *int    `json:"model_id,omitempty"`
	ModelName  *string `json:"model_name,omitempty"`
	Remark     *string `json:"remark,omitempty"`
}
