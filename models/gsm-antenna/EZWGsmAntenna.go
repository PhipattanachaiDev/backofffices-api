package models

import "database/sql"

type EZWGetGsmAntennaStatus struct {
	StatusId   int    // ประเภทไอดีของ Gsm Antenna
	StatusName string // ชื่อประเภทสถานะของ Gsm Antenna
}

type EZWSearchGsmAntenna struct {
	// ไอดีของ Gsm Antenna
	GsmAntennaId int

	// หมายเลขซีเรียลนัมเบอร์ Gsm Antenna
	SerialNo string

	// ไอดีสถานะ Gsm Antenna
	StatusId int

	// รหัสสถานะ Gsm Antenna
	StatusCode string

	// ชื่อสถานะ Gsm Antenna
	StatusName string

	// หมายเหตุ
	Remark string
}

type GsmAntennaGeneralModel struct {
	GsmAntennaID sql.NullInt64  // ไอดีของ Gsm Antenna
	SerialNo     sql.NullString // หมายเลขซีเรียลนัมเบอร์ Gsm Antenna
	StatusID     sql.NullInt64  // ไอดีสถานะ Gsm Antenna
	StatusName   sql.NullString // ชื่อสถานะ Gsm Antenna
	Remark       sql.NullString // หมายเหตุ
}
