package models

import "database/sql"

type EZWBatteryStatus struct {

	// สถานะไอดีของ Battery
	StatusId int

	// ชื่อสถานะของ Battery
	StatusName string
}

type EZWSearchBattery struct {

	// ไอดีของ Battery
	BatteryId int

	// หมายเลขซีเรียลนัมเบอร์ Battery
	SerialNo string

	// ไอดีสถานะ Battery
	StatusId int

	// ชื่อสถานะแบตเตอรี่
	StatusName string

	// หมายเหตุ
	Remark string
}

type BatteryGeneralModel struct {
	BatteryID  sql.NullInt64  // ไอดีแบตเตอรี่
	SerialNo   sql.NullString // หมายเลขซีเรียลนัมเบอร์แบตเตอรี่
	StatusID   sql.NullInt64  // ไอดีสถานะแบตเตอรี่
	StatusName sql.NullString // ชื่อสถานะแบตเตอรี่
	Remark     sql.NullString // หมายเหตุ
}
