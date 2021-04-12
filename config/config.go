package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

//全局配置
type Config struct {
	*AppConfig
	*DbConfig
	*LogConfig
	*QiniuConfig
}

type AppConfig struct {
	AppMode  string
	HttpPort string
	JwtKey   string
}

//数据库配置
type DbConfig struct {
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
}

//日志配置
type LogConfig struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

type QiniuConfig struct {
	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
}

var Conf = new(Config)

func init() {
	//env := os.Getenv("env")
	//if env == "" {
	//
	//}
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadDbData(file)
	LoadServer(file)
	LoadQiniu(file)
	LoadLog(file)
}
func LoadServer(file *ini.File) {
	Conf.AppConfig = &AppConfig{
		AppMode:  file.Section("server").Key("AppMode").MustString("debug"),
		HttpPort: file.Section("server").Key("HttpPort").MustString(":3000"),
		JwtKey:   file.Section("server").Key("JwtKey").MustString("89js82js72"),
	}
}

func LoadDbData(file *ini.File) {
	Conf.DbConfig = &DbConfig{
		Db:         file.Section("database").Key("Db").MustString("debug"),
		DbHost:     file.Section("database").Key("DbHost").MustString("localhost"),
		DbPort:     file.Section("database").Key("DbPort").MustString("3306"),
		DbUser:     file.Section("database").Key("DbUser").MustString("ginblog"),
		DbPassWord: file.Section("database").Key("DbPassWord").MustString("admin123"),
		DbName:     file.Section("database").Key("DbName").MustString("ginblog"),
	}
}

func LoadQiniu(file *ini.File) {
	AccessKey := file.Section("qiniu").Key("AccessKey").String()
	SecretKey := file.Section("qiniu").Key("SecretKey").String()
	Bucket := file.Section("qiniu").Key("Bucket").String()
	QiniuSever := file.Section("qiniu").Key("QiniuSever").String()
	Conf.QiniuConfig = &QiniuConfig{
		AccessKey:  AccessKey,
		SecretKey:  SecretKey,
		Bucket:     Bucket,
		QiniuSever: QiniuSever,
	}
}

func LoadLog(file *ini.File) {
	MaxSize, _ := file.Section("log").Key("MaxSize").Int()
	MaxAge, _ := file.Section("log").Key("MaxAge").Int()
	MaxBackups, _ := file.Section("log").Key("MaxBackups").Int()
	Conf.LogConfig = &LogConfig{
		Level:      file.Section("log").Key("Level").String(),
		Filename:   file.Section("log").Key("Filename").String(),
		MaxSize:    MaxSize,
		MaxAge:     MaxAge,
		MaxBackups: MaxBackups,
	}

}
