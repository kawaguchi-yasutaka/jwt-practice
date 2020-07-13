package handler

import (
	"encoding/json"
	"io/ioutil"
	"jwt-practice/interfaces"
	"jwt-practice/web/request"
	"jwt-practice/web/response"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	jwtGenerator interfaces.IJwtGenerator
}

func NewHandler(jwtGenerator interfaces.IJwtGenerator) Handler {
	return Handler{
		jwtGenerator: jwtGenerator,
	}
}

const (
	USER_NAME     = "master"
	PASSWORD      = "password"
	TOKEN_EXPIRED = 60 * 60 * 24
)

func (handler Handler) LoginHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var loginRequest request.LoginRequest
	if err := json.Unmarshal(body, &loginRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !(loginRequest.Name == USER_NAME && loginRequest.Password == PASSWORD) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := handler.jwtGenerator.GenerateToken(map[string]interface{}{
		"sub": loginRequest.Name,
		"exp": time.Now().Unix() + TOKEN_EXPIRED,
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(&response.LoginResponse{Token: token})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
