package strategy

import "fmt"

type SalaryCalculator struct {
	Algorithm SalaryAlgorithm
	BaseInfo  BaseInfo
}

func (s *SalaryCalculator) CalculateSalary() {
	fmt.Println(s.Algorithm.Calculate(s.BaseInfo))
}
