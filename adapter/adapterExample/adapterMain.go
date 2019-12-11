package main

import (
	. "go-design-pattern/adapter/adapterExample/adapter"
)

func main() {
	printer := StringPrinter{}
	stringGenerator := StringGenerator{}
	printer.PrintString(&stringGenerator)

	intGenerator := GeneratorForInt{}
	adapter := AdapterOfGeneratorForInt{}
	adapter.IntValue = intGenerator.IntGenerator()

	intArrayGenerator := GeneratorForIntArray{}
	adapterArray := AdapterOfGeneratorForArray{}
	adapterArray.IntArray = intArrayGenerator.IntArrayGenerator()

	printer.PrintString(&adapter)
	printer.PrintString(&adapterArray)

}
