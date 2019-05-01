// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package gorm

import (
	"fmt"
	"testing"
	"time"
)

var engine = new(GormEngine)
var tableName = "travis_test_db.test_tab"

type User struct {
	Id       uint64 `gorm:"type:BIGINT(20) UNSIGNED;PRIMARY_KEY;NOT NULL;"json:"id"`
	UserName string `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT '';"json:"user_name"`
	Sex      int8   `gorm:"type:TINIYINT(4);NOT NULL;DEFAULT 0;"json:"sex"`
	Address  string `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT '';"json:"address"`
	Email    string `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT '';"json:"email"`
	Ctime    uint32 `gorm:"type:INT(11) UNSIGNED;NOT NULL" json:"ctime"`
	Mtime    uint32 `gorm:"type:INT(11) UNSIGNED;NOT NULL" json:"mtime"`
}

func init() {
	err := engine.Init("mysql", "root@(127.0.0.1:3306)/?timeout=3000ms&readTimeout=1s&charset=utf8mb4", 20, 20, time.Hour)
	if err != nil {
		panic("Init gorm engine failed, error is: " + err.Error())
	}
	engine.db.LogMode(true)
}

func TestGormEngine_Init(t *testing.T) {

	if err := engine.db.DB().Ping(); err != nil {
		t.Error("Ping gorm engine failed")
	}
}

func TestGormEngine(t *testing.T) {

	// test create
	now := uint32(time.Now().Unix())
	user := User{
		Address:  "china",
		Email:    "demo@gmail.com",
		Sex:      0,
		UserName: "tome",
		Ctime:    now,
		Mtime:    now,
	}
	if d := engine.Table(tableName).Create(&user); d.Error() != nil {
		t.Error("create failed")
	}
	// test query
	var rows []*User
	if d := engine.Table(tableName).Where("user_name != '' ").Find(&rows); d.Error() != nil {
		t.Error("query failed")
	}
	fmt.Println(rows)

	// test update
	now = uint32(time.Now().Unix())
	if d := engine.Table(tableName).Where("user_name != ''").Updates(map[string]interface{}{
		"sex":     1,
		"address": "Hong Kong",
		"mtime":   now,
	}); d.Error() != nil {
		t.Error("update failed")
	}

	// test scan rows
	results, err := engine.Table(tableName).ScanRows(false)
	if err != nil {
		t.Errorf("scan rows failed, error is %v", err)
	}
	fmt.Println(results)

	// test delete
	engine.Table(tableName).Where("address = ?", "Hong Kong").Delete(User{})

}
