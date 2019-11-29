package example

type Car interface {
	Drive() string
}

type CarType struct {
	CarName string
	Size    int
}
