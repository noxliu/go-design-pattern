package factoryMethod

type LenovoFactory struct {
}

func (*LenovoFactory) GetPc(brandName string, screenSize int, memorySize int) Pc {
	pctype := PcType{
		MemorySize: memorySize,
		Brand:      brandName,
		ScreenSize: screenSize,
	}

	lenovoPc := LenovoPc{
		PcType: pctype,
	}

	return lenovoPc
}

func (*LenovoFactory) GetMouse(brandName string, wireless bool) Mouse {
	mouseType := MouseType{
		Brand:    brandName,
		Wireless: wireless,
	}

	lenovoMouse := LenovoMouse{
		MouseType: mouseType,
	}
	return lenovoMouse
}
