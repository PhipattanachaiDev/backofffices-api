package models

import "database/sql"

type EZWCardReaderStatusModel struct {
	// ไอดีของสถานะเครื่องอ่านบัตร
	StatusId int

	// ชื่อของสถานะเครื่องอ่านบัตร
	StatusName string
}

type EZWSearchCardReader struct {
	// ไอดีของเครื่องอ่านบัตร
	CardReaderId int

	// หมายเลขเครื่องอ่านบัตร
	SerialNo string

	// ไอดีของแบรนด์เครื่องอ่านบัตร
	BrandId int

	// ชื่อแบรนด์ของเครื่องอ่านบัตร
	BrandName string

	// ไอดีของรุ่นเครื่องอ่านบัตร
	ModelId sql.NullInt32

	// ชื่อรุ่นของเครื่องอ่านบัตร
	ModelName sql.NullString

	// ไอดีสถานะของเครื่องอ่านบัตร
	StatusId int

	// ชื่อสถานะของเครื่องอ่านบัตร
	StatusName string

	// หมายเหตุเพิ่มเติม
	Remark sql.NullString
}

type EZWCardReaderBrand struct {
	// ไอดีของแบรนด์เครื่องอ่านบัตร
	BrandId int

	// ชื่อแบรนด์เครื่องอ่านบัตร
	BrandName string
}

type EZWCardReaderModel struct {
	// ไอดีของรุ่นเครื่องอ่านบัตร
	ModelId int

	// ชื่อรุ่นเครื่องอ่านบัตร
	ModelName string
}

type CardReaderGeneralModel struct {
	CardReaderID sql.NullInt64  // ไอดีของเครื่องอ่านบัตร
	BrandID      sql.NullInt64  // ไอดีของแบรนด์เครื่องอ่านบัตร
	BrandName    sql.NullString // ชื่อแบรนด์เครื่องอ่านบัตร
	ModelID      sql.NullInt64  // ไอดีของรุ่นเครื่องอ่านบัตร
	ModelName    sql.NullString // ชื่อรุ่นเครื่องอ่านบัตร
	SerialNo     sql.NullString // หมายเลขเครื่องอ่านบัตร
	StatusID     sql.NullInt64  // ไอดีสถานะของเครื่องอ่านบัตร
	StatusName   sql.NullString // ชื่อสถานะของเครื่องอ่านบัตร
	Remark       sql.NullString // หมายเหตุเพิ่มเติม
}
