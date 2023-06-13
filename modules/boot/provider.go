package boot

import (
	"atom/http/modules/boot/controller"

	"github.com/rogeecn/atom/container"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: controller.Provide},
	}
}
