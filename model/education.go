package model

type Education struct {
	Id            int    `json:"id"`
	UserProfileId int    `json:"user_profile_id"`
	Title         string `json:"title"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	CampusName    string `json:"campus_name"`
	Description   string `json:"description"`
	Sequence      int    `json:"sequence"`
	Status        int    `json:"status"`
	CreatedBy     string `json:"created_by"`
	CreatedAt     string `json:"created_at"`
	UpdatedBy     string `json:"updated_by"`
	UpdatedAt     string `json:"updated_at"`
}
