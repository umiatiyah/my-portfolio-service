package model

import "gopkg.in/guregu/null.v4"

type Contact struct {
	Id            int         `json:"id"`
	UserProfileId int         `json:"user_profile_id"`
	Address       string      `json:"address"`
	ContactPerson string      `json:"contact_person"`
	Sequence      int         `json:"sequence"`
	Status        int         `json:"status"`
	CreatedBy     string      `json:"created_by"`
	CreatedAt     string      `json:"created_at"`
	UpdatedBy     null.String `json:"updated_by"`
	UpdatedAt     null.String `json:"updated_at"`
}
