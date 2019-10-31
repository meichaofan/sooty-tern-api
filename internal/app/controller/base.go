package controller

import (
	"go.uber.org/dig"
	"sooty-tern/internal/app/controller/api"
)

func Inject(container *dig.Container) error {
	container.Provide(api.NewUser)
	container.Provide(api.NewLoginInfo)
	return nil
}
