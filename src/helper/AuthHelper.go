package helper

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccessTokenExpires  int64
	RefreshTokenExpires int64
}

// createToken creates a new token for a specific username
func CreateToken(username string) (*TokenDetails, error) {
	ACCESS_SECRET := os.Getenv("ACCESS_SECRET")
	REFRESH_SECRET := os.Getenv("REFRESH_SECRET")
	var err error
	token := &TokenDetails{}
	token.AccessTokenExpires = time.Now().Add(time.Minute * 15).Unix()
	token.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()

	// Create Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = token.AccessTokenExpires

	jwtAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = jwtAccessToken.SignedString([]byte(ACCESS_SECRET))
	if err != nil {
		return nil, err
	}

	// Create Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["username"] = username
	rtClaims["exp"] = token.RefreshTokenExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token.RefreshToken, err = rt.SignedString([]byte(REFRESH_SECRET))
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Create Password Hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Compare Password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	if len(tokenString) < 1 {
		return nil, errors.New("Token not given!")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
