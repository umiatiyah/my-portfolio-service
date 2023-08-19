package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"portfolio-go/response"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	SECRET_KEY = os.Getenv("SECRET_KEY")
)

func CreateToken(id int, name string) (response.TokenResponse, error) {
	if id == 0 {
		return response.TokenResponse{}, nil
	}
	expiredTime, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION"))
	if err != nil {
		expiredTime = 1 //Token expires after 1 hour
	}

	SECRET_AUDIENCE := os.Getenv("SECRET_AUDIENCE")
	claims := &response.JWTClaim{
		Username: name,
		StandardClaims: jwt.StandardClaims{
			Subject:   "JWT Portfolio",
			Id:        strconv.Itoa(id),
			Audience:  SECRET_AUDIENCE,
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiredTime)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Println(err)
	}

	return response.TokenResponse{
		Username:    claims.Username,
		TokenAccess: token,
		TokenType:   "Bearer",
		ExpiresAt:   claims.StandardClaims.ExpiresAt,
	}, err
}

func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.Header("Content-Type", "application/json")
			response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			log.Println("Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
