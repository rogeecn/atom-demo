package controller

import (
	"atom/http/docs"

	"github.com/rogeecn/atom-addons/providers/swagger"
	"github.com/rogeecn/atom/contracts"
)

func newInitializer(
	swagger *swagger.Swagger,
) contracts.Initial {
	swagger.Load(docs.SwaggerSpec)
	return nil
}
