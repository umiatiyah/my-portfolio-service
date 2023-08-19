package service

import (
	"context"
	"log"
	"portfolio-go/infra"
	"portfolio-go/middleware"
	"portfolio-go/request"
	"portfolio-go/response"
	"portfolio-go/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c context.Context) ([]response.UserResponse, error) {
	qry := "SELECT username, email FROM users"
	users, err := infra.DB.Query(qry)
	if err != nil {
		log.Println("error querying")
		return nil, err
	}

	var userList []response.UserResponse
	for users.Next() {
		var user response.UserResponse
		users.Scan(&user.Username, &user.Email)
		userList = append(userList, user)
	}
	return userList, nil
}

func Register(c context.Context, user request.UserRequest) (response.UserResponse, error) {
	qry := `INSERT INTO users (username, password, email, status, created_by, created_at) 
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING username, email`
	userInserted, err := infra.DB.Query(qry, user.Username, utils.HashAndSalt([]byte(user.Password)), user.Email, user.Status, user.CreatedBy, time.Now())
	if err != nil {
		log.Println("error inserting user")
	}

	var userResponse response.UserResponse
	for userInserted.Next() {
		userInserted.Scan(&userResponse.Username, &userResponse.Email)
	}
	return userResponse, err
}

func Login(c context.Context, auth request.AuthRequest) (response.TokenResponse, error) {
	qry := "SELECT id, username, email, password FROM users WHERE username = $1"
	row, err := infra.DB.Query(qry, auth.Username)
	if err != nil {
		log.Println("error get user")
		return response.TokenResponse{}, err
	}

	var user response.UserResponseVerify
	for row.Next() {
		row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	}

	err = utils.VerifyPassword(user.Password, auth.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return response.TokenResponse{}, err
	}
	token, _ := middleware.CreateToken(user.Id, auth.Username)
	return token, nil
}
