package dto

type CreateUserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int8   `json:"age"`
}

type CreateUserResponse struct {
	UserID int64 `json:"user_id"`
}
