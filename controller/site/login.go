package site

import (
	"fmt"
	"go-skeleton/pkg/auth"
	"go-skeleton/pkg/errors"
	"go-skeleton/services"
	"go-skeleton/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var form LoginForm
	if c.ShouldBind(&form) == nil {
		user, err := services.UserService.SignUp(form.User, form.Password, form.Password)
		if err != nil {
			c.JSON(http.StatusOK, utils.JsonError(err))
			return
		}
		c.JSON(http.StatusOK, utils.JsonData(user))
		return
	}
	c.JSON(200, utils.JsonErrorMsg("缺少参数"))
	return
}

func Login(c *gin.Context) {
	var form LoginForm
	if c.ShouldBind(&form) == nil {
		//测试账号 admin  123456
		userService := services.UserService
		user := userService.Take(map[string]interface{}{"username": form.User})

		if user != nil && utils.ValidatePassword(user.Password, form.Password) {
			token, err := auth.GenerateToken(int(user.ID))
			if err != nil {
				c.JSON(200, utils.JsonError(err))
				return
			}
			//refreshToken应该放入数据库或者缓存中
			refreshToken, err := services.UserTokenService.GenerateRefreshToken(int64(user.ID))
			if err != nil {
				c.JSON(200, utils.JsonError(err))
				return
			}
			c.JSON(200, utils.JsonData(gin.H{
				"refreshToken": refreshToken,
				"token":        token,
				"user":         user,
			}))
			return
		} else {
			c.JSON(200, utils.JsonErrorMsg("unauthorized"))
			return
		}
	} else {
		c.JSON(200, utils.JsonErrorMsg("缺少参数"))
		return
	}
}

//刷新token
func RefreshAccessToken(c *gin.Context) {
	refreshToken := c.PostForm("refreshToken")
	if refreshToken == "" {
		c.JSON(200, utils.JsonErrorMsg("refreshToken不能为空"))
	}
	userId, err := services.UserTokenService.GetUserIdByToken(refreshToken)
	if userId == 0 {
		c.JSON(200, utils.JsonCodeError(errors.RefreshTokenError))
		return
	}
	token, err := auth.GenerateToken(int(userId))
	if err != nil {
		c.JSON(200, utils.JsonError(err))
		return
	}
	c.JSON(200, utils.JsonData(gin.H{
		"token": token,
	}))
	return
}

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	// 上传文件至指定目录
	//dst := "/"
	//if err := c.SaveUploadedFile(file, dst); err != nil {
	//	c.String(http.StatusBadRequest, "上传失败")
	//}
	c.String(200, fmt.Sprintf("%s", file.Filename))
}

func multiUpload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files, _ := form.File["upload[]"]
	for _, file := range files {
		dst := "/"
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusBadRequest, "上传失败")
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
func SomeDataFromReader(c *gin.Context) {
	//c.DataFromReader()
}
