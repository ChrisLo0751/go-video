package dto

type UserResponse struct {
	ID    uint   `json:"id"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
