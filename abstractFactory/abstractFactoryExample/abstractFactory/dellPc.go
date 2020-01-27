package factoryMethod

import "strconv"

type DellPc struct {
	PcType PcType
}

func (b DellPc) ShowPcInfo() string {
	return b.PcType.Brand + ", 屏幕尺寸:" + strconv.Itoa(b.PcType.ScreenSize) + ", 内存:" + strconv.Itoa(b.PcType.MemorySize)
}
