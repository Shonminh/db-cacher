// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package gorm

import (
	"database/sql"
	"time"

	"db-cacher/db"
	"db-cacher/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
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
	openDb, err := gorm.Open(dbName, dsn)
	if err != nil {
		return errors.WithMessage(err, "gorm.Open")
	}
	openDb.DB().SetConnMaxLifetime(connMaxLifetime)
	openDb.DB().SetMaxIdleConns(maxIdleConns)
	openDb.DB().SetMaxOpenConns(maxOpenConns)
	if openDb.DB().Ping() != nil {
		return errors.WithMessage(err, "Ping failed")
	}
	engine.db = openDb
	return nil
}

func (engine *GormEngine) Table(name string) db.Engine {
	return &GormEngine{db: engine.db.Table(name)}
}

func (engine *GormEngine) Where(query interface{}, args ...interface{}) db.Engine {
	return &GormEngine{db: engine.db.Where(query, args...)}
}

func (engine *GormEngine) Find(out interface{}, where ...interface{}) db.Engine {
	return &GormEngine{db: engine.db.Find(out, where...)}

}

func (engine *GormEngine) Create(value interface{}) db.Engine {
	return &GormEngine{db: engine.db.Create(value)}

}

func (engine *GormEngine) Limit(limit interface{}) db.Engine {
	return &GormEngine{db: engine.db.Limit(limit)}

}

func (engine *GormEngine) Order(value interface{}, reorder ...bool) db.Engine {
	return &GormEngine{db: engine.db.Order(value, reorder...)}

}

func (engine *GormEngine) Select(query interface{}, args ...interface{}) db.Engine {
	return &GormEngine{db: engine.db.Select(query, args...)}
}

func (engine *GormEngine) Updates(values interface{}, ignoreProtectedAttrs ...bool) db.Engine {
	return &GormEngine{db: engine.db.Updates(values, ignoreProtectedAttrs...)}
}

func (engine *GormEngine) Delete(value interface{}, where ...interface{}) db.Engine {
	return &GormEngine{db: engine.db.Delete(value, where...)}
}


func (engine *GormEngine) ScanRows(ignoreError bool) (results []map[string]string, err error) {
	rows, err := engine.db.Rows()
	if err != nil {
		return nil, errors.WithMessage(err, "gorm.Rows")
	}
	defer rows.Close()
	results = make([]map[string]string, 0, 1024)
	columns, err := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if  err != nil && ignoreError {
			logger.LogWarnf("gorm.Scan failed, error is %v", err.Error())
			continue
		}
		if err != nil {
			return results, errors.WithMessage(err, "gorm.Scan")
		}
		result := make(map[string]string)
		for i, col := range values {
			var value string
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			result[columns[i]] = value
		}
		results = append(results, result)
	}
	return results, nil
}

func (engine *GormEngine) Error() error {
	return engine.db.Error
}
