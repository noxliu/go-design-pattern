package adapter

import "fmt"

type Printer struct {
}

func (*Printer) PrinterForString(string generatorI) {
	fmt.Println(string.StringGenerator())
}
