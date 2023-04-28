package system

import (
	"atom/http/modules/system/controller"
	"atom/http/modules/system/routes"

	"github.com/rogeecn/atom/container"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: controller.Provide},
		{Provider: routes.Provide},
	}
}
