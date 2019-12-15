package main

import (
	. "go-design-pattern/adapter/adapterExample/adapter"
)

func main() {
	//只支持字符串类型的打印工具
	printer := StringPrinter{}

	stringGenerator := StringGenerator{}

	intGenerator := IntGenerator{}
	//将int类型适配到string类型的适配器
	adapter := IntToStringAdapter{intGenerator.IntGenerator()}

	intArrayGenerator := GeneratorForIntArray{}
	//将int数组适配到string类型的适配器
	adapterArray := IntArrayToStringAdapter{intArrayGenerator.IntArrayGenerator()}

	printer.PrintString(&stringGenerator)
	printer.PrintString(&adapter)
	printer.PrintString(&adapterArray)
}
