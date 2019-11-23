package main

import (
	"fmt"
	"strings"
)

type Printer interface {
	PrintWord(word string)
}

type PrintUpperCase struct {
}

func (*PrintUpperCase) PrintWord(word string) {
	fmt.Println(strings.ToUpper(word))
}

type PrintLowerCase struct {
}

func (*PrintLowerCase) PrintWord(word string) {
	fmt.Println(strings.ToLower(word))
}

type StateContext struct {
	state int //如果是0则是lower，如果是1则是upper
}

func (c *StateContext) PrintWord(word string) {
	upperPrinter := PrintUpperCase{}
	lowerPrinter := PrintLowerCase{}
	if c.state == 0 {
		lowerPrinter.PrintWord(word)
		c.state = 1
	} else {
		upperPrinter.PrintWord(word)
		c.state = 0
	}

}

type State interface {
	WriteName(name string)
}

func main() {
	stateContext := StateContext{}
	//调用Context的打印方法，调用内部切换状态
	stateContext.PrintWord("TestWord")
	stateContext.PrintWord("TestWord")
	stateContext.PrintWord("TestWord")
	stateContext.PrintWord("TestWord")
	stateContext.PrintWord("TestWord")
}
