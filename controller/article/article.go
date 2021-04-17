package article

import (
	"encoding/json"
	"fmt"
	"gin-test/model"
	"gin-test/pkg/gredis"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	var articleInfo model.Article
	id, _ := strconv.Atoi(c.Query("id"))
	key := "article:" + strconv.Itoa(id)
	data, err := gredis.RedisClient.Get(gredis.Ctx, key).Result()
	fmt.Println(data, "66666666666")
	if data == "" || err != nil {
		art := model.Article{}
		articleInfo, _ = art.GetArticleById(id)
		cacheValue, _ := json.Marshal(articleInfo)
		errSet := gredis.RedisClient.Set(gredis.Ctx, key, cacheValue, 60*60*time.Second).Err()
		fmt.Println(errSet, "111111111111")
	} else {
		_ = json.Unmarshal([]byte(data), &articleInfo)
		fmt.Println("get cache", "22222222222222")
	}

	reply, _ := gredis.RedisClient.SetNX(gredis.Ctx, "test", "test value", time.Second*60).Result()
	fmt.Println(reply, "33333333333333333")

	res, err := gredis.RedisClient.SetNX(gredis.Ctx, "test2", "test value2", time.Second*55).Result()
	fmt.Println(res, err, "44444444444444")
	res2, err2 := gredis.RedisClient.Get(gredis.Ctx, "test2").Result()
	fmt.Println(res2, err2, "55555555555555555")

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

//测试协程拉取
func MyChan(c *gin.Context) {
	id1, _ := strconv.Atoi(c.DefaultQuery("id1", "3"))
	id2, _ := strconv.Atoi(c.DefaultQuery("id2", "4"))
	art1Chan := getArt(id1)
	art2Chan := getArt(id2)

	art1 := <-art1Chan
	art2 := <-art2Chan
	data := make(map[string]interface{})
	data["art1"] = art1
	data["art2"] = art2
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
	//ctx := context.Background()
	//context.WithCancel(ctx)
	//context.WithTimeout(ctx, time.Microsecond)
	//context.WithDeadline(ctx, time.Now().Add(20))
	//context.WithValue(ctx, "anc", "aaaa")
}

func getArt(id int) <-chan model.Article {
	artChan := make(chan model.Article)
	go func(id int) {

		artModel := model.Article{}
		art, _ := artModel.GetArticleById(id)
		artChan <- art
	}(id)
	return artChan
}
