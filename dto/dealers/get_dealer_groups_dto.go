package dealers

type GetDealerGroupsResponse struct {

	// รหัสกลุ่มดีลเลอร์
	// example: 1
	DealerGroupId int `json:"group_id" example:"1" extensions:"x-order=0"`

	// ชื่อกลุ่มดีลเลอร์
	// example: ดีลเลอร์ทั่วไป
	DealerGroupName string `json:"group_name" example:"ดีลเลอร์ทั่วไป" extensions:"x-order=1"`
}
