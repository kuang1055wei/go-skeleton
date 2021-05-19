package site

import (
	"fmt"
	"gin-test/services"
	"gin-test/utils"
	"gin-test/utils/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var form LoginForm
	if c.ShouldBind(&form) == nil {
		//真实情况应该这里是登陆验证逻辑
		if form.User == "user" && form.Password == "password" {
			uid := 1
			token, err := auth.GenerateToken(uid)
			if err != nil {
				c.JSON(200, utils.JsonError(err))
				return
			}
			//refreshToken应该放入数据库或者缓存中
			refreshToken, err := services.UserTokenService.GenerateRefreshToken(int64(uid))
			if err != nil {
				c.JSON(200, utils.JsonError(err))
				return
			}

			c.JSON(200, utils.JsonData(gin.H{
				"refreshToken": refreshToken,
				"token":        token,
			}))
			return
		} else {
			c.JSON(200, utils.JsonErrorMsg("unauthorized"))
			return
		}
	} else {
		c.JSON(200, utils.JsonErrorMsg("缺少参数"))
	}
}

//刷新token
func RefreshAccessToken(c *gin.Context) {
	refreshToken := c.PostForm("refreshToken")
	if refreshToken == "" {
		c.JSON(200, utils.JsonErrorMsg("refreshToken不能为空"))
	}
	userId, err := services.UserTokenService.GetRefreshTokenUserId(refreshToken)
	if userId == 0 {
		c.JSON(200, utils.JsonErrorMsg("refreshToken不合法"))
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
