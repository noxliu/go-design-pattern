package main

import (
	"fmt"
	. "go-design-pattern/factoryMethod/withoutFactory/example"
)

func main() {
	carTypeBenz := CarType{
		CarName: "benz",
		Size:    10,
	}

	benz := Benz{CarType: carTypeBenz}
	fmt.Println(benz.Drive())

	carTypeBmw := CarType{
		CarName: "bmw",
		Size:    9,
	}

	bmw := Benz{CarType: carTypeBmw}
	fmt.Println(bmw.Drive())
}
