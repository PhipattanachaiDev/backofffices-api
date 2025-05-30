package models

import "encoding/json"

type EZWAuthenticationModel struct {
	UserId   int             // รหัสผู้ใช้งาน
	UserName string          // ชื่อผู้ใช้งาน
	RoleId   int             // รหัสบทบาท
	RoleName string          // ชื่อบทบาท
	Access   json.RawMessage // สิทธิ์การเข้าถึง
}
