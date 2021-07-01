package site

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {

	go func(name string) {
		fmt.Println(name)
	}("哈哈哈哈哈")

	c.String(http.StatusOK, "你好")

}
