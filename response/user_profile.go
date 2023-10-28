package response

import (
	"portfolio-go/model"
)

type UserProfileResponse struct {
	UserProfile model.UserProfile   `json:"user_profile"`
	SocialMedia []model.SocialMedia `json:"social_media"`
	Hobby       []model.Hobby       `json:"hobby"`
	Experience  []model.Experience  `json:"experience"`
	Education   []model.Education   `json:"education"`
	Contact     []model.Contact     `json:"contact"`
}
