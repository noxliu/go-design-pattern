package singletonExamples

type eagerSingleton struct{}

var eagerSingletonInstance *eagerSingleton

func GetInstance() *eagerSingleton {
	if eagerSingletonInstance == nil {
		eagerSingletonInstance = &eagerSingleton{}
	}
	return eagerSingletonInstance
}
