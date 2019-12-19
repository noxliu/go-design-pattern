package chainOfResponsibility

import "fmt"

type DepartmentManager struct {
}

func (*DepartmentManager) HaveRight(money int) bool {
	return money < 5000
}

func (*DepartmentManager) HandleFeeRequest(name string, money int) bool {
	fmt.Printf("Dep manager permit %s %d fee request\n", name, money)
	return true
}
