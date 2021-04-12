package test

import (
	"fmt"
	"math/rand"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// 记录日志并使用zap.Xxx(key, val)记录相关字段
	var (
		age = 18
	)
	var name string = "hello kw 你好啊"
	zap.L().Debug("this is hello func", zap.String("user", name), zap.Int("age", age))
	// var arr [3]int
	// var result = [...]int{1, 2, 3, 4, 100}
	// var a = [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	c.String(http.StatusOK, fmt.Sprint(rand.Intn(10000)))
	c.String(http.StatusOK, name)
}

func HelloWord(c *gin.Context) {
	c.String(http.StatusOK, "hello word")
}

type user struct {
	ID   int
	Name string
	Age  int
}

func TestJson(c *gin.Context) {
	var a []string
	var s = []string{}
	var b = []int{1, 3, 4, 5}
	fmt.Println(a, b, s)

	//map[KeyType]ValueType
	var m map[string]string
	var n = map[string]string{"a": "b"}
	var mn = make(map[int]string, 8)
	fmt.Println(m, n)
	mn[1] = "张三"
	fmt.Println(mn, len(mn))

	v, isExists := mn[1] //判断map中键是否存在
	fmt.Println(v, isExists)

	var sliceMap = make([]map[string]string, 10) //元素为map类型的切片
	var mapSlice = make(map[string][]string, 10) //值为切片类型的map
	mapSlice["my"] = []string{"hhhh", "klhkjh"}
	fmt.Println(sliceMap, mapSlice)

	// var arr = [...]int{1, 2, 3}
	// var data = gin.H{"status": 1, "message": "", "data": arr}
	allUsers := []user{{ID: 123, Name: "张三", Age: 20}, {ID: 456, Name: "李四", Age: 25}}
	c.JSON(http.StatusOK, allUsers)
}
