package models

type EZWSearchDealers struct {
	DealerId      int    // รหัสตัวแทนจำหน่าย
	DealerName    string // ชื่อตัวแทนจำหน่าย
	DealerGroupId int64  // รหัสกลุ่มตัวแทนจำหน่าย
	DealerGroup   string // กลุ่มตัวแทนจำหน่าย
	DealerDetail  string // รายละเอียดตัวแทนจำหน่าย
	DealerStatus  string // สถานะตัวแทนจำหน่าย
}
