package web

import (
	"jwt-practice/infra"
	"jwt-practice/web/handler"
	"net/http"
)

func Init(infra infra.Infra) {
	_handler := handler.NewHandler(infra.JwtGenerator)
	http.HandleFunc("/login", _handler.LoginHandler)
	http.ListenAndServe(":8080", nil)
}
