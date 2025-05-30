package models

import "database/sql"

type EZWGetTrackerBom struct {

	// ไอดีของ Tracker Bom
	TrackerBomId int

	// ไอดีของกล่องติดตาม
	TrackerId sql.NullInt64

	// หมายเลขอีมี่ของกล่องติดตาม
	TrackerCode sql.NullString

	// ไอดีของ Gps
	GpsId sql.NullInt64

	// ชื่อยี่ห้อ Gps
	GpsBrandName sql.NullString

	// ชื่อรุ่น Gps
	GpsModelName sql.NullString

	// ไอดีของ Gps Antenna
	GpsAntennaId sql.NullInt64

	// หมายเลข S/N ของ Gps Antenna
	GpsAntennaSerialNo sql.NullString

	// ไอดีของ GSM Antenna
	GsmAntennaId sql.NullInt64

	// หมายเลข S/N ของ Gsm Antenna
	GsmAntennaSerialNo sql.NullString

	// ไอดีของเครื่องอ่านบัตร
	CardReaderId sql.NullInt64

	// หมายเลข S/N ของ เครื่องอ่านบัตร
	CardReaderSerialNo sql.NullString

	// ชื่อแบรนด์ของ เครื่องอ่านบัตร
	CardReaderBrandName sql.NullString

	// ชื่อรุ่นของ เครื่องอ่านบัตร
	CardReaderModelName sql.NullString
	// ไอดีของซิมการ์ด
	SimId sql.NullInt64

	// หมายเลขของซิมการ์ด
	SimNo sql.NullString

	// ไอดีของแบตเตอรี่
	BatteryId sql.NullInt64

	// หมายเลข S/N ของแบตเตอรี่
	BatterySerialNo sql.NullString
}

type EZWInsertTrackerBom struct {

	// ไอดีของ Gps
	GpsId *int

	// ไอดีของ Gps Antenna
	GpsAntennaId *int

	// ไอดีของ Gsm Antenna
	GsmAntennaId *int

	// ไอดีของ Card Reader
	CardReaderId *int

	// ไอดีของ Sim
	SimId *int

	// ไอดีของ Battery
	BatteryId *int
}

type GPSModel struct {
	GpsID    sql.NullInt64
	GpsIMEI  sql.NullString
	SerialNo sql.NullString
}

// Battery
type BatteryModel struct {
	BatteryID sql.NullInt64
	SerialNo  sql.NullString
}

// CardReader
type CardReaderModel struct {
	CardReaderID sql.NullInt64
	SerialNo     sql.NullString
}

// GpsAntenna
type GpsAntennaModel struct {
	GpsAntennaID sql.NullInt64
	SerialNo     sql.NullString
}

// GsmAntenna
type GsmAntennaModel struct {
	GsmAntennaID sql.NullInt64
	SerialNo     sql.NullString
}

// Sim
type SimModel struct {
	SimID sql.NullInt64
	SimNo sql.NullString
}

type EZWTrackerBomGeneralModel struct {
	TrackerBomID       sql.NullInt64
	TrackerBomCode     sql.NullString
	GpsID              sql.NullInt64
	GpsIMEI            sql.NullString
	GpsSerialNo        sql.NullString
	GpsAntennaID       sql.NullInt64
	GpsAntennaSerialNo sql.NullString
	GsmAntennaID       sql.NullInt64
	GsmAntennaSerialNo sql.NullString
	SimID              sql.NullInt64
	SimNo              sql.NullString
	BatteryID          sql.NullInt64
	BatterySerialNo    sql.NullString
	CardReaderId       sql.NullInt64
	CardReaderSerialNo sql.NullString
	IsActive           sql.NullBool
}
