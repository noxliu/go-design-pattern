package factoryMethod

type Factory interface {
	GetPc() Pc
	GetMouse() Mouse
}
