package model

import "gopkg.in/guregu/null.v4"

type Education struct {
	Id            int         `json:"id"`
	UserProfileId int         `json:"user_profile_id"`
	Title         string      `json:"title"`
	StartDate     string      `json:"start_date"`
	EndDate       string      `json:"end_date"`
	CampusName    string      `json:"campus_name"`
	Description   null.String `json:"description"`
	Sequence      int         `json:"sequence"`
	Status        int         `json:"status"`
	CreatedBy     string      `json:"created_by"`
	CreatedAt     string      `json:"created_at"`
	UpdatedBy     null.String `json:"updated_by"`
	UpdatedAt     null.Time   `json:"updated_at"`
}
