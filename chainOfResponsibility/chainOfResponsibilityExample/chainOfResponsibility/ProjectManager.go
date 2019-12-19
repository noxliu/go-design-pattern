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
	fmt.Printf("Project manager permit %s %d fee request\n", name, money)
	return true

}
