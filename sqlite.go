package goframework_redis

import (
	"github.com/kordar/godb"
	log "github.com/kordar/gologger"
	"gorm.io/gorm"
)

var sqlitepool *godb.DbConnPool

func GetSqliteDB(db string) *gorm.DB {
	return sqlitepool.Handle(db).(*gorm.DB)
}

// InitSqliteHandle 初始化Sqlite句柄
func InitSqliteHandle(dbs ...string) {
	sqlitepool = godb.GetDbPool()
	for _, db := range dbs {
		err := AddSqliteInstance(db)
		if err != nil {
			log.Warnf("初始化Sqlite异常，err=%v", err)
		}
	}
}

// AddSqliteInstance 添加Sqlite句柄
func AddSqliteInstance(db string) error {
	sqlitepool = godb.GetDbPool()
	ins := NewGormSqliteConnIns(db, gormConfig())
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
