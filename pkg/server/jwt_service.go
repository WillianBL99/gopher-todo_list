package server

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/willianbl99/todo-list_api/pkg/herr"
)

type JwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() *JwtService {
	return &JwtService{
		secretKey: "secret",
		issuer:    "todo-list",
	}
}

type Claim struct {
	Sum string `json:"sum"`
	jwt.StandardClaims
}

func (js *JwtService) GenerateToken(id string) (string, error) {
	const ONE_DAY = 24 * time.Hour
	token := jwt.New(jwt.SigningMethodHS256)

	cl := token.Claims.(jwt.MapClaims)
	cl["sum"] = id
	cl["exp"] = time.Now().Add(ONE_DAY).Unix()
	cl["iat"] = time.Now().Unix()
	cl["iss"] = js.issuer

	t, err := token.SignedString([]byte(js.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (js *JwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, js.jwtValidate)

	return err == nil
}

func (js *JwtService) GetTokenId(token string) (string, error) {
	tk, err := jwt.ParseWithClaims(token, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(js.secretKey), nil
	})

	if err != nil {
		return "", herr.NewApp().InvalidToken
	}

	if claims, ok := tk.Claims.(*Claim); ok && tk.Valid {
		return claims.Sum, nil
	}

	return "", herr.NewApp().InvalidToken
}

func (js *JwtService) jwtValidate(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, herr.NewApp().InvalidToken
	}

	return []byte(js.secretKey), nil
}
