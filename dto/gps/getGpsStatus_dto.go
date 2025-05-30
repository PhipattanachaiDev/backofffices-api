package gps

type GPSStatusResponse struct {
	TypeID   int    `json:"status_id" example:"119001"`
	TypeName string `json:"status_name" example:"Active"`
}
