// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package config

import (
	"fmt"

	"db-cacher/db"
	"db-cacher/logger"
)

type ScanDBConfig struct {
	uniqueId      string
	scanTableName string
	columns       string
	duration      int
	primaryKey    string
	order         int
	limit         int
	whereOption   string
}

type Config struct {
	scanDbConfigMap map[string]*ScanDBConfig
	db              db.Engine
}

var gConfig = Config{scanDbConfigMap: make(map[string]*ScanDBConfig)}

func Register(db db.Engine, configs ...ScanDBConfig) {
	for _, config := range configs {
		gConfig.scanDbConfigMap[config.uniqueId] = &config
	}
	gConfig.db = db
}

func (scanConfig *ScanDBConfig) Worker(engine db.Engine) {

	scanRows, err := engine.Where(scanConfig.whereOption).Order(scanConfig.primaryKey).Limit(scanConfig.limit).ScanRows(true)
	if err != nil {
		logger.LogErrorf("ScanRows failed, error is %v", err)
	}
	fmt.Println(len(scanRows))

}

func (config *Config) GetEngine() db.Engine {
	return config.db
}
