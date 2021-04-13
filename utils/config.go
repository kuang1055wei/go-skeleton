package utils

//推荐用config2方式
import (
	"flag"
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	EnvMode string
	Config  *ini.File
)

func InitConfig() {
	baseDir, _ := os.Getwd()
	envParam := flag.String("env", "dev", "--env dev/test/prod")
	flag.Parse()
	EnvMode = *envParam
	var err error
	Config, err = ini.Load(baseDir + "/config/config.ini") //"../config/app.ini",
	if err != nil {
		fmt.Println(err)
		return
	}
	// 加入环境变量
	Config.ValueMapper = os.ExpandEnv
}

func GetConfig(key string) *ini.Key {
	parts := strings.Split(key, "::")
	section := parts[0]
	keyStr := parts[1]
	return Config.Section(section).Key(keyStr)
}
