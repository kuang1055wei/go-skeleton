package article

import (
	"encoding/json"
	"fmt"
	"gin-test/model"
	"gin-test/pkg/config"
	"gin-test/pkg/gredis"
	"gin-test/pkg/simpleDb"
	"gin-test/pkg/upload"
	"gin-test/services"
	"gin-test/utils"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/spf13/viper"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	var articleInfo model.Article
	articleMap := make(map[string]interface{})
	id, _ := strconv.Atoi(c.Query("id"))
	key := "article:" + strconv.Itoa(id)
	data, err := gredis.Client.Get(gredis.Ctx, key).Result()
	if data == "" || err != nil {
		art := model.Article{}
		articleInfo, err = art.GetArticleById(id)
		if articleInfo != (model.Article{}) {
			cacheValue, _ := json.Marshal(articleInfo)
			_ = gredis.Client.Set(gredis.Ctx, key, cacheValue, 60*60*time.Second).Err()
			_ = json.Unmarshal(cacheValue, &articleMap)
		}
	} else {
		_ = json.Unmarshal([]byte(data), &articleInfo)
		_ = json.Unmarshal([]byte(data), &articleMap)
	}
	//测试连接池
	//var wg sync.WaitGroup
	//for i := 0; i < 1000; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		res2, _ := gredis.RedisClient.Get(gredis.Ctx, "test2").Result()
	//		status := gredis.RedisClient.PoolStats()
	//		data1, _ := json.Marshal(status)
	//		fmt.Println(strconv.Itoa(i), res2, string(data1))
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
	//测试连接池end
	res, _ := gredis.Client.SetNX(gredis.Ctx, "test", "test value", time.Second*60).Result()
	res1, _ := gredis.Client.Set(gredis.Ctx, "test2", "test value2", time.Second*55).Result()
	res2, _ := gredis.Client.Get(gredis.Ctx, "test2").Result()
	redisMap := make(map[string]interface{})
	redisMap["res"] = res
	redisMap["res1"] = res1
	redisMap["res2"] = res2
	redisMap["articleCache"] = data
	redisMap["poolStatus"] = gredis.Client.PoolStats()

	c.JSON(http.StatusOK, gin.H{
		"uid":         c.GetInt("uid"),
		"username":    c.GetString("username"),
		"articleInfo": articleInfo,
		"articleMap":  articleMap,
		"redisRes":    redisMap,
	})
}

func GetArticleList(c *gin.Context) {
	//params := model.NewQueryParams(c)
	//params.LikeByReq("title").PageByReq().Desc("id")

	params := simpleDb.NewQueryParams(c)
	params.LikeByReq("title").PageByReq().Desc("id")
	articleList3, paging := services.ArticleService.FindPageByParams(params)

	//pageSize, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	//page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	//title := c.Query("title")
	//var condition = make(map[string]string)
	//if title != "" {
	//	condition["title"] = title
	//}
	//list1, total1 := model.GetArticleList(condition, pageSize, page)
	//list, total := model.GetArticleList2(condition, pageSize, page)
	c.JSON(http.StatusOK, gin.H{
		"articleList3": articleList3,
		"paging":       paging,

		//"list1":  list1,
		//"total1": total1,
		//"list":   list,
		//"total":  total,
	})
}

func SearchArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	title := c.Query("title")

	list, total, err := model.SearchArticle2(title, pageSize, page)
	fmt.Printf("%+v\n\n", err)
	c.JSON(http.StatusOK, gin.H{
		"list":  list,
		"total": total,
	})
}

//test json:
//{
//	"title": "我是标题",
//	"cid": 1,
//	"desc": "descdescdescdesc",
//	"content": "内容内容内容内容内容内容内容",
//	"img": "asdadadad",
//	"comment_count": 1,
//	"read_count": 1
//}
type ArticleForm struct {
	Id      int    `json:"id" form:"id" binding:"numeric"`
	Title   string `json:"title" form:"title" binding:"required"`
	Cid     uint64 `json:"cid" form:"cid" binding:"required,lt=10"`
	Desc    string `json:"desc" form:"desc" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
	Img     string `json:"img" form:"img" binding:"required"`
	//CommentCount int64  `json:"comment_count" form:"comment_count" binding:"required"`
	//ReadCount    int64  `json:"read_count" form:"read_count" binding:"required"`
}

func EditArticle(c *gin.Context) {
	var artForm ArticleForm

	if err := c.ShouldBind(&artForm); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"type": 1,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"type": 2,
			"errs": errs,
			"msg":  utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		})
		return
	}
	jsonStr, _ := json.Marshal(artForm)
	var art model.Article
	json.Unmarshal(jsonStr, &art)
	result, err := model.EditArticle(int(art.ID), &art)
	c.JSON(http.StatusOK, gin.H{
		"article": artForm,
		"artJson": string(jsonStr),
		"msg":     "success",
		"result":  result,
		"err":     err,
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

func demoFunc(i int) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Hello word %d\n", i)
}

//使用ants
//https://github.com/panjf2000/ants/blob/master/README_ZH.md
func MyChan2(c *gin.Context) {
	defer ants.Release()
	var wg sync.WaitGroup
	//syncCalculateSum := func() {
	//	demoFunc()
	//	wg.Done()
	//}
	pool, _ := ants.NewPool(10)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		//使用自定义池
		_ = pool.Submit(func() {
			demoFunc(i)
			wg.Done()
		})
		//使用默认池
		//_ = ants.Submit(func() {
		//	demoFunc()
		//	wg.Done()
		//})
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

}

//func UploadImg(c *gin.Context)  {
//	file, err := c.FormFile("file")
//	if err != nil {
//		c.JSON(http.StatusOK , gin.H{
//			"err":err.Error(),
//		})
//		return
//	}
//	// c.JSON(200, gin.H{"message": file.Header.Context})
//	c.SaveUploadedFile(file, file.Filename)
//	c.String(http.StatusOK, file.Filename)
//}

func UploadImg(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err1": err,
			"err":  err.Error(),
		})
		return
	}
	if fileHeader == nil {
		c.JSON(http.StatusOK, gin.H{
			"res": 1,
		})
		return
	}

	imageName := upload.GetImageName(fileHeader.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		c.JSON(http.StatusOK, gin.H{
			"res": 2,
		})
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res": 3,
		})
		return
	}

	if err := c.SaveUploadedFile(fileHeader, src); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res": 4,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})

}

func ViperTest(c *gin.Context) {
	opt := gredis.Client.Options()
	optStr, _ := json.Marshal(opt)
	fmt.Printf("%+v", opt)
	c.JSON(200, gin.H{
		"database":  viper.GetViper().Sub("database"),
		"appSize":   viper.GetViper().Get("app.PageSize"),
		"viperTest": viper.GetViper().Get("app.PageSize"),
		"allviper":  viper.GetViper().AllSettings(),
		"redis":     string(optStr),
		"config":    config.Conf,
	})
}
