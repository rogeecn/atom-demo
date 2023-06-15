package routes

import (
	"atom/http/modules/system/controller"

	"github.com/gin-gonic/gin"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom-addons/providers/http"
	"github.com/rogeecn/atom-addons/providers/log"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	return container.Container.Provide(newRoutes, atom.GroupRoutes)
}

func newRoutes(
	userController *controller.UserController,
	svc http.Service,
) http.Route {
	engine := svc.GetEngine().(*gin.Engine)
	group := engine.Group("users")
	log.Info("register route group: %s", group)

	routeUserController(group, userController)

	return nil
}
