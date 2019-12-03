package flyweight

//Flyweight Factory，享元工厂
var F = make(map[int]*string)

func GetMsgTemplate(templateNo int) string {
	msgTemplateBody := F[templateNo]
	if msgTemplateBody == nil {
		//模拟去数据库查询
		messageBody := QueryByType(templateNo)
		F[templateNo] = &messageBody
		msgTemplateBody = &messageBody
	}

	return *msgTemplateBody
}
