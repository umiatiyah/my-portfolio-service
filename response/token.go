package response

type Token struct {
	Token  string `json:"token"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}
