package models

import "database/sql"

type EZWSearchTrackerModel struct {

	// ไอดีรายการวัสดุกล่องติดตาม
	TrackerBomId int

	// ไอดีกล่องติดตาม
	TrackerId int

	// รหัสกล่องติดตาม
	TrackerCode sql.NullString

	// หมายเลขซีเรียลนัมเบอร์กล่องติดตาม
	SerialNo sql.NullString

	// ไอดีรุ่นกล่องติดตาม
	TrackerModelId sql.NullInt64

	// ชื่อรุ่นกล่องติดตาม
	TrackerModelName sql.NullString

	// ไอดียี่ห้อกล่องติดตาม
	TrackerBrandId sql.NullInt64

	// ชื่อยี่ห้อกล่องติดตาม
	TrackerBrandName sql.NullString

	// ไอดีสถานะกล่องติดตาม
	StatusId sql.NullInt64

	// ชื่อสถานะกล่องติดตาม
	StatusName sql.NullString

	// หมายเหตุ
	Remark sql.NullString
}

type EZWTrackerBrand struct {
	BrandId   int
	BrandName sql.NullString
}

type EZWTrackerStatus struct {
	StatusId   int
	StatusName sql.NullString
}

type EZWTrackerModel struct {
	ModelId   int
	ModelName sql.NullString
}

type EZWTrackerUpdateModel struct {
	Message string
}

type EZWInsertTrackerModel struct {
	TrackerCode  string
	SerialNo     string
	TrackerBomId int
	BrandId      int
	ModelId      int
	StatusId     int
	Remark       sql.NullString
	UserId       int
}

type EZWGetTrackerGeneral struct {
	TrackerBomId     sql.NullInt64
	TrackerCode      sql.NullString
	SerialNo         sql.NullString
	TrackerBrandId   sql.NullInt64
	TrackerBrandName sql.NullString
	TrackerModelId   sql.NullInt64
	TrackerModelName sql.NullString
	StatusId         sql.NullInt64
	StatusName       sql.NullString
	Remark           sql.NullString
}
