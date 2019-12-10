package main

import . "go-design-pattern/adapter/adapterExample/adapter"

func main() {
	stringPrinter := Printer{}
	stringPrinter.PrinterForString("")
}
