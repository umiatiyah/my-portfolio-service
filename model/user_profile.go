package model

type UserProfile struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Fullname    string `json:"fullname"`
	Country     string `json:"country"`
	Address     string `json:"address"`
	IsFreelance bool   `json:"is_freelance"`
	PhoneNumber string `json:"phone_number"`
	Status      int    `json:"status"`
	CreatedBy   string `json:"created_by"`
	CreatedAt   string `json:"created_at"`
	UpdatedBy   string `json:"updated_by"`
	UpdatedAt   string `json:"updated_at"`
}
