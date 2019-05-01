// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package cache

import "db-cacher/cache/redis"

type cacherPair struct {
	uniqueId string
	writer   func(in []interface{}) error
	reader   func(key string) (interface{}, error)
}

var gCachePairs = make(map[string]Cacher)

func Register(cacherType CacherType, pairs ...cacherPair) {
	for _, pair := range pairs {
		c := newCacherByType(cacherType, &pair)
		gCachePairs[pair.uniqueId] = c
	}
}

func newCacherByType(cacherType CacherType, pair *cacherPair) Cacher {
	switch cacherType {
	case RedisCacher:
		return redis.New(pair.uniqueId, pair.writer, pair.reader)

	}
	return redis.New(pair.uniqueId, pair.writer, pair.reader)
}

// get chacher by uniqueId
func GetCacher(uniqueId string) Cacher {
	return gCachePairs[uniqueId]
}
