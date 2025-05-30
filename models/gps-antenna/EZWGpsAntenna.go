package models

import "database/sql"

type EZWGpsAntennaStatus struct {

	// สถานะไอดีของ Gps Antenna
	StatusId int

	// ชื่อสถานะของ Gps Antenna
	StatusName string
}

type EZWSearchGpsAntenna struct {

	// ไอดีของ Gps Antenna
	GpsAntennaId int

	// หมายเลขซีเรียลนัมเบอร์ Gps Antenna
	SerialNo string

	// ไอดีสถานะ Gps Antenna
	StatusId int

	// รหัสสถานะ Gps Antenna
	StatusCode string

	// ชื่อสถานะ Gps Antenna
	StatusName string

	// หมายเหตุ
	Remark string
}

type GpsAntennaGeneralModel struct {
	GpsAntennaID sql.NullInt64  // ไอดีของ Gps Antenna
	SerialNo     sql.NullString // หมายเลขซีเรียลนัมเบอร์ Gps Antenna
	StatusID     sql.NullInt64  // ไอดีสถานะ Gps Antenna
	StatusName   sql.NullString // ชื่อสถานะ Gps Antenna
	Remark       sql.NullString // หมายเหตุ
}
