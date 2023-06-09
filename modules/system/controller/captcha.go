package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rogeecn/atom/providers/captcha"
	"github.com/spf13/viper"
)

type CaptchaController struct {
}

func NewCaptchaController() *CaptchaController {
	return &CaptchaController{}
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (c *CaptchaController) Show(ctx *gin.Context) (*captcha.CaptchaResponse, error) {
	return &captcha.CaptchaResponse{
		CaptchaId:     "123",
		PicPath:       "asdf",
		CaptchaLength: viper.GetUint("CAPTCHA_IMG_KEY_LONG"),
		OpenCaptcha:   viper.GetUint("CAPTCHA_IMG_OPEN"),
	}, nil
}
