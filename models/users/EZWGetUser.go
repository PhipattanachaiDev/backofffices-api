package models

type EZWGetUser struct {
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
