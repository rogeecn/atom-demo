package routes

import (
	"atom/modules/system/controller"

	"github.com/gin-gonic/gin"
	"github.com/rogeecn/atom/providers/http"
	"github.com/rogeecn/gen"
)

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
