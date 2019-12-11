package singletonExamples

type lazySingleton struct{}

var lazySingletonInstance *lazySingleton

func GetLazySingletonInstance() *lazySingleton {
	if lazySingletonInstance == nil {
		lazySingletonInstance = &lazySingleton{}
	}
	return lazySingletonInstance
}
