package routes

import (
	"atom/http/modules/system/controller"

	"github.com/gin-gonic/gin"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom-addons/providers/http"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	newRoutes := func(
		user *controller.UserController,
		svc http.Service,
	) http.Route {
		engine := svc.GetEngine().(*gin.Engine)

		routeUserController(engine.Group("/user"), user)

		return nil
	}
	return container.Container.Provide(newRoutes, atom.GroupRoutes)
}
