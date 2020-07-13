package jwtGenerator

import (
	"github.com/dgrijalva/jwt-go"
	"jwt-practice/interfaces"
)

type JwtGenerator struct {
	key []byte
}

func NewJwtGenerator(key []byte) interfaces.IJwtGenerator {
	return JwtGenerator{
		key: key,
	}
}

func (generator JwtGenerator) GenerateToken(content map[string]interface{}) (string, error) {
	//token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims(content))
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	for k, v := range content {
		claims[k] = v
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(generator.key)
	if err != nil {
		return "", err
	}
	return token.SignedString(key)
}
