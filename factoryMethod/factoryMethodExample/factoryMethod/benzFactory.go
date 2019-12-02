package factoryMethod

type BenzFactory struct {
}

func (*BenzFactory) GetCar(carName string) Car {
	cartype := CarType{
		CarName:  carName,
		CarBrand: "benz",
		Size:     10,
	}

	benz := Benz{
		CarType: cartype,
	}

	return benz
}
