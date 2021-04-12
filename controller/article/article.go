package article

import (
	"gin-test/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	art := model.Article{}
	articleInfo, _ := art.GetArticleById(id)
	c.JSON(http.StatusOK, articleInfo)
}

func GetArticleList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	title := c.Query("title")
	var condition = make(map[string]string)
	if title != "" {
		condition["title"] = title
	}
	list, total := model.GetArticleList(condition, pageSize, page)
	c.JSON(http.StatusOK, gin.H{
		"list":  list,
		"total": total,
	})
}
