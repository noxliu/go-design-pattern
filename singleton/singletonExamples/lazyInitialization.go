package singletonExamples

import "sync"

type singleton struct{}

var ins *singleton
var mu sync.Mutex

func GetIns() *singleton {
	mu.Lock()
	defer mu.Unlock()
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}
