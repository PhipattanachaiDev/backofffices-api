package users

type UpdateUserRequest struct {
	UserId      int    `json:"user_id" binding:"required" example:"1" extensions:"x-order=0"`
	Name        string `json:"name" binding:"required" example:"John Doe" extensions:"x-order=1"`
	Username    string `json:"username" binding:"required" example:"johndoe" extensions:"x-order=2"`
	Password    string `json:"password" binding:"required" example:"password123" extensions:"x-order=3"`
	Description string `json:"description" example:"Customer description" extensions:"x-order=4"`
}
