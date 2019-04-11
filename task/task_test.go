// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package task

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimeTask(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(2)
	task1 := NewTimeTask(time.Second*4, func() error {
		fmt.Println("task1 job")
		return nil
	}, false, &group)
	task2 := NewTimeTask(time.Second*4, func() error {
		fmt.Println("task2 job")
		return nil
	}, true, &group)
	task1.Start()
	task2.Start()
	group.Wait()
}
