package tracker

type EZWSearchTrackerRequest struct {

	// รหัสกล่องติดตาม
	TrackerCode string `json:"tracker_code" example:"2024" extensions:"x-order=0" description:"รหัสกล่องติดตาม"`

	// ไอดียี่ห้อกล่องติดตาม
	BrandId int `json:"brand_id" example:"146001" extensions:"x-order=1" description:"ไอดียี่ห้อกล่องติดตาม"`

	// ไอดีรุ่นกล่องติดตาม
	ModelId int `json:"model_id" example:"147001" extensions:"x-order=2" description:"ไอดีรุ่นกล่องติดตาม"`

	// ไอดีสถานะกล่องติดตาม
	StatusId int `json:"status_id" example:"128001" extensions:"x-order=3" description:"ไอดีสถานะกล่องติดตาม"`

	// Remark (หมายเหตุ)
	Remark string `json:"remark" example:"eZView" extensions:"x-order=4" description:"หมายเหตุ"`
}

type EZWSearchTrackerModelResponse struct {

	// ไอดีรายการวัสดุกล่องติดตาม
	TrackerBomId int `json:"tracker_bom_id" example:"2024" extensions:"x-order=0" description:"ไอดีรายการวัสดุกล่องติดตาม"`

	// ไอดีกล่องติดตาม
	TrackerId int `json:"tracker_id" example:"2024" extensions:"x-order=1" description:"ไอดีรายการวัสดุกล่องติดตาม"`

	// รหัสกล่องติดตาม
	TrackerCode string `json:"tracker_code" example:"2024" extensions:"x-order=2" description:"ไอดีรายการวัสดุกล่องติดตาม"`

	// หมายเลขซีเรียลนัมเบอร์กล่องติดตาม
	SerialNo string `json:"serial_no" example:"2024" extensions:"x-order=3" description:"ไอดีรายการวัสดุกล่องติดตาม"`

	// ไอดีรุ่นกล่องติดตาม
	TrackerModelId int `json:"tracker_model_id" example:"2024" extensions:"x-order=4" description:"ไอดีรายการวัสดุกล่องติดตาม"`

	// ชื่อรุ่นกล่องติดตาม
	TrackerModelName string `json:"tracker_model_name" example:"2024" extensions:"x-order=5" description:"ไอดีรายการวัสดุกล่องติดตาม"`

	// ไอดียี่ห้อกล่องติดตาม
	TrackerBrandId int `json:"tracker_brand_id" example:"2024" extensions:"x-order=6" description:"ไอดีรายการวัสดุกล่องติดตาม"`

	// ชื่อยี่ห้อกล่องติดตาม
	TrackerBrandName string `json:"tracker_brand_name" example:"2024" extensions:"x-order=7" description:"ไอดีรายการวัสดุกล่องติดตาม"`

	// ไอดีสถานะกล่องติดตาม
	StatusId int `json:"status_id" example:"2024" extensions:"x-order=8" description:"ไอดีรายการวัสดุกล่องติดตาม"`

	// ชื่อสถานะกล่องติดตาม
	StatusName string `json:"status_name" example:"2024" extensions:"x-order=9" description:"ไอดีรายการวัสดุกล่องติดตาม"`

	// หมายเหตุ
	Remark string `json:"remark" example:"2024" extensions:"x-order=10" description:"ไอดีรายการวัสดุกล่องติดตาม"`
}

type EZWTrackerBrandResponse struct {

	// ไอดียี่ห้อกล่องติดตาม
	BrandId int `json:"brand_id" extensions:"x-order=0" description:"ไอดียี่ห้อกล่องติดตาม"`

	// ชื่อยี่ห้อกล่องติดตาม
	BrandName string `json:"brand_name" extensions:"x-order=1" description:"ชื่อยี่ห้อกล่องติดตาม"`
}

type EZWTrackerStatusResponse struct {

	// ไอดีสถานะกล่องติดตาม
	StatusId int `json:"status_id" extensions:"x-order=0" description:"ไอดีสถานะกล่องติดตาม"`

	// ชื่อสถานะกล่องติดตาม
	StatusName string `json:"status_name" extensions:"x-order=1" description:"ชื่อสถานะกล่องติดตาม"`
}

