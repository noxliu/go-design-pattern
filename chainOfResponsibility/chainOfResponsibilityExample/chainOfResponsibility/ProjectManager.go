package chainOfResponsibility

import "fmt"

type ProjectManager struct {
}

func (*ProjectManager) HaveRight(money int) bool {
	return money < 200
}

func (*ProjectManager) HandleFeeRequest(name string, money int) bool {
	fmt.Printf("项目经理处理 %s 的 %d 报销请求\n", name, money)
	return true

}
