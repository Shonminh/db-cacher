// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package db

import (
	"time"
)

// Engine is interface to operate db
type Engine interface {
	Init(dbName, dsn string, maxIdleConns, maxOpenConns int, connMaxLifetime time.Duration) error
}
