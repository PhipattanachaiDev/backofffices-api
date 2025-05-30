package customers

type GetCustomersByConditionsRequest struct {

	// รหัสลูกค้า
	CustomerId string `json:"customer_id"`

	// ชื่อลูกค้า
	CustomerName string `json:"customer_name"`

	// กลุ่มลูกค้า
	CustomerGroup int64 `json:"customer_group"`

	// รายละเอียดลูกค้า
	CustomerDetail string `json:"customer_detail"`

	// สถานะลูกค้า
	CustomerStatus int64 `json:"customer_status"`
}

type GetCustomersByConditionsResponse struct {

	// รหัสลูกค้า
	CustomerId int `json:"customer_id"`

	// ชื่อลูกค้า
	CustomerName string `json:"customer_name"`

	// รหัสกลุ่มลูกค้า
	CustomerGroupId int64 `json:"customer_group_id"`

	// กลุ่มลูกค้า
	CustomerGroup string `json:"customer_group"`

	// รายละเอียดลูกค้า
	CustomerDetail string `json:"customer_detail"`

	// สถานะลูกค้า
	CustomerStatus string `json:"customer_status"`
}
