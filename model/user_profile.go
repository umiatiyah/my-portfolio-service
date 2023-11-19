package model

import "gopkg.in/guregu/null.v4"

type UserProfile struct {
	Id          int         `json:"id"`
	UserId      int         `json:"user_id"`
	Fullname    string      `json:"fullname"`
	Country     string      `json:"country"`
	Address     string      `json:"address"`
	IsFreelance null.Bool   `json:"is_freelance"`
	PhoneNumber null.String `json:"phone_number"`
	Status      int         `json:"status"`
	CreatedBy   string      `json:"created_by"`
	CreatedAt   string      `json:"created_at"`
	UpdatedBy   null.String `json:"updated_by"`
	UpdatedAt   null.Time   `json:"updated_at"`
	Description null.String `json:"description"`
}
