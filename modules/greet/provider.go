package greet

import (
	"atom/http/modules/greet/controller"

	"github.com/rogeecn/atom/container"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: controller.Provide},
	}
}
