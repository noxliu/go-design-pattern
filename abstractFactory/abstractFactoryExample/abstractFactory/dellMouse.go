package factoryMethod

type DellMouse struct {
	MouseType MouseType
}

func (b DellMouse) UseMouse() string {
	isWireLess := "有线"
	if b.MouseType.Wireless {
		isWireLess = "无线"
	}
	return b.MouseType.Brand + "," + isWireLess
}
