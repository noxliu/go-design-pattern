package singletonExamples

import "sync"

type lazySingletonWithLock struct{}

var lazySingletonWithLockInstance *lazySingletonWithLock
var mu sync.Mutex

func GetLazySingletonWithLockInstance() *lazySingletonWithLock {
	mu.Lock()
	defer mu.Unlock()

	if lazySingletonWithLockInstance == nil {
		lazySingletonWithLockInstance = &lazySingletonWithLock{}
	}
	return lazySingletonWithLockInstance
}
