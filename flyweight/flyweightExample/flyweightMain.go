package main

import (
	"fmt"
	. "go-design-pattern/flyweight/flyweightExample/flyweight"
)

func main() {
	message1 := Message{"三星手机", "122", "122", 1}
	message2 := Message{"华为手机", "122", "133", 2}
	message3 := Message{"oppo手机", "133", "133", 3}
	message4 := Message{"小米手机", "144", "13", 1}
	message5 := Message{"荣耀手机", "1433", "133", 1}

	fmt.Println(GetMsgTemplate(message1.TemplateNo))
	fmt.Println(GetMsgTemplate(message2.TemplateNo))
	fmt.Println(GetMsgTemplate(message3.TemplateNo))
	fmt.Println(GetMsgTemplate(message4.TemplateNo))
	fmt.Println(GetMsgTemplate(message5.TemplateNo))

	fmt.Println(len(F))
}
