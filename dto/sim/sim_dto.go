package sim

type EZWSimStatusModelResponse struct {

	// ประเภทไอดีของสถานะซิมการ์ด
	StatusId int `json:"status_id" example:"126002" extensions:"x-order=0" description:"ประเภทไอดีของสถานะซิมการ์ด"`

	// ชื่อประเภทสถานะของซิมการ์ด
	StatusName string `json:"status_name" example:"Inactive" extensions:"x-order=1" description:"ชื่อประเภทสถานะของซิมการ์ด"`
}
type EZWSimOperatorModelResponse struct {

	// ประเภทไอดีของผู้ให้บริการซิมการ์ด
	OperatorId int `json:"operator_id" example:"125001" extensions:"x-order=0" description:"ประเภทไอดีของผู้ให้บริการซิมการ์ด"`

	// ชื่อประเภทผู้ให้บริการซิมการ์ด
	OperatorName string `json:"operator_name" example:"AIS" extensions:"x-order=1" description:"ชื่อประเภทผู้ให้บริการซิมการ์ด"`
}

type EZWSearchSimResponse struct {

	// ไอดีของซิมการ์ด
	SimId int `json:"sim_id" example:"1" extensions:"x-order=0" description:"ไอดีของซิมการ์ด"`

	// หมายเลขซิมการ์ด
	SimNo string `json:"sim_no" example:"66672455949" extensions:"x-order=1" description:"หมายเลขซิมการ์ด"`

	// ไอดีสถานะซิมการ์ด
	StatusId int `json:"status_id" example:"126002" extensions:"x-order=2" description:"ไอดีสถานะซิมการ์ด"`

	// รหัสสถานะซิมการ์ด
	StatusCode string `json:"status_code" example:"Active" extensions:"x-order=3" description:"รหัสสถานะซิมการ์ด"`

	// ชื่อสถานะซิมการ์ด
	StatusName string `json:"status_name" example:"Active" extensions:"x-order=4" description:"ชื่อสถานะซิมการ์ด"`

	// ไอดีผู้ให้บริการ
	OperatorId int `json:"operator_id" example:"125001" extensions:"x-order=5" description:"ไอดีผู้ให้บริการ"`

	// รหัสผู้ให้บริการ
	OperatorCode string `json:"operator_code" example:"TrueMove" extensions:"x-order=6" description:"รหัสผู้ให้บริการ"`

	// ชื่อผู้ให้บริการ
	OperatorName string `json:"operator_name" example:"TrueMove H" extensions:"x-order=7" description:"ชื่อผู้ให้บริการ"`

	// หมายเหตุ
	Remark string `json:"remark" example:"eZView PLUS" extensions:"x-order=8" description:"หมายเหตุ"`
}

type EZWSearchSimRequest struct {

	// หมายเลขซิมการ์ด
	// required: true
	SimNo string `json:"sim_no" example:"66672455949" extensions:"x-order=0"  description:"หมายเลขซิมการ์ด"`

	// ไอดีผู้ให้บริการซิม
	// required: true
	OperatorId int `json:"operator_id" example:"99" extensions:"x-order=1" description:"ไอดีผู้ให้บริการซิม"`

	// ไอดีสถานะซิม
	// required: true
	StatusId int `json:"status_id" example:"99" extensions:"x-order=2" description:"ไอดีสถานะซิม"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"eZView PLUS" extensions:"x-order=3" description:"หมายเหตุ"`
}

type EZWInsertSimRequest struct {

	// เลขที่ซิมการ์ด
	// required: true
	SimNo string `json:"sim_no" example:"66672455949" extensions:"x-order=0"  description:"เลขที่ซิมการ์ด"`

	// ไอดีผู้ให้บริการซิม
	// required: true
	OperatorId int `json:"operator_id" example:"125002" extensions:"x-order=1" description:"ไอดีผู้ให้บริการซิม"`

	// ไอดีสถานะซิม
	// required: true
	StatusId int `json:"status_id" example:"126002" extensions:"x-order=2" description:"ไอดีสถานะซิม"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"สำหรับทดสอบใน API" extensions:"x-order=3" description:"หมายเหตุ"`
}

type EZWInsertSimResponse struct {

	// รายละเอียดตอบกลับ
	// required: true
	Message string `json:"remark" example:"eZView PLUS" extensions:"x-order=0" description:"รายละเอียดตอบกลับ"`
}

type EZWUpdateSimRequest struct {

	// ไอดีซิมการ์ด
	// required: true
	SimId int `json:"sim_id" example:"1" extensions:"x-order=0"  description:"ไอดีซิมการ์ด"`

	// หมายเลขซิมการ์ด
	// required: true
	SimNo string `json:"sim_no" example:"66672455949" extensions:"x-order=1"  description:"หมายเลขซิมการ์ด"`

	// ไอดีผู้ให้บริการซิม
	// required: true
	OperatorId int `json:"operator_id" example:"125001" extensions:"x-order=2" description:"ไอดีผู้ให้บริการซิม"`

	// ไอดีสถานะซิม
	// required: true
	StatusId int `json:"status_id" example:"126002" extensions:"x-order=3" description:"ไอดีสถานะซิม"`

	// หมายเหตุ
	// required: true
	Remark string `json:"remark" example:"สำหรับทดสอบใน API" extensions:"x-order=4" description:"หมายเหตุ"`
}
type EZWUpdateSimResponse struct {

	// รายละเอียดตอบกลับ
	// required: true
	Message string `json:"remark" example:"eZView PLUS" extensions:"x-order=0" description:"รายละเอียดตอบกลับ"`
}

type EZWGetSimGeneralRequest struct {
	// ไอดีของซิมการ์ด
	SimId *int `json:"sim_id" extensions:"x-order=0" description:"ไอดีของซิมการ์ด"`
}

type EZWGetSimGeneralResponse struct {

	// ไอดีของซิมการ์ด
	SimID *int `json:"sim_id" extensions:"x-order=0" description:"ไอดีของซิมการ์ด"`

	// หมายเลขซิมการ์ด
	SimNo *string `json:"sim_no" extensions:"x-order=1" description:"หมายเลขซิมการ์ด"`

	// ไอดีผู้ให้บริการซิมการ์ด
	OperatorID *int `json:"operator_id" extensions:"x-order=2" description:"ไอดีผู้ให้บริการซิมการ์ด"`

	// รหัสผู้ให้บริการซิมการ์ด
	OperatorName *string `json:"operator_name" extensions:"x-order=3" description:"รหัสผู้ให้บริการซิมการ์ด"`

	// ไอดีสถานะซิมการ์ด
	StatusID *int `json:"status_id" extensions:"x-order=4" description:"ไอดีสถานะซิมการ์ด"`

	// รหัสสถานะซิมการ์ด
	StatusName *string `json:"status_name" extensions:"x-order=5" description:"รหัสสถานะซิมการ์ด"`

	// หมายเลขสถานะซิมการ์ด
	Remark *string `json:"remark" extensions:"x-order=6" description:"หมายเลขสถานะซิมการ์ด"`
}
