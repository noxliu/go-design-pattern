package factoryMethod

type BmwFactory struct {
}

func (*BmwFactory) GetCar(carName string) Car {
	cartype := CarType{
		CarName:  carName,
		CarBrand: "bmw",
		Size:     10,
	}

	bmw := Bmw{
		CarType: cartype,
	}

	return bmw
}
