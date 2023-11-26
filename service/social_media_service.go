package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"portfolio-go/infra"
	"portfolio-go/model"
	"time"
)

func GetSocialMediasByUser(c context.Context, username string) ([]model.SocialMedia, error) {
	qry := `SELECT sm.*FROM social_media sm
	JOIN user_profile up ON sm.user_profile_id = up.id
	JOIN users u ON up.user_id = u.id WHERE u.username = $1`
	socialMedias, err := infra.DB.Query(qry, username)
	if err != nil {
		log.Println("error querying social media: ", err)
		return nil, err
	}

	var socialMediaList []model.SocialMedia
	for socialMedias.Next() {
		var socialMedia model.SocialMedia
		socialMedias.Scan(&socialMedia.Id, &socialMedia.UserProfileId, &socialMedia.Title, &socialMedia.Link, &socialMedia.Status,
			&socialMedia.CreatedBy, &socialMedia.CreatedAt, &socialMedia.UpdatedBy, &socialMedia.UpdatedAt)
		socialMediaList = append(socialMediaList, socialMedia)
	}

	if socialMediaList == nil {
		return nil, errors.New("data not found")
	}

	return socialMediaList, nil
}

func GetDetailSocialMedia(c context.Context, id int) (model.SocialMedia, error) {
	qry := `SELECT sm.*FROM social_media sm WHERE sm.Id = $1`
	detSocialMedia, err := infra.DB.Query(qry, id)
	if err != nil {
		log.Println("error querying social media: ", err)
		return model.SocialMedia{}, err
	}

	var socialMedia model.SocialMedia
	for detSocialMedia.Next() {
		detSocialMedia.Scan(&socialMedia.Id, &socialMedia.UserProfileId, &socialMedia.Title, &socialMedia.Link, &socialMedia.Status,
			&socialMedia.CreatedBy, &socialMedia.CreatedAt, &socialMedia.UpdatedBy, &socialMedia.UpdatedAt)
	}
	if socialMedia.Id == 0 {
		return model.SocialMedia{}, errors.New("data not found")
	}

	return socialMedia, nil
}

func CreateSocialMedia(c context.Context, socialMedia model.SocialMedia) (model.SocialMedia, error) {
	qry := `INSERT INTO social_media (user_profile_id, title, link, status, created_by, created_at) 
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	smInserted, err := infra.DB.Query(qry, socialMedia.UserProfileId, socialMedia.Title, socialMedia.Link, socialMedia.Status, socialMedia.CreatedBy, time.Now())
	if err != nil {
		log.Println("error inserting social media: ", err)
	}

	for smInserted.Next() {
		smInserted.Scan(&socialMedia.UserProfileId, &socialMedia.Title, &socialMedia.Link, &socialMedia.Status,
			&socialMedia.CreatedBy, &socialMedia.CreatedAt, &socialMedia.UpdatedBy, &socialMedia.UpdatedAt)
	}
	return socialMedia, err
}

func UpdateSocialMedia(c context.Context, socialMedia model.SocialMedia) (model.SocialMedia, error) {
	qry := `UPDATE social_media
			SET user_profile_id=$2, title=$3, link=$4, status=$5, updated_by=$6, updated_at=$7
			WHERE id=$1 RETURNING *`
	smUpdated, err := infra.DB.Query(qry, socialMedia.Id, socialMedia.UserProfileId, socialMedia.Title, socialMedia.Link, socialMedia.Status, socialMedia.UpdatedBy, time.Now())
	if err != nil {
		log.Println("error updating social media: ", err)
	}

	for smUpdated.Next() {
		smUpdated.Scan(&socialMedia.UserProfileId, &socialMedia.Title, &socialMedia.Link, &socialMedia.Status,
			&socialMedia.CreatedBy, &socialMedia.CreatedAt, &socialMedia.UpdatedBy, &socialMedia.UpdatedAt)
	}
	return socialMedia, err
}

func DeleteSocialMedia(c context.Context, id int) (string, error) {
	_, err := GetDetailSocialMedia(c, id)
	if err != nil {
		return "", errors.New("data not found")
	}

	qry := `DELETE FROM social_media WHERE id=$1`
	_, err = infra.DB.Query(qry, id)
	if err != nil {
		log.Println("error deleting social media: ", err)
	}

	msg := fmt.Sprintf("social media with id %d successfully deleted", id)
	return msg, err
}
