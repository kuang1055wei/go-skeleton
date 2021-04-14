package article

import (
	"fmt"
	"gin-test/model"
	"io/ioutil"
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

func MyHttp(c *gin.Context) {
	//自定义头什么的
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://www.leiphone.com", nil)
	req.Header.Add("auth", "test")
	resp2, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
	}
	body2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println(body2)
	defer resp2.Body.Close()

	//普通的
	//resp, err := http.Get("http://www.leiphone.com")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%V", resp)
	c.JSON(http.StatusOK, gin.H{
		//"body":    string(body),
		"body2": string(body2),
		//"content": resp.ContentLength,
	})
}
