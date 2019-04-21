// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package db

import (
	"time"
)

// Engine is interface to operate db
type Engine interface {
	Init(dbName, dsn string, maxIdleConns, maxOpenConns int, connMaxLifetime time.Duration) error
	Where(query interface{}, args ...interface{}) Engine
	Find(out interface{}, where ...interface{}) Engine
	Create(value interface{}) Engine
	Limit(limit interface{}) Engine
	Order(value interface{}, reorder ...bool) Engine
	Select(query interface{}, args ...interface{}) Engine
	ScanRows() (results []map[string]string, err error)
	Table(name string) Engine
	Error() error
	Delete(value interface{}, where ...interface{}) Engine
	Updates(values interface{}, ignoreProtectedAttrs ...bool) Engine
}
