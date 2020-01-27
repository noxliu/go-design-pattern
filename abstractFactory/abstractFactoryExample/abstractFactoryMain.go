package main

import (
	"fmt"
	factoryMethod "go-design-pattern/abstractFactory/abstractFactoryExample/abstractFactory"
)

func main() {
	dellFactory := factoryMethod.DellFactory{}
	lenovoFactory := factoryMethod.LenovoFactory{}

	dellMouse := dellFactory.GetMouse("dell 101", false)
	dellPc := dellFactory.GetPc("dell xps", 14, 16)

	lenovoMouse := lenovoFactory.GetMouse("lenovo 455", true)
	lenovoPc := lenovoFactory.GetPc("lenovo yoga", 11, 8)

	fmt.Println(dellMouse.UseMouse())
	fmt.Println(dellPc.ShowPcInfo())
	fmt.Println(lenovoMouse.UseMouse())
	fmt.Println(lenovoPc.ShowPcInfo())
}
