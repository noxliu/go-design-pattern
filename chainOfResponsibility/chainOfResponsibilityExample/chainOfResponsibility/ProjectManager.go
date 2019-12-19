package chainOfResponsibility

import "fmt"

type ProjectManager struct {
}

func NewProjectManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &ProjectManager{},
	}
}

func (*ProjectManager) HaveRight(money int) bool {
	return money < 500
}

func (*ProjectManager) HandleFeeRequest(name string, money int) bool {
	fmt.Printf("总经理处理 %s 的 %d 报销请求\n", name, money)
	return true

}
