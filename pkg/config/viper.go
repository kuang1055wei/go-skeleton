package config

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//tag:   https://github.com/mitchellh/mapstructure
//全局配置
type Config struct {
	ServerConfig `mapstructure:"server"`
	DbConfig     `mapstructure:"database"`
	LogConfig    `mapstructure:"log"`
	QiniuConfig  `mapstructure:"qiniu"`
	RedisConfig  `mapstructure:"redis"`
	AppConfig    `mapstructure:"app"`
	//Other map[string]interface{} `mapstructure:",remain"`
}

type AppConfig struct {
	JwtKey string `mapstructure:"JwtKey"`

	PrefixUrl string `mapstructure:"PrefixUrl"`

	RuntimeRootPath string `mapstructure:"RuntimeRootPath"`

	ImageSavePath  string   `mapstructure:"ImageSavePath"`
	ImageMaxSize   int      `mapstructure:"ImageMaxSize"`
	ImageAllowExts []string `mapstructure:"ImageAllowExts"`

	ExportSavePath string `mapstructure:"ExportSavePath"`
	QrCodeSavePath string `mapstructure:"QrCodeSavePath"`
	FontSavePath   string `mapstructure:"FontSavePath"`
}

type ServerConfig struct {
	AppMode      string        `mapstructure:"AppMode"`
	HttpPort     string        `mapstructure:"HttpPort"`
	ReadTimeout  time.Duration `mapstructure:"ReadTimeout"`
	WriteTimeout time.Duration `mapstructure:"WriteTimeout"`
}

//数据库配置
type DbConfig struct {
	Db         string `mapstructure:"Db"`
	DbHost     string `mapstructure:"DbHost"`
	DbPort     string `mapstructure:"DbPort"`
	DbUser     string `mapstructure:"DbUser"`
	DbPassWord string `mapstructure:"DbPassWord"`
	DbName     string `mapstructure:"DbName"`
}

//日志配置
type LogConfig struct {
	Level       string `mapstructure:"Level"`
	Filename    string `mapstructure:"Filename"`
	MaxSize     int    `mapstructure:"MaxSize"`
	MaxAge      int    `mapstructure:"MaxAge"`
	MaxBackups  int    `mapstructure:"MaxBackups"`
	LogSavePath string `mapstructure:"LogSavePath"`
	TimeFormat  string `mapstructure:"TimeFormat"`
}

type QiniuConfig struct {
	AccessKey  string `mapstructure:"AccessKey"`
	SecretKey  string `mapstructure:"SecretKey"`
	Bucket     string `mapstructure:"Bucket"`
	QiniuSever string `mapstructure:"QiniuSever"`
}

type RedisConfig struct {
	Host        string        `mapstructure:"Host"`
	Password    string        `mapstructure:"Password"`
	MaxIdle     int           `mapstructure:"MaxIdle"`
	MaxActive   int           `mapstructure:"MaxActive"`
	IdleTimeout time.Duration `mapstructure:"IdleTimeout"`
}

var Conf = new(Config)

func InitConfig() error {
	//env := pflag.String("env" , "dev" , "环境变量: dev | test | product")
	//pflag.Parse()
	//fmt.Printf("%+v\n",*env)
	//env := os.Getenv("env")
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	_ = viper.Unmarshal(&Conf)
	viper.WatchConfig() //监听配置文件改动
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info(fmt.Sprintf("配置文件修改成功:%s", in.Name))
		_ = viper.Unmarshal(Conf)
	})
	//配置中心示例
	//viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.json")
	//viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	//err := viper.ReadRemoteConfig()

	return nil
}
