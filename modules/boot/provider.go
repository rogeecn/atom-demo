package boot

import (
	"atom/http/docs"

	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom-addons/providers/swagger"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: provideSwagger},
	}
}

func provideSwagger(opts ...opt.Option) error {
	return container.Container.Provide(func(swagger *swagger.Swagger) contracts.Initial {
		swagger.Load(docs.SwaggerSpec)
		return nil
	}, atom.GroupInitial)
}
