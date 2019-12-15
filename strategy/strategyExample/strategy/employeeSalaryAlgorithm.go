package strategy

type Employee struct {
}

func (*Employee) Calculate(b BaseInfo) int {
	return b.BaseSalary + 1000
}
