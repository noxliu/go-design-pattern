package strategy

type Employ struct {
	BaseInfo         BaseInfo
	SalaryCalculator SalaryCalculator
}

func CalculateSalary() int {
	return 100
}
