package factoryMethod

type Mouse interface {
	UseMouse() string
}

type MouseType struct {
	Brand    string
	Wireless bool
}
