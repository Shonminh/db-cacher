// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"time"
)

var mysqlDb = "mysql"

type GormEngine struct {
	db *gorm.DB
}

func New() *GormEngine {
	var gormEngine = new(GormEngine)
	return gormEngine
}

func (engine *GormEngine) Init(dbName, dsn string, maxIdleConns, maxOpenConns int, connMaxLifetime time.Duration) error {
	if dbName == "" {
		dbName = mysqlDb
	}
	db, err := gorm.Open(dbName, dsn)
	if err != nil {
		return errors.WithMessage(err, "gorm.Open")
	}
	db.DB().SetConnMaxLifetime(connMaxLifetime)
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)
	if db.DB().Ping() != nil {
		return errors.WithMessage(err, "Ping failed")
	}
	engine.db = db
	return nil
}

