
# go-skeleton
gin框架脚手架案例 日志zap    
数据库: [GORM文档](https://learnku.com/docs/gorm/v2)       
验证: [validations](https://github.com/go-playground/validator)  
配置:[viper](https://github.com/spf13/viper)  
redis:[goredis](https://github.com/go-redis/redis)  
缓存:[redis cache](https://github.com/go-redis/cache)  
优雅关机:[参考](https://www.liwenzhou.com/posts/Go/graceful_shutdown/)  
jwt:[jwt-go](https://github.com/dgrijalva/jwt-go)  
协程池:[ants](https://github.com/panjf2000/ants)  
分布式队列：[machinery](https://github.com/RichardKnop/machinery#retry-tasks)  
swagger选项: [swagger](https://swaggo.github.io/swaggo.io/declarative_comments_format/general_api_info.html)
swagger go文档：[swagger](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)

## 常用方法
**包文件**:pkg  
**中间件**:middleware  
**配置**:config/config.ini  
**加密解密**: pkg/encrypt.go  
**常用工具方法**: utils/utils.go  
**验证翻译中文** : utils/trans.go  

## 命令 
- swagger   
   1. cd server
   2. swag init
   3. http://localhost:8000/swagger/index.html


## 运行
1. cd server
2. go mod tidy
3. go mod download
4. air or go run main.go
5. localhost:8000
