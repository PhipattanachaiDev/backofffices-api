package gps

type GetDistrictRequest struct {
	ProvinceID int `json:"province_id" binding:"required"`
}

type DistrictResponse struct {
	DistrictID   uint   `json:"district_id" example:"1"`
	DistrictName string `json:"district_name" example:"เขต พระนคร"`
}
