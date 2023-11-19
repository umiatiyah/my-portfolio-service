package service

import (
	"context"
	"log"
	"portfolio-go/infra"
	"portfolio-go/model"
	"portfolio-go/response"
)

func GetUserProfile(c context.Context, id int) (response.UserProfileResponse, error) {
	qry := "SELECT up.* FROM user_profile up WHERE id = $1"
	users, err := infra.DB.Query(qry, id)
	if err != nil {
		log.Println("error querying", qry)
		return response.UserProfileResponse{}, err
	}

	var user response.ProfileResponse
	var userList response.UserProfileResponse
	for users.Next() {
		users.Scan(&user.Id, &user.UserId, &user.Fullname, &user.Country, &user.Address, &user.IsFreelance, &user.PhoneNumber, &user.Status,
			&user.CreatedBy, &user.CreatedAt, &user.UpdatedBy, &user.UpdatedAt, &user.Description)
	}

	qry2 := "select (select e.title from experience e where e.user_profile_id = $1 order by e.sequence limit 1) as current_job, (select sm.link from social_media sm where sm.title = 'github' and sm.user_profile_id = $1 limit 1) as github_link"
	profile, err := infra.DB.Query(qry2, id)
	if err != nil {
		log.Println("error querying", qry2)
		return response.UserProfileResponse{}, err
	}

	for profile.Next() {
		profile.Scan(&user.CurrentJob, &user.GithubLink)
	}

	userList.UserProfile = user

	socialMedias, err := GetUserSocialMedia(c, id)
	hobbies, err := GetUserHobby(c, id)
	experience, err := GetUserExperience(c, id)
	education, err := GetUserEducation(c, id)
	userList.SocialMedia = socialMedias
	userList.Hobby = hobbies
	userList.Experience = experience
	userList.Education = education
	return userList, nil
}

func GetUserSocialMedia(c context.Context, userProfileId int) ([]model.SocialMedia, error) {
	qry := "SELECT sm.* FROM social_media sm WHERE user_profile_id = $1"
	socials, err := infra.DB.Query(qry, userProfileId)
	if err != nil {
		log.Println("error querying")
		return nil, err
	}

	var socialMedias []model.SocialMedia
	for socials.Next() {
		var social model.SocialMedia
		socials.Scan(&social.Id, &social.UserProfileId, &social.Title, &social.Link, &social.Status,
			&social.CreatedBy, &social.CreatedAt, &social.UpdatedBy, &social.UpdatedAt)
		socialMedias = append(socialMedias, social)
	}
	return socialMedias, nil
}

func GetUserHobby(c context.Context, userProfileId int) ([]model.Hobby, error) {
	qry := "SELECT h.* FROM hobby h WHERE user_profile_id = $1 ORDER BY Sequence"
	hobbies, err := infra.DB.Query(qry, userProfileId)
	if err != nil {
		log.Println("error querying")
		return nil, err
	}

	var hobbyList []model.Hobby
	for hobbies.Next() {
		var hobby model.Hobby
		hobbies.Scan(&hobby.Id, &hobby.UserProfileId, &hobby.Title, &hobby.Description, &hobby.Sequence, &hobby.Status,
			&hobby.CreatedBy, &hobby.CreatedAt, &hobby.UpdatedBy, &hobby.UpdatedAt)
		hobbyList = append(hobbyList, hobby)
	}
	return hobbyList, nil
}

func GetUserExperience(c context.Context, userProfileId int) ([]model.Experience, error) {
	qry := "SELECT e.* FROM experience e WHERE user_profile_id = $1 ORDER BY Sequence"
	experiences, err := infra.DB.Query(qry, userProfileId)
	if err != nil {
		log.Println("error querying")
		return nil, err
	}

	var experienceList []model.Experience
	for experiences.Next() {
		var experience model.Experience
		experiences.Scan(&experience.Id, &experience.UserProfileId, &experience.Title, &experience.StartDate, &experience.EndDate,
			&experience.CompanyName, &experience.Description, &experience.Sequence, &experience.Status,
			&experience.CreatedBy, &experience.CreatedAt, &experience.UpdatedBy, &experience.UpdatedAt)
		experienceList = append(experienceList, experience)
	}
	return experienceList, nil
}

func GetUserEducation(c context.Context, userProfileId int) ([]model.Education, error) {
	qry := "SELECT e.* FROM education e WHERE user_profile_id = $1 ORDER BY Sequence"
	educations, err := infra.DB.Query(qry, userProfileId)
	if err != nil {
		log.Println("error querying")
		return nil, err
	}

	var educationList []model.Education
	for educations.Next() {
		var education model.Education
		educations.Scan(&education.Id, &education.UserProfileId, &education.Title, &education.StartDate, &education.EndDate,
			&education.CampusName, &education.Description, &education.Sequence, &education.Status,
			&education.CreatedBy, &education.CreatedAt, &education.UpdatedBy, &education.UpdatedAt)
		educationList = append(educationList, education)
	}
	return educationList, nil
}
