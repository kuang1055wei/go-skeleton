package api

import (
	"context"
	"go-skeleton/pkg/gcaptcha"
	"go-skeleton/pkg/gredis"
	"go-skeleton/pkg/jsonresult"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type CaptchaController struct {
}

//生成验证码
func (captcha *CaptchaController) GenerateCaptcha(c *gin.Context) {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      5,
		ShowLineOptions: base64Captcha.OptionShowSlimeLine,
		Length:          5,
		Source:          base64Captcha.TxtSimpleCharaters,
		//BgColor:         nil,
		//Fonts:           nil,
	}
	driver = driverString.ConvertFonts()

	store := gcaptcha.NewRedisStore(context.TODO(), gredis.GetRedis())
	captchaObj := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captchaObj.Generate()
	if err != nil {
		c.JSON(http.StatusOK, jsonresult.JsonError(err))
		return
	}
	c.JSON(http.StatusOK, jsonresult.JsonData(gin.H{
		"captchaId": id,
		"data":      b64s,
	}))
	return
}

type VerifyCaptcha struct {
	Id    string `form:"id"`
	Value string `form:"value"`
}

//验证验证码
func (captcha *CaptchaController) VerifyCaptcha(c *gin.Context) {
	var verify VerifyCaptcha
	err := c.ShouldBindQuery(&verify)
	if err != nil {
		c.JSON(http.StatusOK, jsonresult.JsonError(err))
		return
	}
	store := gcaptcha.NewRedisStore(context.TODO(), gredis.GetRedis())
	result := store.Verify(verify.Id, verify.Value, true)
	c.JSON(http.StatusOK, jsonresult.JsonData(gin.H{
		"verify_result": result,
	}))
}
