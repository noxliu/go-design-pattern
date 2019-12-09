package main

import (
	. "go-design-pattern/strategy/strategyExample/strategy"
)

func main() {
	employeeZhangsan := BaseInfo{"998876878", "张三", 1000}
	sal := SalaryCalculator{&Employee{}, employeeZhangsan}
	sal.CalculateSalary()

}

func test(userType string) {
	if userType == "A" {
		//..
	} else if userType == "B" {
		//..
	} else if userType == "C" {
		//..
	} else if userType == "D" {
		//..
	} else {
		//..
	}
}
