package auth_helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("6E+60q835r7P")

func CreateToken(user_id int, user_token, user_name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token":   user_token,
		"uid":     user_id,
		"name":    user_name,
		"created": time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(secretKey)
}

func OpenToken(token string) (*AuthResult, error) {
	var parsed_token *jwt.Token
	parsed_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// _, ok := token.Method.(*jwt.SigningMethodECDSA)
		// if !ok {
		// 	return nil, nil
		// }
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if parsed_token == nil {
		return nil, errors.New("token is null")
	}
	claims, ok := parsed_token.Claims.(jwt.MapClaims)
	if !ok || !parsed_token.Valid {
		return nil, errors.New("token is not valid")
	}
	result := AuthResult{}
	var uid float64
	if user_id_claim, ok := claims["uid"]; !ok {
		return nil, errors.New("uid not parsed")
	} else if uid, ok = user_id_claim.(float64); !ok {
		return nil, errors.New("uid not parsed")
	} else {
		result.Uid = int(uid)
	}
	if token_claim, ok := claims["token"]; !ok {
		return nil, errors.New("token not parsed")
	} else if result.Token, ok = token_claim.(string); !ok {
		return nil, errors.New("token not parsed")
	}
	if name_claim, ok := claims["name"]; !ok {
		return nil, errors.New("name not parsed")
	} else if result.Name, ok = name_claim.(string); !ok {
		return nil, errors.New("name not parsed")
	}
	if date_claim, ok := claims["created"]; !ok {
		return nil, errors.New("created not parsed")
	} else if date, ok := date_claim.(float64); !ok {
		return nil, errors.New("created not parsed")
	} else {
		result.Created = time.Unix(int64(date), 0)
	}
	return &result, nil
}
