package adapter

import "fmt"

type Printer struct {
}

func (*Printer) PrinterForString(string string) {
	fmt.Println(string)
}
