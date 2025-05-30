package dealers

type GetDealerStatusResponse struct {
	// รหัสสถานะตัวแทนจำหน่าย
	// example: 1
	StatusId int `json:"status_id" example:"1" extensions:"x-order=0"`

	// ชื่อสถานะตัวแทนจำหน่าย
	// example: ใช้งานได้
	StatusName string `json:"status_name" example:"ใช้งานได้" extensions:"x-order=1"`
}
