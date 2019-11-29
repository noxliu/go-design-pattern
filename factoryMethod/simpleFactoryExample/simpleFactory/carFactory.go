package simpleFactory

func getCar(carType string) Car {
	if carType == "benz" {
		benz := Benz{}
		return benz
	} else {
		return nil
	}
}
