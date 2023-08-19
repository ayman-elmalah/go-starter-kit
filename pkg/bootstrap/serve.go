package bootstrap

import (
	"go-starter-kit/pkg/config"
	"go-starter-kit/pkg/database"
	"go-starter-kit/pkg/routing"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()
}
