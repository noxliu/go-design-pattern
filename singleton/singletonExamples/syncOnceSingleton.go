package singletonExamples

import "sync"

type syncOnceSingleton struct{}

var syncOnceInstance *syncOnceSingleton
var once sync.Once

func GetSyncOnceInstance() *syncOnceSingleton {
	once.Do(func() {
		syncOnceInstance = &syncOnceSingleton{}
	})
	return syncOnceInstance
}
