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

func (c *CaptchaController) Show(ctx *gin.Context) (*captcha.CaptchaResponse, error) {
	return &captcha.CaptchaResponse{
		CaptchaId:     "123",
		PicPath:       "asdf",
		CaptchaLength: viper.GetUint("CAPTCHA_IMG_KEY_LONG"),
		OpenCaptcha:   viper.GetUint("CAPTCHA_IMG_OPEN"),
	}, nil
}
