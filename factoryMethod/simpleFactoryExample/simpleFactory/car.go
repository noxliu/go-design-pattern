package simpleFactory

//API is interface
type Car interface {
	Drive(name string) string
}

type CarType struct {
	carName string
	size    int
}
