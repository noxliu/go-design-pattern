package factoryMethod

type Car interface {
	Drive() string
}

type CarType struct {
	CarName  string
	CarBrand string
	Size     int
}
