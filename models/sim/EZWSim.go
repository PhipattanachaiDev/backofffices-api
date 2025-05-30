package models

import "database/sql"

type EZWGetSimStatus struct {
	StatusId   int    // ประเภทของไอดีซิมการ์ด
	StatusName string // ชื่อประเภทสถานะของซิมการ์ด
}
type EZWGetSimOperator struct {
	OperatorId   int    // ประเภทของไอดีซิมการ์ด
	OperatorName string // ชื่อประเภทสถานะของซิมการ์ด
}

type EZWSearchSim struct {
	SimId        int    // ไอดีของซิมการ์ด
	SimNo        string // หมายเลขซิมการ์ด
	StatusId     int    // ไอดีสถานะซิมการ์ด
	StatusCode   string // รหัสสถานะซิมการ์ด
	StatusName   string // ชื่อสถานะซิมการ์ด
	OperatorId   int    // ไอดีผู้ให้บริการ
	OperatorCode string //รหัสผู้ให้บริการ
	OperatorName string // ชื่อผู้ให้บริการ
	Remark       string // หมายเหตุ

}

type EZWInsertSim struct {
	Message string // รายละเอียดตอบกลับ
}
type EZWUPdateSim struct {
	Message string // รายละเอียดตอบกลับ
}

type SimGeneralModel struct {
	SimID        sql.NullInt64  // ไอดีของซิมการ์ด
	SimNo        sql.NullString // หมายเลขซิมการ์ด
	OperatorID   sql.NullInt64  // ไอดีผู้ให้บริการ
	OperatorName sql.NullString // ชื่อผู้ให้บริการ
	StatusID     sql.NullInt64  // ไอดีสถานะซิมการ์ด
	StatusName   sql.NullString // ชื่อสถานะซิมการ์ด
	Remark       sql.NullString // หมายเหตุ
}
