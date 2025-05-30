package responses

import "time"

// Generic Success Response for Swagger
type SuccessResponseSwagger[T any] struct {
	// สถานะของ response
	Success bool `json:"success" example:"true" extensions:"x-order=0"`

	// รหัสสถานะของ response
	Status int `json:"status" example:"200" extensions:"x-order=1"`

	// เวลาที่ response ถูกสร้างขึ้น
	Timestamp time.Time `json:"timestamp" example:"2025-01-01T00:00:00.000000000Z" extensions:"x-order=2"`

	// ข้อความสถานะของ response
	Message string `json:"message" example:"OK" extensions:"x-order=3"`

	// ข้อมูลที่ต้องการส่งกลับ
	Data T `json:"data,omitempty" extensions:"x-order=4"`
}

// Create Response for Swagger (201)
type CreateResponseSwagger struct {
	// สถานะของ response
	Success bool `json:"success" example:"true" extensions:"x-order=0"`

	// รหัสสถานะของ response
	Status int `json:"status" example:"201" extensions:"x-order=1"`

	// เวลาที่ response ถูกสร้างขึ้น
	Timestamp time.Time `json:"timestamp" example:"2025-01-01T00:00:00.000000000Z" extensions:"x-order=2"`

	// ข้อความสถานะของ response
	Message string `json:"message" example:"Created" extensions:"x-order=3"`
}

type CreateDataResponseSwagger[T any] struct {
	// สถานะของ response
	Success bool `json:"success" example:"true" extensions:"x-order=0"`

	// รหัสสถานะของ response
	Status int `json:"status" example:"201" extensions:"x-order=1"`

	// เวลาที่ response ถูกสร้างขึ้น
	Timestamp time.Time `json:"timestamp" example:"2025-01-01T00:00:00.000000000Z" extensions:"x-order=2"`

	// ข้อความสถานะของ response
	Message string `json:"message" example:"Created" extensions:"x-order=3"`

	// ข้อมูลที่ต้องการส่งกลับ
	Data T `json:"data,omitempty" extensions:"x-order=4"`
}

// Update Response for Swagger (200)
type UpdateResponseSwagger struct {
	// สถานะของ response
	Success bool `json:"success" example:"true" extensions:"x-order=0"`

	// รหัสสถานะของ response
	Status int `json:"status" example:"200" extensions:"x-order=1"`

	// เวลาที่ response ถูกสร้างขึ้น
	Timestamp time.Time `json:"timestamp" example:"2025-01-01T00:00:00.000000000Z" extensions:"x-order=2"`

	// ข้อความสถานะของ response
	Message string `json:"message" example:"OK" extensions:"x-order=3"`
}

// Bad Request Response for Swagger (400)
type BadRequestResponseSwagger struct {
	// สถานะของ response
	Success bool `json:"success" example:"false" extensions:"x-order=0"`

	// รหัสสถานะของ response
	Status int `json:"status" example:"400" extensions:"x-order=1"`

	// เวลาที่ response ถูกสร้างขึ้น
	Timestamp time.Time `json:"timestamp" example:"2025-01-01T00:00:00.000000000Z" extensions:"x-order=2"`

	// ข้อความสถานะของ response
	Message string `json:"message" example:"Bad Request"  extensions:"x-order=3"`
}

// Unauthorized Response for Swagger (401)
type UnauthorizedResponseSwagger struct {
	// สถานะของ response
	Success bool `json:"success" example:"false" extensions:"x-order=0"`

	// รหัสสถานะของ response
	Status int `json:"status" example:"401" extensions:"x-order=1"`

	// เวลาที่ response ถูกสร้างขึ้น
	Timestamp time.Time `json:"timestamp" example:"2025-01-01T00:00:00.000000000Z" extensions:"x-order=2"`

	// ข้อความสถานะของ response
	Message string `json:"message" example:"Unauthorized" extensions:"x-order=3"`
}

// Forbidden Response for Swagger (403)
type ForbiddenResponseSwagger struct {
	// สถานะของ response
	Success bool `json:"success" example:"false" extensions:"x-order=0"`

	// รหัสสถานะของ response
	Status int `json:"status" example:"403" extensions:"x-order=1"`

	// เวลาที่ response ถูกสร้างขึ้น
	Timestamp time.Time `json:"timestamp" example:"2025-01-01T00:00:00.000000000Z" extensions:"x-order=2"`

	// ข้อความสถานะของ response
	Message string `json:"message" example:"Forbidden" extensions:"x-order=3"`
}

// Not Found Response for Swagger (404)
type NotFoundResponseSwagger struct {
	// สถานะของ response
	Success bool `json:"success" example:"false" extensions:"x-order=0"`

	// รหัสสถานะของ response
	Status int `json:"status" example:"404" extensions:"x-order=1"`

	// เวลาที่ response ถูกสร้างขึ้น
	Timestamp time.Time `json:"timestamp" example:"2025-01-01T00:00:00.000000000Z" extensions:"x-order=2"`

	// ข้อความสถานะของ response
	Message string `json:"message" example:"Not Found" extensions:"x-order=3"`
}

// Internal Server Error Response for Swagger (500)
type InternalServerErrorResponseSwagger struct {
	// สถานะของ response
	Success bool `json:"success" example:"false" extensions:"x-order=0"`

	// รหัสสถานะของ response
	Status int `json:"status" example:"500" extensions:"x-order=1"`

	// เวลาที่ response ถูกสร้างขึ้น
	Timestamp time.Time `json:"timestamp" example:"2025-01-01T00:00:00.000000000Z" extensions:"x-order=2"`

	// ข้อความสถานะของ response
	Message string `json:"message" example:"Internal Server Error" extensions:"x-order=3"`
}
