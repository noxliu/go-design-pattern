package strategy

type SalaryAlgorithm interface {
	Calculate(b BaseInfo) int
}

type BaseInfo struct {
	UserId     string
	Name       string
	BaseSalary int
}
