package flyweight

import "strings"

type Message struct {
	ProductName string
	OrderNo     string
	TrackNo     string
	TemplateNo  int
}

func (m *Message) Merge(MessageTemplate string) string {
	message := strings.ReplaceAll(MessageTemplate, "\\<\\<ProductName\\>\\>", m.ProductName)
	message = strings.ReplaceAll(message, "<<OrderNo>>", m.OrderNo)
	message = strings.ReplaceAll(message, "<<TrackNo>>", m.TrackNo)
	return message
}
