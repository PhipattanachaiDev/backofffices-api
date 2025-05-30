package models

type EZWSearchCustomers struct {
	CustomerId      int    // รหัสลูกค้า
	CustomerName    string // ชื่อลูกค้า
	CustomerGroupId int64  // รหัสกลุ่มลูกค้า
	CustomerGroup   string // กลุ่มลูกค้า
	CustomerDetail  string // รายละเอียดลูกค้า
	CustomerStatus  string // สถานะลูกค้า
}
