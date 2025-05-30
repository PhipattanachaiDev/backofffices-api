package dealers

type GetDealersByConditionsRequest struct {

	// รหัสตัวแทนจำหน่าย
	DealerId string `json:"dealer_id"`

	// ชื่อตัวแทนจำหน่าย
	DealerName string `json:"dealer_name"`

	// กลุ่มตัวแทนจำหน่าย
	DealerGroup int64 `json:"dealer_group"`

	// รายละเอียดตัวแทนจำหน่าย
	DealerDetail string `json:"dealer_detail"`

	// สถานะตัวแทนจำหน่าย
	DealerStatus int64 `json:"dealer_status"`
}

type GetDealersByConditionsResponse struct {

	// รหัสตัวแทนจำหน่าย
	DealerId int `json:"dealer_id"`

	// ชื่อตัวแทนจำหน่าย
	DealerName string `json:"dealer_name"`

	// รหัสกลุ่มตัวแทนจำหน่าย
	DealerGroupId int64 `json:"dealer_group_id"`

	// กลุ่มตัวแทนจำหน่าย
	DealerGroup string `json:"dealer_group"`

	// รายละเอียดตัวแทนจำหน่าย
	DealerDetail string `json:"dealer_detail"`

	// สถานะตัวแทนจำหน่าย
	DealerStatus string `json:"dealer_status"`
}
