package main

import (
	. "go-design-pattern/strategy/strategyExample/strategy"
)

func main() {
	employeeZhangsan := BaseInfo{"995646878", "张三", 1000}
	employeeSalaryCalculator := SalaryCalculator{&Employee{}, employeeZhangsan}
	employeeSalaryCalculator.CalculateSalary()

	managerLisi := BaseInfo{"998877985", "李四", 1000}
	managerSalaryCalculator := SalaryCalculator{&Manager{}, managerLisi}
	managerSalaryCalculator.CalculateSalary()

	salesManagerWangwu := BaseInfo{"998811225", "王五", 1000}
	salesManagerSalaryCalculator := SalaryCalculator{&SalesManager{}, salesManagerWangwu}
	salesManagerSalaryCalculator.CalculateSalary()
}
