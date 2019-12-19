package chainOfResponsibility

import "fmt"

type GeneralManager struct {
}

func (*GeneralManager) HaveRight(money int) bool {
	return true
}

func (*GeneralManager) HandleFeeRequest(name string, money int) bool {
	fmt.Printf("General manager permit %s %d fee request\n", name, money)
	return true
}
