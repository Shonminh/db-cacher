// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package task

import (
	"db-cacher/logger"
	"sync"
	"time"
)

type TimeTask struct {
	du       time.Duration
	isClose  bool
	lazyLoad bool
	job      func() error
	rMutex   sync.RWMutex
	wg       *sync.WaitGroup
}

func NewTimeTask(duration time.Duration, job func() error, lazyLoad bool, group *sync.WaitGroup) *TimeTask {
	return &TimeTask{
		du:       duration,
		job:      job,
		isClose:  false,
		lazyLoad: lazyLoad,
		wg:       group,
	}
}

func (t *TimeTask) Start() {
	go t.run()
}

func (t *TimeTask) run() {
	for false == t.checkIsClose() {
		ticker := time.NewTicker(t.du)

		if t.lazyLoad {
			<-ticker.C
		}
		if err := t.job(); err != nil {
			logger.LogErrorf("[Start]err is:%v", err.Error())
		}
		if t.lazyLoad == false {
			<-ticker.C
		}
		ticker.Stop()
	}
	t.wg.Done()
}

func (t *TimeTask) checkIsClose() bool {
	t.rMutex.RLock()
	defer t.rMutex.RUnlock()
	return t.isClose
}

func (t *TimeTask) Close() {
	t.rMutex.Lock()
	defer t.rMutex.Unlock()
	t.isClose = true
}
