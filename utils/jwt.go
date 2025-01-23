package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "f9df2d30544bc5591b090483529169a9f86ba3b20c02737ba1c1b52b7789570f956b7b3f146d213566f49365a6abe0a63fe47df7c1bd2ca030642a8953be18f3a4574fd8efdc328f9a9624fb85b6903fe11efc97c4677c454e91e2963990cae8d72b5a54873ed02ad8aadd48479f8a07695c676c1c0835e13ecc765349a205a24ed88bf14618fb5e30c01e19663062147fdc4ef51b10d3e05b767a9b399c697272d65679a756fe153fe94340e8c3d3e1a9b0863bbe93bee1a69abc0d54ce712ec06c056db1d522b2301f0b5a2bd35b80226e49a5672feae57cc27e35ebc6f34c2fd29fe2d0e6c5c220697e7e92964b76dc437bb7cf3997f44a133e6cf0e9da8e"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	isValidToken := parsedToken.Valid

	if !isValidToken {
		return 0, errors.New("Invalid Token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid Token Claims")
	}

	// email := claims["eamil"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
