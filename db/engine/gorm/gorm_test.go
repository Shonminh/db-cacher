// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package gorm

import (
	"testing"
	"time"
)

func TestGormEngine_Init(t *testing.T) {
	engine := New()
	err := engine.Init("mysql", "root@(127.0.0.1:3306)/?timeout=3000ms&readTimeout=1s&charset=utf8mb4", 20, 20, time.Hour)
	if err != nil {
		t.Error("Init gorm engine failed")
	}
	if err = engine.db.DB().Ping(); err != nil {
		t.Error("Ping gorm engine failed")
	}
}
