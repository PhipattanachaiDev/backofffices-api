package customers

type GetCustomerGroupsResponse struct {

	// รหัสกลุ่มลูกค้า
	// example: 1
	CustomerGroupId int `json:"group_id" example:"1" extensions:"x-order=0"`

	// ชื่อกลุ่มลูกค้า
	// example: ลูกค้าทั่วไป
	CustomerGroupName string `json:"group_name" example:"ลูกค้าทั่วไป" extensions:"x-order=1"`
}
