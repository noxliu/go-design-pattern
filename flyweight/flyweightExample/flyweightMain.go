package main

import (
	"fmt"
	. "go-design-pattern/flyweight/flyweightExample/flyweight"
)

func main() {
	fmt.Println(GetFLyWeight("a"))
	fmt.Println(GetFLyWeight("a"))
	fmt.Println(GetFLyWeight("c"))
	fmt.Println(GetFLyWeight("b"))
	fmt.Println(GetFLyWeight("a"))
	fmt.Println(GetFLyWeight("a"))
	fmt.Println(GetFLyWeight("b"))
	fmt.Println(GetFLyWeight("d"))
	fmt.Println(len(F))
}
