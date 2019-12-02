package factoryMethod

type BenzFactory struct {
}

func (*BenzFactory) GetCar(carName string) Car {
	cartype := CarType{
		carName: carName,
		size:    10,
	}

	benz := Benz{
		CarType: cartype,
	}

	return benz

}
