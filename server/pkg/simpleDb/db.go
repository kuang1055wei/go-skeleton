package simpleDb

import (
	"database/sql"
	"fmt"
	"go-skeleton/pkg/config"
	"log"
	"os"
	"time"

	"go.uber.org/zap"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

//appMode常量
type AppMode string

const (
	DEBUG AppMode = "debug"
	PROD  AppMode = "prod"
)

type dbConfig struct {
	DbUser          string
	DbPassWord      string
	DbHost          string
	DbPort          string
	DbName          string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
	AppMode         AppMode //根据此变量选择日志模式
}

var err error

func dbDial(cfg *dbConfig) error {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DbUser,
		cfg.DbPassWord,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)

	//根据环境选择debug
	var newLogger logger.Interface
	if cfg.AppMode == DEBUG {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: 200 * time.Millisecond,
				LogLevel:      logger.Info, //log级别改此项即可
				Colorful:      true,
			},
		)
	} else {
		newLogger = logger.Default.LogMode(logger.Silent)
	}

	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：silent
		Logger: newLogger, //logger.Default.LogMode(logger.Silent)
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
			TablePrefix:   "",
		},
	})

	if err != nil {
		return err
	}
	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	//_ = db.AutoMigrate(&User{}, &Article{}, &Category{}, Profile{}, Comment{})

	//下面两种设置连接池的方法效果一样
	if sqlDB, err = db.DB(); err == nil {
		// SetMaxIdleCons 设置连接池中的最大闲置连接数。
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

		// SetMaxOpenCons 设置数据库的最大连接数量。
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

		// SetConnMaxLifetiment 设置连接的最大可复用时间。
		sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

		//最大空闲时间
		sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	} else {
		zap.L().Warn("mysql conns error:", zap.String("error", err.Error()))
	}
	//这里是多数据源的连接池
	//db.Use(
	//	dbresolver.Register(dbresolver.Config{}).
	//		SetConnMaxIdleTime(time.Hour).
	//		SetConnMaxLifetime(24 * time.Hour).
	//		SetMaxIdleConns(10).
	//		SetMaxOpenConns(105),
	//)

	return nil
}

func InitDb() error {
	cfg := &dbConfig{
		DbUser:          config.Conf.DbConfig.DbUser,
		DbPassWord:      config.Conf.DbConfig.DbPassWord,
		DbHost:          config.Conf.DbConfig.DbHost,
		DbPort:          config.Conf.DbConfig.DbPort,
		DbName:          config.Conf.DbConfig.DbName,
		MaxOpenConns:    10,
		MaxIdleConns:    100,
		ConnMaxIdleTime: 24 * time.Hour,
		ConnMaxLifetime: 24 * time.Hour,
		AppMode:         AppMode(config.Conf.AppMode),
	}

	return dbDial(cfg)
}

// 获取数据库链接
func DB() *gorm.DB {
	return db
}

// 关闭连接
func CloseDB() {
	if sqlDB == nil {
		return
	}
	if err := sqlDB.Close(); nil != err {
		log.Printf("Disconnect from database failed: %s", err.Error())
	}
}
