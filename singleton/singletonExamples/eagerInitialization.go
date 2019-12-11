package singletonExamples

type singleton struct{}

var ins *singleton

func GetInstance() *singleton {
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}
