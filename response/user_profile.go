package response

type UserProfileResponse struct {
	UserProfile interface{} `json:"user_profile"`
	SocialMedia interface{} `json:"social_media"`
	Hobby       interface{} `json:hobby"`
	Experience  interface{} `json:"experience"`
	Education   interface{} `json:"education"`
	Contact     interface{} `json:"contact"`
}
