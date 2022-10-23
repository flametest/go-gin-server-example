package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewMysqlDB(config *Config) (*gorm.DB, error) {
	var logLevel logger.LogLevel
	if config.Debug {
		logLevel = logger.Info
	} else {
		logLevel = logger.Warn
	}
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: false,
		},
		Logger:                                   logger.Default.LogMode(logLevel),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}
	return db, err
}
