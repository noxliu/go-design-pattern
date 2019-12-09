package main

import (
	. "go-design-pattern/strategy/strategyExample/strategy"
)

func main() {
	employeeZhangsan := BaseInfo{
		UserId:     "998876878",
		Name:       "张三",
		BaseSalary: 1000,
	}
	employeeCalculator := Employee{}

	sal := SalaryCalculator{
		Algorithm: &employeeCalculator,
		BaseInfo:  employeeZhangsan,
	}
	sal.CalculateSalary()

}
