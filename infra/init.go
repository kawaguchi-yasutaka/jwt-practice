package infra

import (
	"io/ioutil"
	"jwt-practice/infra/jwtGenerator"
	"jwt-practice/infra/jwtHandler"
	"jwt-practice/interfaces"
	"os"
)

type Infra struct {
	JwtGenerator interfaces.IJwtGenerator
	JwtHandler   interfaces.IJwtHandler
}

func InitInfra() Infra {
	fSKey, err := os.Open("./demo")
	if err != nil {
		panic(err)
	}
	sKey, err := ioutil.ReadAll(fSKey)
	if err != nil {
		panic(err)
	}
	defer fSKey.Close()
	fPKey, err := os.Open("./demo.pub")
	if err != nil {
		panic(err)
	}
	pKey, err := ioutil.ReadAll(fPKey)
	if err != nil {
		panic(err)
	}
	defer fPKey.Close()
	jwtGenerator := jwtGenerator.NewJwtGenerator(sKey)
	jwtHandler := jwtHandler.NewJwtHandler(pKey)
	return Infra{
		JwtGenerator: jwtGenerator,
		JwtHandler:   jwtHandler,
	}
}
