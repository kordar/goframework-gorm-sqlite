package goframework_gorm_sqlite

import (
	log "github.com/kordar/gologger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormSqliteConnIns struct {
	name string
	ins  *gorm.DB
}

func NewGormSqliteConnIns(name string, dsn string, config *gorm.Config) *GormSqliteConnIns {
	ins, err := gorm.Open(sqlite.Open(dsn), config)
	if err != nil {
		log.Errorf("[godb-sqlite] 初始化sqlite失败,dsn=%s,err=%v", dsn, err)
		return nil
	}
	return &GormSqliteConnIns{name: name, ins: ins}
}

func (c GormSqliteConnIns) GetName() string {
	return c.name
}

func (c GormSqliteConnIns) GetInstance() interface{} {
	return c.ins
}

func (c GormSqliteConnIns) Close() error {
	if s, err := c.ins.DB(); err == nil {
		return s.Close()
	} else {
		return err
	}
}
