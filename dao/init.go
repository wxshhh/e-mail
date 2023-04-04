package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var _db *gorm.DB

func DataBase(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	// 设置连接
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256, // string类型字段的默认长度
		SkipInitializeWithVersion: true,
		DisableDatetimePrecision:  true, // 禁止datetime精度，mysql5.6之前不支持
		DontSupportRenameIndex:    true, // 不支持重命名索引，mysql5.7之前不支持
		DontSupportRenameColumn:   true, // 不支持重命名列，mysql8之前不支持
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}, // 命名策略，单数化，不加s
		Logger: ormLogger,
	})
	if err != nil {
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池
	sqlDB.SetMaxOpenConns(100) // 打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	// 主从配值
	_ = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(connWrite)},                      // 写操作
		Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, // 读操作（用两个数据库）
		Policy:   dbresolver.RandomPolicy{},                                    // 策略：负载均衡
	}))

	// 数据库迁移
	migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
