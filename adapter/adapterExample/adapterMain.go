package main

import (
	. "go-design-pattern/adapter/adapterExample/adapter"
)

func main() {
	printer := Printer{}
	stringGenerator := GeneratorForString{}
	printer.PrinterForString(&stringGenerator)

	intGenerator := GeneratorForInt{}
	adapter := AdapterOfGeneratorForInt{}
	adapter.IntValue = intGenerator.IntGenerator()

	intArrayGenerator := GeneratorForIntArray{}
	adapterArray := AdapterOfGeneratorForArray{}
	adapterArray.IntArray = intArrayGenerator.IntArrayGenerator()

	printer.PrinterForString(&adapter)
	printer.PrinterForString(&adapterArray)

}
