package flyweight

import "strings"

type Message struct {
	ProductName string
	OrderNo     string
	TrackNo     string
	TemplateNo  int
}

func (m *Message) Merge(MessageTemplate string) string {
	message := strings.ReplaceAll(MessageTemplate, "\\<\\<productName\\>\\>", m.ProductName)
	message = strings.ReplaceAll(MessageTemplate, "<<orderNo>>", m.OrderNo)
	message = strings.ReplaceAll(MessageTemplate, "<<trackNo>>", m.TrackNo)
	return message
}
