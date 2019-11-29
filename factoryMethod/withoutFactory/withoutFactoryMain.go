package main

import (
	. "go-design-pattern/factoryMethod/withoutFactory/example"
)

func main() {
	carType := CarType{
		CarName: "benz",
		Size:    10,
	}

	benz := Benz{CarType: carType}
	benz.Drive()
}
