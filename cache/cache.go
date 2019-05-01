// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package cache

import (
	// "time"
	//
	// "github.com/gin-contrib/cache/persistence"
	// "github.com/gin-contrib/cache/persistence/"
	// "github.com/gin-contrib/cache/persistence"
)

type Cacher interface {
	// persistence.CacheStore
	Write(in []interface{}) error
	Read(key string) (interface{}, error)
}

type CacherType int

const (
	MemoryCacher CacherType = iota
	RedisCacher
)
