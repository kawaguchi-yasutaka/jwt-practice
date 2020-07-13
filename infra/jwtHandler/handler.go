package jwtHandler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"jwt-practice/interfaces"
	"time"
)

type JwtHandler struct {
	key []byte
}

func NewJwtHandler(key []byte) interfaces.IJwtHandler {
	return JwtHandler{
		key: key,
	}
}

func (handler JwtHandler) Verify(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(handler.key)
	})
	if err != nil {
		return err
	}
	data := token.Claims.(jwt.MapClaims)
	if ok := data.VerifyExpiresAt(time.Now().Unix(), true); !ok {
		return fmt.Errorf("tokne expired")
	}
	return nil
}
