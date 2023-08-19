package response

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserResponseVerify struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
