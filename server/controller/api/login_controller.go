package api

import (
	"go-skeleton/pkg/auth"
	"go-skeleton/pkg/errors"
	"go-skeleton/pkg/jsonresult"
	"go-skeleton/services"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (l *LoginController) Login(c *gin.Context) {
	var form LoginForm
	if c.ShouldBind(&form) == nil {
		userService := services.UserService
		user := userService.Take(map[string]interface{}{"username": form.User})

		//真实情况应该这里是登陆验证逻辑
		if form.User == "user" && form.Password == "password" {
			uid := 1
			token, err := auth.GenerateToken(uid)
			if err != nil {
				c.JSON(200, jsonresult.JsonError(err))
				return
			}
			//refreshToken应该放入数据库或者缓存中
			refreshToken, err := services.UserTokenService.GenerateRefreshToken(int64(uid))
			if err != nil {
				c.JSON(200, jsonresult.JsonError(err))
				return
			}
			c.JSON(200, jsonresult.JsonData(gin.H{
				"refreshToken": refreshToken,
				"token":        token,
				"user":         user,
			}))
			return
		} else {
			c.JSON(200, jsonresult.JsonErrorMsg("unauthorized"))
			return
		}
	} else {
		c.JSON(200, jsonresult.JsonErrorMsg("缺少参数"))
	}
}

//刷新token
func (l *LoginController) RefreshAccessToken(c *gin.Context) {
	refreshToken := c.PostForm("refreshToken")
	if refreshToken == "" {
		c.JSON(200, jsonresult.JsonErrorMsg("refreshToken不能为空"))
	}
	userId, err := services.UserTokenService.GetUserIdByToken(refreshToken)
	if userId == 0 {
		c.JSON(200, jsonresult.JsonCodeError(errors.RefreshTokenError))
		return
	}
	token, err := auth.GenerateToken(int(userId))
	if err != nil {
		c.JSON(200, jsonresult.JsonError(err))
		return
	}
	c.JSON(200, jsonresult.JsonData(gin.H{
		"token": token,
	}))
	return
}
