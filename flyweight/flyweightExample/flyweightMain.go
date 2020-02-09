package main

import (
	. "./flyweight"
	"fmt"
)

func main() {
	message1 := Message{"三星手机", "4884741244", "8514484122", 1}
	message2 := Message{"华为手机", "1211223122", "3221234133", 2}
	message3 := Message{"oppo手机", "7656754133", "9556435133", 3}
	message4 := Message{"小米手机", "8976897144", "8477474413", 1}
	message5 := Message{"荣耀手机", "5437659433", "0000211133", 1}
	message6 := Message{"格力空调", "7656223133", "0056423323", 3}
	message7 := Message{"长虹冰箱", "7623323133", "0056435133", 3}
	message8 := Message{"格兰仕微波炉", "72533878122", "5452321133", 3}

	fmt.Println(message1.Merge(GetMsgTemplate(message1.TemplateNo)))
	fmt.Println(message2.Merge(GetMsgTemplate(message2.TemplateNo)))
	fmt.Println(message3.Merge(GetMsgTemplate(message3.TemplateNo)))
	fmt.Println(message4.Merge(GetMsgTemplate(message4.TemplateNo)))
	fmt.Println(message5.Merge(GetMsgTemplate(message5.TemplateNo)))
	fmt.Println(message6.Merge(GetMsgTemplate(message6.TemplateNo)))
	fmt.Println(message7.Merge(GetMsgTemplate(message7.TemplateNo)))
	fmt.Println(message8.Merge(GetMsgTemplate(message8.TemplateNo)))

	fmt.Println("缓存的模版数量：", len(F))
}
