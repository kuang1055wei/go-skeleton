package simpleDb

import (
	"database/sql"
	"fmt"
	"gin-test/pkg/config"
	"log"
	"os"
	"time"

	"gorm.io/plugin/dbresolver"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

var err error

func InitDb() error {
	//dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	config.Conf.DbUser,
	//	config.Conf.DbPassWord,
	//	config.Conf.DbHost,
	//	config.Conf.DbPort,
	//	config.Conf.DbName,
	//)
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.DbConfig.DbUser,
		config.Conf.DbConfig.DbPassWord,
		config.Conf.DbConfig.DbHost,
		config.Conf.DbConfig.DbPort,
		config.Conf.DbConfig.DbName,
	)
	//根据环境选择debug
	var newLogger logger.Interface
	if config.Conf.ServerConfig.AppMode == "debug" {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: 200 * time.Millisecond,
				LogLevel:      logger.Silent,
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

	db.Use(
		dbresolver.Register(dbresolver.Config{}).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)

	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	//_ = db.AutoMigrate(&User{}, &Article{}, &Category{}, Profile{}, Comment{})

	//sqlDB, _ := db.DB()
	//// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	//sqlDB.SetMaxIdleConns(10)
	//
	//// SetMaxOpenCons 设置数据库的最大连接数量。
	//sqlDB.SetMaxOpenConns(100)
	//
	//// SetConnMaxLifetiment 设置连接的最大可复用时间。
	//sqlDB.SetConnMaxLifetime(10 * time.Second)

	return nil
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