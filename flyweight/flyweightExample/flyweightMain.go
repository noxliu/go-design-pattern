package main

import (
	"fmt"
	. "go-design-pattern/flyweight/flyweightExample/flyweight"
)

func main() {
	fmt.Println("--")
	fmt.Println(GetFLyWeight("a11"))
	fmt.Println(GetFLyWeight("a12"))
	fmt.Println(GetFLyWeight("a13"))
	fmt.Println(GetFLyWeight("a11"))
	fmt.Println(GetFLyWeight("a12"))
	fmt.Println(GetFLyWeight("a13"))
	fmt.Println(GetFLyWeight("a11"))
	fmt.Println(len(F))
}
