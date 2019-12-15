package strategy

import (
	"fmt"
	"strconv"
)

type SalaryCalculator struct {
	Algorithm SalaryAlgorithm
	BaseInfo  BaseInfo
}

func (s *SalaryCalculator) CalculateSalary() {
	fmt.Println("员工ID: " + s.BaseInfo.UserId + ", 员工姓名: " +
		s.BaseInfo.Name + ", 基本工资: " +
		strconv.Itoa(s.BaseInfo.BaseSalary) +
		", 计算后实发工资: " +
		strconv.Itoa(s.Algorithm.Calculate(s.BaseInfo)))
}
