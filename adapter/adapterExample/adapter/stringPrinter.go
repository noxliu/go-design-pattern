package adapter

import "fmt"

type StringPrinterI interface {
	StringPrinter(string string)
}

type Printer struct {
}

func (*Printer) StringPrinter(string stringGeneratorI) {
	fmt.Println(string.stringGenerator())
}
