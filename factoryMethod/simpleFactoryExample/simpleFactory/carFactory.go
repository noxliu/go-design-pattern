package simpleFactory

func GetCar(carType string) Car {
	carType1 := CarType{
		carName: carType,
		size:    10,
	}
	if carType == "benz" {
		benz := Benz{CarType: carType1}
		return benz
	} else if carType == "bmw" {
		bmw := Bmw{CarType: carType1}
		return bmw
	} else {
		return nil
	}
}
