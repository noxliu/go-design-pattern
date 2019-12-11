package singletonExamples

import "sync"

type lazySingletonWithLock struct{}

var lazySingletonWithLockInstance *lazySingletonWithLock
var mu sync.Mutex

func GetIns() *lazySingletonWithLock {
	mu.Lock()
	defer mu.Unlock()

	if lazySingletonWithLockInstance == nil {
		lazySingletonWithLockInstance = &lazySingletonWithLock{}
	}
	return lazySingletonWithLockInstance
}
