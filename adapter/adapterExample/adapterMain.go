package main

import (
	. "go-design-pattern/adapter/adapterExample/adapter"
)

func main() {
	printer := Printer{}
	stringGenerator := StringGenerator{}
	printer.StringPrinter(&stringGenerator)

	intGenerator := GeneratorForInt{}
	adapter := AdapterOfGeneratorForInt{}
	adapter.IntValue = intGenerator.IntGenerator()

	intArrayGenerator := GeneratorForIntArray{}
	adapterArray := AdapterOfGeneratorForArray{}
	adapterArray.IntArray = intArrayGenerator.IntArrayGenerator()

	printer.StringPrinter(&adapter)
	printer.StringPrinter(&adapterArray)

}
