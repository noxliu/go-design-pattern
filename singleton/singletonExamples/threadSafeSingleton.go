package singletonExamples

import "sync"

type threadSafeSingleton struct{}

var threadSafeInstance *threadSafeSingleton
var mu1 sync.Mutex

func GetThreadSafeInstance() *threadSafeSingleton {
	if threadSafeInstance == nil {
		mu1.Lock()
		defer mu1.Unlock()
		if threadSafeInstance == nil {
			threadSafeInstance = &threadSafeSingleton{}
		}
	}
	return threadSafeInstance
}
