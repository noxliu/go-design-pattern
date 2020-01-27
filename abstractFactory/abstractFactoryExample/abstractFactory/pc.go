package factoryMethod

type Pc interface {
	ShowPcInfo() string
}

type PcType struct {
	Brand      string
	ScreenSize int
	MemorySize int
}
