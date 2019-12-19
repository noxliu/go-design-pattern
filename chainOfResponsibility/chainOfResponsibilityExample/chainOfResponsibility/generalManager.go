package chainOfResponsibility

import "fmt"

type GeneralManager struct {
}

func (*GeneralManager) HaveRight(money int) bool {
	return true
}

func (*GeneralManager) HandleFeeRequest(name string, money int) bool {
	fmt.Printf("总经理处理 %s 的 %d 报销请求\n", name, money)
	return true
}
