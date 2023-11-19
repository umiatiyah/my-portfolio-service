package model

import "gopkg.in/guregu/null.v4"

type Experience struct {
	Id            int         `json:"id"`
	UserProfileId int         `json:"user_profile_id"`
	Title         string      `json:"title"`
	StartDate     string      `json:"start_date"`
	EndDate       null.Time   `json:"end_date"`
	CompanyName   string      `json:"company_name"`
	Description   null.String `json:"description"`
	Sequence      int         `json:"sequence"`
	Status        int         `json:"status"`
	CreatedBy     string      `json:"created_by"`
	CreatedAt     string      `json:"created_at"`
	UpdatedBy     null.String `json:"updated_by"`
	UpdatedAt     null.Time   `json:"updated_at"`
}
