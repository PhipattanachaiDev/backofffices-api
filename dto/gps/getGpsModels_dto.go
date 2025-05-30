package gps

type GPSModelRequest struct {
	TypeID int `json:"type_id" binding:"required"`
}

type GPSModelResponse struct {
	TypeID   int    `json:"model_id" example:"137001"`
	TypeName string `json:"model_name" example:"T399L"`
}
