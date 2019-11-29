package simpleFactory

func getCar(carType string) Car {
	if carType == "benz" {
		car := CarType{
			carName: carType,
			size:    10,
		}

		return car
	} else {
		return nil
	}
}
