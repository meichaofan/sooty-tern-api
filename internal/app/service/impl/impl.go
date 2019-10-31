package impl

import (
	"go.uber.org/dig"
	"sooty-tern/internal/app/service"
	"sooty-tern/internal/app/service/impl/internal"
)

func Inject(container *dig.Container) error {
	container.Provide(internal.NewUser, dig.As(new(service.IUser)))
	container.Provide(internal.NewLoginInfo, dig.As(new(service.ILoginInfo)))
	return nil
}
