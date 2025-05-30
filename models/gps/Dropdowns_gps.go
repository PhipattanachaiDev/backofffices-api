package models

type GPSStatus struct {
	TypeID   int
	TypeName string
}

type GPSModel struct {
	TypeID   int
	TypeName string
}

type GPSBrand struct {
	TypeID   int
	TypeName string
}

type SearchGPS struct {
	GpsID        int
	GpsIMEI      string
	SerialNumber string
	StatusID     int
	StatusName   string
	Remark       string
	BrandName    string
	BrandCode    string
	ModelName    string
	ModelCode    string
}
