package server

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	secretKey string
	issuer string
}

func NewJwtService() *JwtService {
	return &JwtService{
		secretKey: "secret",
		issuer: "todo-list",
	}
}

type Claim struct {
	sum string `json:"sum"`
	jwt.StandardClaims
}

func (js *JwtService) GenerateToken(id string) (string, error) {
	const DAY = 24 * time.Hour
	cl := &Claim{
		sum: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(DAY).Unix(),
			Issuer: js.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cl)

	t, err := token.SignedString([]byte(js.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (js *JwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(js.secretKey), nil
	})

	return err == nil
}