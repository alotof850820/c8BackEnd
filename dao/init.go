package dao

import (
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

// 初始化数据库
func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN:                       connRead,
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度,MySQL 5.6 之前的版本中不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的版本中不支持
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的版本中不支持
			SkipInitializeWithVersion: false, // 根据版本自动配置
		},
	), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		PrepareStmt: true, // 不支持嵌套事務
	})
	if err != nil {
		return
	}
	// setPool
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                 // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Second * 30) // 设置连接可复用的最大时间
	_db = db

	// 主從配置 dbresolver在多个数据库之间进行读写分离和负载均衡。
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(connWrite)},                       // 寫操作
		Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connWrite)}, // 讀操作
		Policy:   dbresolver.RandomPolicy{},
	}))

	// 自動建表
	Migration()
}

// *gin.Context是HTTP请求的上下文信息
// 基于现有的数据库客户端实例，创建一个新的数据库客户端实例，并使用传入的上下文信息。
// 这样可以方便地在不同的上下文环境中使用数据库客户端实例。
func NewDBClient(ctx *context.Context) *gorm.DB {
	db := _db
	return db.WithContext(*ctx)
}
