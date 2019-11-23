package main

import "fmt"

type FlyWeight struct {
	data string
}

var F = make(map[string]*FlyWeight)

func GetFLyWeight(fileName string) *FlyWeight {
	image := F[fileName]
	if image == nil {
		image = &FlyWeight{data: fileName}
		F[fileName] = image
	}

	return image
}

func main() {
	fmt.Println(GetFLyWeight("a11"))
	fmt.Println(GetFLyWeight("a12"))
	fmt.Println(GetFLyWeight("a13"))
	fmt.Println(GetFLyWeight("a11"))
	fmt.Println(GetFLyWeight("a12"))
	fmt.Println(GetFLyWeight("a13"))
	fmt.Println(GetFLyWeight("a11"))
	fmt.Println(len(F))
}
