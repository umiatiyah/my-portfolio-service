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
		log.Println("error querying")
		return response.UserProfileResponse{}, err
	}

	var userList response.UserProfileResponse
	for users.Next() {
		var user model.UserProfile
		users.Scan(&user.Id, &user.UserId, &user.Fullname, &user.Country, &user.Address, &user.IsFreelance, &user.PhoneNumber, &user.Status,
			&user.CreatedBy, &user.CreatedAt, &user.UpdatedBy, &user.UpdatedAt)
		userList.UserProfile = user
	}

	socialMedias, err := GetUserSocialMedia(c, id)
	hobbies, err := GetUserHobby(c, id)
	experience, err := GetUserExperience(c, id)
	education, err := GetUserEducation(c, id)
	contacts, err := GetUserContact(c, id)
	userList.SocialMedia = socialMedias
	userList.Hobby = hobbies
	userList.Experience = experience
	userList.Education = education
	userList.Contact = contacts
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
	qry := "SELECT h.* FROM hobby h WHERE user_profile_id = $1"
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
	qry := "SELECT e.* FROM experience e WHERE user_profile_id = $1"
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
	qry := "SELECT e.* FROM education e WHERE user_profile_id = $1"
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

func GetUserContact(c context.Context, userProfileId int) ([]model.Contact, error) {
	qry := "SELECT c.* FROM contact c WHERE user_profile_id = $1"
	contacts, err := infra.DB.Query(qry, userProfileId)
	if err != nil {
		log.Println("error querying")
		return nil, err
	}

	var contactList []model.Contact
	for contacts.Next() {
		var contact model.Contact
		contacts.Scan(&contact.Id, &contact.UserProfileId, &contact.Address, &contact.ContactPerson, &contact.Sequence, &contact.Status,
			&contact.CreatedBy, &contact.CreatedAt, &contact.UpdatedBy, &contact.UpdatedAt)
		contactList = append(contactList, contact)
	}
	return contactList, nil
}
