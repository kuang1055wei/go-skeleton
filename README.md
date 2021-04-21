
# gin-test
gin框架脚手架案例 日志zap    
数据库: [GORM文档](https://learnku.com/docs/gorm/v2)       
验证: [validations](https://github.com/go-playground/validator)  
配置:[viper](https://github.com/spf13/viper)
redis:[goredis](https://github.com/go-redis/redis)
优雅关机:[参考](https://www.liwenzhou.com/posts/Go/graceful_shutdown/)
jwt:[jwt-go](github.com/dgrijalva/jwt-go)

## 常用方法
**中间件**:middleware
**配置**:config/config.ini
**加密解密**: pkg/encrypt.go
**常用工具方法**: utils/utils.go
**验证翻译中文** : utils/trans.go

# 运行
1、go mod tidy     
2、go mod download     
3、air or go run main.go    
4、localhost:8000