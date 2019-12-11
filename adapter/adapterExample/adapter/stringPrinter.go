package adapter

import "fmt"

type StringPrinterI interface {
	StringPrinter(string string)
}

type StringPrinter struct {
}

func (*StringPrinter) PrintString(string stringGeneratorI) {
	fmt.Println(string.stringGenerator())
}
