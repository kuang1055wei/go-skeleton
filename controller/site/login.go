package site

import (
	"fmt"
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
		if form.User == "user" && form.Password == "password" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	} else {
		c.JSON(401, gin.H{"status": "missing params"})
	}
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
