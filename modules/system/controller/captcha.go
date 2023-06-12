package controller

import (
	"atom/http/modules/system/dto"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rogeecn/atom-addons/providers/captcha"
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

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/:id/:action [post]
func (c *CaptchaController) Go(ctx *gin.Context, id uint, action string, body *dto.UserInfoBody) (*captcha.CaptchaResponse, error) {
	return &captcha.CaptchaResponse{
		CaptchaId:     fmt.Sprintf("id: %d", id),
		PicPath:       fmt.Sprintf("action: %s, username: %s", action, body.Username),
		CaptchaLength: viper.GetUint("CAPTCHA_IMG_KEY_LONG"),
		OpenCaptcha:   viper.GetUint("CAPTCHA_IMG_OPEN"),
	}, nil
}
