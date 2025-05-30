package gps

// GetSubDistrictRequest struct
type GetSubDistrictRequest struct {
	DistrictID int `json:"district_id" binding:"required"`
}

// SubDistrictResponse struct
type SubDistrictResponse struct {
	SubDistrictID   uint   `json:"sub_district_id" example:"1"`
	SubDistrictName string `json:"sub_district_name" example:"แขวง วัดอรุณ"`
	ZipCode         string `json:"zip_code" example:"10200"`
}
