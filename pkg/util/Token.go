package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/SGDIEGO/CleanCode/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

/*CONSTANTS*/
const (
	HeaderName = "token"
)

// Function to create JWT
func CreateToken(UserLog *entity.User, secret []byte) (string, error) {
	TimeExpired := time.Now().Add(5 * time.Minute)

	Userclaim := entity.UserClaim{
		UserName: UserLog.UserName,
		Email:    UserLog.Email,
		Password: UserLog.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(TimeExpired),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Userclaim)

	return token.SignedString(secret)
}

func GetTokenValue(requestToken string, secret []byte) (*entity.User, error) {

	var UserC entity.UserClaim
	token, err := jwt.ParseWithClaims(requestToken, &UserC, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	User := entity.User{
		UserId:   UserC.UserId,
		UserName: UserC.UserName,
		Email:    UserC.Email,
		Password: UserC.Password,
	}

	return &User, nil
}

func TokenValid(c *gin.Context, secretKey []byte) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	// Validate
	token := c.Query(HeaderName)
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
