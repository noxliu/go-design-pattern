package main

import (
	. "go-design-pattern/factoryMethod/withoutFactory/example"
)

func main() {
	carTypeBenz := CarType{
		CarName: "benz",
		Size:    10,
	}

	benz := Benz{CarType: carTypeBenz}
	benz.Drive()

}
