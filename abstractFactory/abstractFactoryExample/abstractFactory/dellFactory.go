package factoryMethod

type DellFactory struct {
}

func (*DellFactory) GetPc(brandName string, screenSize int, memorySize int) Pc {
	pctype := PcType{
		MemorySize: memorySize,
		Brand:      brandName,
		ScreenSize: screenSize,
	}

	dellPc := DellPc{
		PcType: pctype,
	}

	return dellPc
}

func (*DellFactory) GetMouse(brandName string, wireless bool) Mouse {
	mouseType := MouseType{
		Brand:    brandName,
		Wireless: wireless,
	}

	dellMouse := DellMouse{
		MouseType: mouseType,
	}
	return dellMouse
}
