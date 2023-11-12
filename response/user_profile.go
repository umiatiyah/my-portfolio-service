package response

import (
	"portfolio-go/model"
)

type UserProfileResponse struct {
	UserProfile ProfileResponse     `json:"user_profile"`
	SocialMedia []model.SocialMedia `json:"social_media"`
	Hobby       []model.Hobby       `json:"hobby"`
	Experience  []model.Experience  `json:"experience"`
	Education   []model.Education   `json:"education"`
	Contact     []model.Contact     `json:"contact"`
}

type ProfileResponse struct {
	model.UserProfile
	CurrentJob string `json:"current_job"`
	GithubLink string `json:"github_link"`
}
