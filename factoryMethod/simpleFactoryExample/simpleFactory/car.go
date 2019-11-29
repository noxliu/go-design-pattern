package simpleFactory

//API is interface
type Car interface {
	Drive() string
}

type CarType struct {
	carName string
	size    int
}
