package web

import (
	"jwt-practice/infra"
	"jwt-practice/web/handler"
	"net/http"
)

func Init(infra infra.Infra) {
	_handler := handler.NewHandler(infra.JwtGenerator, infra.JwtHandler)
	http.HandleFunc("/login", _handler.Login)
	http.HandleFunc("/check_token", _handler.CheckToken)
	http.ListenAndServe(":8080", nil)
}
