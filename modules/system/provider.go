package system

import (
	"atom/http/modules/system/controller"
	"atom/http/modules/system/dao"
	"atom/http/modules/system/routes"
	"atom/http/modules/system/service"

	"github.com/rogeecn/atom/container"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: dao.Provide},
		{Provider: service.Provide},
		{Provider: controller.Provide},
		{Provider: routes.Provide},
	}
}
