package goframework_gorm_sqlite

import (
	"log/slog"

	"github.com/kordar/godb"
	"gorm.io/gorm"
)

var (
	sqlitepool = godb.NewDbPool()
	dbLogLevel = "info"
)

func GetSqliteDB(db string) *gorm.DB {
	return sqlitepool.Handle(db).(*gorm.DB)
}

func SetDbLogLevel(level string) {
	dbLogLevel = level
}

func gormConfig() *gorm.Config {
	mysqlConfig := gorm.Config{}
	mysqlConfig.Logger = newSlogGormLogger(dbLogLevel)
	return &mysqlConfig
}

// AddSqliteInstances 批量添加Sqlite句柄
func AddSqliteInstances(dbs map[string]string) {
	for db, dsn := range dbs {
		ins := NewGormSqliteConnIns(db, dsn, gormConfig())
		if ins == nil {
			continue
		}
		err := sqlitepool.Add(ins)
		if err != nil {
			slog.Warn("[godb-sqlite] 初始化Sqlite异常", "err", err)
		}
	}
}

// AddSqliteInstance 添加Sqlite句柄
func AddSqliteInstance(db string, dsn string) error {
	ins := NewGormSqliteConnIns(db, dsn, gormConfig())
	return sqlitepool.Add(ins)
}

// RemoveSqliteInstance 移除Sqlite句柄
func RemoveSqliteInstance(db string) {
	sqlitepool.Remove(db)
}

// HasSqliteInstance sqlite句柄是否存在
func HasSqliteInstance(db string) bool {
	return sqlitepool != nil && sqlitepool.Has(db)
}
