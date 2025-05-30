package customers

type GetUsersCustomerByIdResponse struct {
	UserId      int    `json:"user_id"`     // user_id
	UserName    string `json:"user_name"`   // user_name
	Name        string `json:"name"`        // name
	Description string `json:"description"` // description
}
