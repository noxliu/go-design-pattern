package factoryMethod

type LenovoMouse struct {
	MouseType MouseType
}

func (b LenovoMouse) UseMouse() string {
	isWireLess := "有线"
	if b.MouseType.Wireless {
		isWireLess = "无线"
	}
	return b.MouseType.Brand + "," + isWireLess
}
