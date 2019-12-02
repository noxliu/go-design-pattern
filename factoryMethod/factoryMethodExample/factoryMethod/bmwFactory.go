package factoryMethod

type BmwFactory struct {
}

func (*BmwFactory) GetCar(carName string) Car {
	cartype := CarType{
		carName: carName,
		size:    10,
	}

	bmw := Bmw{
		CarType: cartype,
	}

	return bmw
}
