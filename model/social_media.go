package model

import "gopkg.in/guregu/null.v4"

type SocialMedia struct {
	Id            int         `json:"id"`
	UserProfileId int         `json:"user_profile_id"`
	Title         string      `json:"title"`
	Link          string      `json:"link"`
	Status        int         `json:"status"`
	CreatedBy     string      `json:"created_by"`
	CreatedAt     string      `json:"created_at"`
	UpdatedBy     null.String `json:"updated_by"`
	UpdatedAt     null.Time   `json:"updated_at"`
}
