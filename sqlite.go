package goframework_gorm_sqlite

import (
	"github.com/kordar/godb"
	log "github.com/kordar/gologger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	if dbLogLevel == "error" {
		mysqlConfig.Logger = logger.Default.LogMode(logger.Error)
	}
	if dbLogLevel == "warn" {
		mysqlConfig.Logger = logger.Default.LogMode(logger.Warn)
	}
	if dbLogLevel == "info" {
		mysqlConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	return &mysqlConfig
}

// InitSqliteHandle 初始化Sqlite句柄
func InitSqliteHandle(dbs map[string]string) {
	for db, dsn := range dbs {
		ins := NewGormSqliteConnIns(db, dsn, gormConfig())
		if ins == nil {
			continue
		}
		err := sqlitepool.Add(ins)
		if err != nil {
			log.Warnf("初始化Sqlite异常，err=%v", err)
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
