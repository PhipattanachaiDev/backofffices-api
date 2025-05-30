package models

import "database/sql"

type EZWGetGpsGeneralModel struct {
	GpsImei    sql.NullString
	SerialNo   sql.NullString
	StatusId   sql.NullInt64
	StatusName sql.NullString
	BrandId    sql.NullInt64
	BrandName  sql.NullString
	ModelId    sql.NullInt64
	ModelName  sql.NullString
	Remark     sql.NullString
}
