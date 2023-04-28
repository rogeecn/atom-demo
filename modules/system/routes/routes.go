package routes

import (
	"atom/http/modules/system/controller"

	"github.com/gin-gonic/gin"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/http"
	"github.com/rogeecn/atom/utils/opt"
	"github.com/rogeecn/gen"
)

func Provide(opts ...opt.Option) error {
	return container.Container.Provide(NewRoute, atom.GroupRoutes)
}

type Route struct {
	captcha *controller.CaptchaController
	engine  *gin.Engine
}

func NewRoute(captcha *controller.CaptchaController, svc http.Service) http.Route {
	engine := svc.GetEngine().(*gin.Engine)
	return &Route{captcha: captcha, engine: engine}
}

func (r *Route) Register() {
	r.engine.GET("/", gen.DataFunc(r.captcha.Show))
}
