package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/typings"
)

type AccountTokenStruct struct {
	jwt.Claims
	UserID      string               `json:"userID"`
	Email       string               `json:"email"`
	IssuedAt    string               `json:"issuedAt"`
	AccountType typings.AccountType  `json:"accountType"`
	AccountRole entities.AccountRole `json:"accountRole"`
}

var jwtSecretKey = []byte(config.EnvConfig.JWT_SECRET)

func IssueToken(ats *AccountTokenStruct) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":         ats.UserID,
		"userID":      ats.UserID,
		"accountRole": ats.AccountRole,
		"iss":         "luxxe",
		"email":       ats.Email,
		"aud":         ats.AccountType,                                                                    // Audience (user role)                                                                        // Issuer
		"exp":         time.Now().Add(time.Duration(config.EnvConfig.JWT_EXPIRY) * 24 * time.Hour).Unix(), // Expiration time
		"iat":         strconv.Itoa(int(time.Now().Unix())), 
	})
	tokenString, err := claims.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func VerifyToken(tokenString string) (*AccountTokenStruct, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var currentAccountRole entities.AccountRole
		accountRoleClaims := claims["accountRole"].(string)
		if accountRoleClaims == string(entities.AccountRoleUser) {
			currentAccountRole = entities.AccountRoleUser
		} else {
			currentAccountRole = entities.AccountRoleAdmin
		}

		accountTokenStruct := &AccountTokenStruct{
			UserID:      claims["userID"].(string),
			AccountRole: currentAccountRole,
			IssuedAt:    claims["iat"].(string),
			Email:       claims["email"].(string),
			AccountType: typings.AccountType(claims["aud"].(string)),
		}
		return accountTokenStruct, nil
	}

	return nil, fmt.Errorf("unable to retrieve token")
}
