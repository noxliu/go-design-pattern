package main

import (
	. "go-design-pattern/adapter/adapterExample/adapter"
)

func main() {
	//只支持字符串类型的打印工具
	printer := StringPrinter{}

	stringGenerator := StringGenerator{}

	intGenerator := IntGenerator{}
	adapter := IntToStringAdapter{intGenerator.IntGenerator()}

	intArrayGenerator := GeneratorForIntArray{}
	adapterArray := ArrayToStringAdapter{intArrayGenerator.IntArrayGenerator()}

	printer.PrintString(&stringGenerator)
	printer.PrintString(&adapter)
	printer.PrintString(&adapterArray)

}
