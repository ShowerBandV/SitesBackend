package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type ResultSet struct {
	Data   string
	Msg    string
	Status int
}
type LoginResult struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := &LoginResult{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "xuxi",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte("fiberSites"))

	return token, err
}
func ParseToken(token string) (bool, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &LoginResult{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("fiberSites"), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*LoginResult); ok && tokenClaims.Valid {
			fmt.Println(claims)
			return true, nil
		}
	}

	return false, err
}
