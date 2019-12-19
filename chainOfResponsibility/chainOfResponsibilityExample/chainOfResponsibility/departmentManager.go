package chainOfResponsibility

import "fmt"

type DepartmentManager struct {
}

func (*DepartmentManager) HaveRight(money int) bool {
	return money < 500
}

func (*DepartmentManager) HandleFeeRequest(name string, money int) bool {
	fmt.Printf("主管经理处理 %s 的 %d 报销请求\n", name, money)
	return true
}
