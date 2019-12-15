package singletonExamples

type eagerSingleton struct{}

var eagerSingletonInstance *eagerSingleton = &eagerSingleton{}

func GetEagerSingletonInstance() *eagerSingleton {
	return eagerSingletonInstance
}
