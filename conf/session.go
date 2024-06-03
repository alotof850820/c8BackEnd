package conf

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 配置和控制数据库操作的行为，根据具体的需求来调整和优化数据库操作。
func Session() {
	var _db *gorm.DB
	tx := _db.Session(&gorm.Session{
		PrepareStmt:              true,                                 //减少语法解析和执行的开销，并防止SQL注入攻击。
		SkipHooks:                true,                                 //跳过所有的钩子函数
		DisableNestedTransaction: true,                                 //禁用嵌套事务。
		Logger:                   logger.Default.LogMode(logger.Error), //记录錯誤信息级别及以上的日志
	},
	)

	tx.Debug()
}
