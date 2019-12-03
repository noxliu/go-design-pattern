package flyweight

type MsgTemplate struct {
	data string
}

var F = make(map[string]*MsgTemplate)

func GetMsgTemplate(msgTemplate string) *MsgTemplate {
	msgTemplateBody := F[msgTemplate]
	if msgTemplateBody == nil {
		msgTemplateBody = &MsgTemplate{data: msgTemplate}
		F[msgTemplate] = msgTemplateBody
	}

	return msgTemplateBody
}
