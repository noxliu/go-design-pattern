package strategy

type SalaryCalculator interface {
	CalculateSalary() int
}

type BaseInfo struct {
	userId     string
	Name       string
	BaseSalary int
}
