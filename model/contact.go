package model

type Contact struct {
	Id            int    `json:"id"`
	UserProfileId int    `json:"user_profile_id"`
	Address       string `json:"address"`
	ContactPerson string `json:"contact_person"`
	Sequence      int    `json:"sequence"`
	Status        int    `json:"status"`
	CreatedBy     string `json:"created_by"`
	CreatedAt     string `json:"created_at"`
	UpdatedBy     string `json:"updated_by"`
	UpdatedAt     string `json:"updated_at"`
}