type EZWGetTrackerModelRequest struct {

	// ไอดีของยี่ห้อกล่องติดตาม
	ParentId int `json:"parent_id" example:"146001" extensions:"x-order=0" description:"ไอดีของยี่ห้อกล่องติดตาม"`
}

type EZWTrackerModelResponse struct {

	// ไอดีรุ่นกล่องติดตาม
	ModelId int `json:"model_id" extensions:"x-order=0" description:"ไอดีรุ่นกล่องติดตาม"`

	// ชื่อรุ่นกล่องติดตาม
	ModelName string `json:"model_name" extensions:"x-order=1" description:"ชื่อรุ่นกล่องติดตาม"`
}

type EZWUpdateTrackerRequest struct {

	// ไอดีกล่องติดตาม
	TrackerId int `json:"tracker_id"  example:"1" extensions:"x-order=0" description:"ไอดีกล่องติดตาม"`

	// รหัสกล่องติดตาม
	TrackerCode string `json:"tracker_code"  example:"TRK001" extensions:"x-order=1" description:"รหัสกล่องติดตาม"`

	// หมายเลขซีเรียลนัมเบอร์กล่องติดตาม
	SerialNo string `json:"serial_no"  example:"SN123456" extensions:"x-order=2" description:"หมายเลขซีเรียลนัมเบอร์กล่องติดตาม"`

	// ไอดีกลุ่มอุปกรณ์ติดตาม
	TrackerBomId int `json:"tracker_bom_id"  example:"2" extensions:"x-order=3" description:"ไอดีกลุ่มอุปกรณ์ติดตาม"`

	// ไอดียี่ห้อกล่องติดตาม
	BrandId int `json:"brand_id"  example:"146001" extensions:"x-order=4" description:"ไอดียี่ห้อกล่องติดตาม"`

	// ไอดีรุ่นกล่องติดตาม
	ModelId int `json:"model_id"  example:"147001" extensions:"x-order=5" description:"ไอดีรุ่นกล่องติดตาม"`

	// ไอดีสถานะกล่องติดตาม
	StatusId int `json:"status_id"  example:"128001" extensions:"x-order=6" description:"ไอดีสถานะกล่องติดตาม"`

	// หมายเหตุ
	Remark string `json:"remark" example:"Updated tracker details" extensions:"x-order=7" description:"หมายเหตุ"`
}

type EZWUpdateTrackerResponse struct {

	// หมายเหตุ
	Message string `json:"message" example:"Update successful" description:"หมายเหตุ"`
}

type EZWInsertTrackerRequest struct {
	TrackerCode  string `json:"tracker_code" binding:"required" extensions:"x-order=0" example:"202309010001"`
	SerialNo     string `json:"serial_no" binding:"required"  extensions:"x-order=1" example:"SN202309010001"`
	TrackerBomId int    `json:"tracker_bom_id" binding:"required"  extensions:"x-order=2" example:"1"`
	BrandId      int    `json:"brand_id" binding:"required"  extensions:"x-order=3" example:"146001"`
	ModelId      int    `json:"model_id" binding:"required"  extensions:"x-order=4" example:"147002"`
	StatusId     int    `json:"status_id" binding:"required"  extensions:"x-order=5" example:"128001"`
	Remark       string `json:"remark"  extensions:"x-order=6" example:"New tracker" `
}

type EZWInsertTrackerResponse struct {
	Message string `json:"message" example:"Insert successful: tracker_id = 123"`
}

type EZWRequestGetTrackerGeneral struct {
	TrackerId int `json:"tracker_id"  binding:"required" example:"1"`
}

type EZWResponseGetTrackerGeneral struct {
	TrackerBomId     *int    `json:"tracker_bom_id,omitempty"`
	TrackerCode      *string `json:"tracker_code,omitempty"`
	SerialNo         *string `json:"serial_no,omitempty"`
	TrackerBrandId   *int    `json:"tracker_brand_id,omitempty"`
	TrackerBrandName *string `json:"tracker_brand_name,omitempty"`
	TrackerModelId   *int    `json:"tracker_model_id,omitempty"`
	TrackerModelName *string `json:"tracker_model_name,omitempty"`
	StatusId         *int    `json:"status_id,omitempty"`
	StatusName       *string `json:"status_name,omitempty"`
	Remark           *string `json:"remark,omitempty"`
}
