package simpleFactory

func getCar(carType string) Car{
	if carType == "benz" {
		car := CarType{
			carName: carType,
			size:    10,
		}s

		return car
	} else {
		return nil
	}
}
