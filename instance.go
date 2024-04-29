package goframework_redis

import (
	"github.com/kordar/gocfg"
	log "github.com/kordar/gologger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormSqliteConnIns struct {
	name string
	ins  *gorm.DB
}

func gormConfig() *gorm.Config {
	dbLogLevel := gocfg.GetSystemValue("gorm_log_level")
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

func NewGormSqliteConnIns(name string, config *gorm.Config) *GormSqliteConnIns {
	dsn := gocfg.GetSectionValue(name, "data")
	ins, err := gorm.Open(sqlite.Open(dsn), config)
	if err != nil {
		log.Errorf("[sqlite] 初始化sqlite失败,dsn=%s,err=%v", dsn, err)
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
