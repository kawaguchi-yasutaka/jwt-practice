package main

import (
	"jwt-practice/infra"
	"jwt-practice/web"
)

func main() {
	infra := infra.InitInfra()
	web.Init(infra)
}
