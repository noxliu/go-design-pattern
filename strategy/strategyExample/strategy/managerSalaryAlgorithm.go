package strategy

type Manager struct {
}

func (*Manager) Calculate(b BaseInfo) int {
	return b.BaseSalary + 2000
}
