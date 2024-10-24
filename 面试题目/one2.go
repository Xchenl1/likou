package main

import (
	"fmt"
	"sync"
)

//func DoLoadAll() error
//func DoSyncData() error
//// 请实现以下接口，
//type Syncer interface {
//	Notify()
//	Startup()
//}
//用以支撑如下能力
//syncer := &XXXSyncer()
//1)调用 syncer.Startup() 的时候，会调用 DoLoadAll(), 并且保证一个进程内只会调用一次；
//2)当调用 syncer.Notify() 时，会调用 DoSyncData(), 并且保证不会并发调用。

type Syncer interface {
	Notify()
	Startup()
}

type XXXSyncer struct {
	Once1  sync.Once
	Mutex1 sync.Mutex
}

func (X *XXXSyncer) Notify() {
	X.Mutex1.Lock()
	defer X.Mutex1.Unlock()
	DoSyncData()
}

func (X *XXXSyncer) Startup() {
	X.Once1.Do(func() {
		err := DoLoadAll()
		if err == nil {
			fmt.Println("只实现一次")
		}
	})
}
func DoLoadAll() error {
	return nil
}

var num = 0

func DoSyncData() error {
	num++
	fmt.Println(num)
	return nil
}

func main() {
	a := sync.WaitGroup{}
	a.Add(1)
	XXXSyncer := &XXXSyncer{}
	XXXSyncer.Startup()
	go func() {
		XXXSyncer.Notify()
		a.Done()
	}()
	XXXSyncer.Notify()
	a.Wait()
}
