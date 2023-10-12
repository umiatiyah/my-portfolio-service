package model

type Hobby struct {
	Id            int    `json:"id"`
	UserProfileId int    `json:"user_profile_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Sequence      int    `json:"sequence"`
	Status        int    `json:"status"`
	CreatedBy     string `json:"created_by"`
	CreatedAt     string `json:"created_at"`
	UpdatedBy     string `json:"updated_by"`
	UpdatedAt     string `json:"updated_at"`
}
