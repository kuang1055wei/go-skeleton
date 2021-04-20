package utils

//这种加载配置方式更优雅。推荐用此方式

import (
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

//全局配置
type Config struct {
	ServerConfig `ini:"server"`
	DbConfig     `ini:"database"`
	LogConfig    `ini:"log"`
	QiniuConfig  `ini:"qiniu"`
	RedisConfig  `ini:"redis"`
	AppConfig    `ini:"app"`
}

type AppConfig struct {
	JwtKey string `ini:"JwtKey"`

	PrefixUrl string `ini:"PrefixUrl"`

	RuntimeRootPath string `ini:"RuntimeRootPath"`

	ImageSavePath  string   `ini:"ImageSavePath"`
	ImageMaxSize   int      `ini:"ImageMaxSize"`
	ImageAllowExts []string `ini:"ImageAllowExts"`

	ExportSavePath string `ini:"ExportSavePath"`
	QrCodeSavePath string `ini:"QrCodeSavePath"`
	FontSavePath   string `ini:"FontSavePath"`
}

type ServerConfig struct {
	AppMode      string        `ini:"AppMode"`
	HttpPort     string        `ini:"HttpPort"`
	JwtKey       string        `ini:"JwtKey"`
	ReadTimeout  time.Duration `ini:"ReadTimeout"`
	WriteTimeout time.Duration `ini:"WriteTimeout"`
}

//数据库配置
type DbConfig struct {
	Db         string `ini:"Db"`
	DbHost     string `ini:"DbHost"`
	DbPort     string `ini:"DbPort"`
	DbUser     string `ini:"DbUser"`
	DbPassWord string `ini:"DbPassWord"`
	DbName     string `ini:"DbName"`
}

//日志配置
type LogConfig struct {
	Level      string `ini:"Level"`
	Filename   string `ini:"Filename"`
	MaxSize    int    `ini:"MaxSize"`
	MaxAge     int    `ini:"MaxAge"`
	MaxBackups int    `ini:"MaxBackups"`
}

type QiniuConfig struct {
	AccessKey  string `ini:"AccessKey"`
	SecretKey  string `ini:"SecretKey"`
	Bucket     string `ini:"Bucket"`
	QiniuSever string `ini:"QiniuSever"`
}

type RedisConfig struct {
	Host        string        `ini:"Host"`
	Password    string        `ini:"Password"`
	MaxIdle     int           `ini:"MaxIdle"`
	MaxActive   int           `ini:"MaxActive"`
	IdleTimeout time.Duration `ini:"IdleTimeout"`
}

var Conf = new(Config)

func InitConfig() {
	err := ini.MapTo(Conf, "config/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
}
