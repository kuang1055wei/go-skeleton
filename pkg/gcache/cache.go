package gcache

import (
	"encoding/json"
	"gin-test/pkg/gredis"
	"sync"

	"github.com/go-redis/cache/v8"

	"github.com/go-redis/redis/v8"
)

var instance *cache.Cache
var once sync.Once

func GetCache() *cache.Cache {
	once.Do(func() {
		ring := redis.NewRing(&redis.RingOptions{
			Addrs: map[string]string{
				"cache_server": "",
			},
			NewClient: func(name string, opt *redis.Options) *redis.Client {
				//直接使用gredis的客户端，上面的Addrs不能省略，
				return gredis.Client
			},
		})
		instance = cache.New(&cache.Options{
			Redis: ring,
			//LocalCache: cache.NewTinyLFU(1000, time.Minute),//本地缓存，可以使用内存提速
			Marshal: func(i interface{}) ([]byte, error) {
				//使用官方json包
				j, err := json.Marshal(i)
				return j, err
			},
			Unmarshal: func(bytes []byte, i interface{}) error {
				err := json.Unmarshal(bytes, i)
				return err
			},
		})
	})
	return instance
}

//func Demo() {
//	//直接设置缓存
//	_ = GetCache().Set(&cache.Item{
//		Key:   "myKey2",
//		Value: 1,
//		TTL:   time.Minute * 5,
//	})
//
//	//只会加载一次，如果不存在，则调用do方法加载
//	_ = GetCache().Once(&cache.Item{
//		Key:   "myKey",
//		Value: nil,
//		TTL:   time.Minute * 10,
//		Do: func(item *cache.Item) (interface{}, error) {
//			res := map[string]string{
//				"a": "aaaa",
//				"b": "bbbb",
//			}
//			return res, nil
//		},
//	})
//
//	//获取
//	var val interface{}
//	GetCache().Get(context.TODO(), "myKey", &val)
//	fmt.Printf("%+v\n", val)
//}
