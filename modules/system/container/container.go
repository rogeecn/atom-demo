package container

import (
	"atom/modules/system/controller"
	"atom/modules/system/routes"

	"github.com/rogeecn/atom/container"
	"go.uber.org/dig"
)

func Provide() error {
	if err := container.Container.Provide(controller.NewCaptchaController); err != nil {
		return err
	}

	if err := container.Container.Provide(routes.NewRoute, dig.Group("route")); err != nil {
		return err
	}
	return nil
}
