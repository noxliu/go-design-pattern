package strategy

type SalesManager struct {
}

func (*SalesManager) Calculate(b BaseInfo) int {
	return b.BaseSalary * 2
}
