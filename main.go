package main

import (
	"github.com/mniudanri/go-auth-paseto/api"
	"github.com/mniudanri/go-auth-paseto/util"
)

func main() {
	config := util.LoadConfig(".")

	server := api.InitServer(config)
	_ = server
}