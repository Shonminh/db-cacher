// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package redis

import (
	"time"

	"github.com/gin-contrib/cache/persistence"
)

var redisCacherPrefix = "redisCacherPrefixKey_"
var gRedisStore *persistence.RedisStore

type RedisCacher struct {
	key string
	// persistence.RedisStore
	writer func(in []interface{}) error
	reader func(key string) (interface{}, error)
	rs     *persistence.RedisStore
}

func New(uniqueId string, writer func(in []interface{}) error, reader func(key string) (interface{}, error)) *RedisCacher {
	c := RedisCacher{}
	c.key = redisCacherPrefix + uniqueId
	c.writer = writer
	c.reader = reader
	c.rs = gRedisStore
	return &c
}

// init redis function
func Init(host, password string, expireTime time.Duration) {
	redisStore := persistence.NewRedisCache(host, password, expireTime)
	gRedisStore = redisStore
}

func (c *RedisCacher) Write(in []interface{}) error {
	return c.writer(in)
}

func (c *RedisCacher) Read(key string) (interface{}, error) {
	return c.reader(key)
}
