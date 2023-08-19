package request

type UserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Status    int    `json:"status"`
	CreatedBy string `json:"created_by"`
}
