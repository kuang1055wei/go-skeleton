package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var data = map[string]string{
		"title":   "我是主页",
		"content": "大家好",
	}

	c.HTML(http.StatusOK, "site/index.html", data)
	//c.String(http.StatusOK, "index")
}
