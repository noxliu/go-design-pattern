package visitor

import "fmt"

type VisitBody struct {
}

func (*VisitBody) VisitBody(body Body) string {
	fmt.Println(body.BodyColor)
	fmt.Println(body.Size)
	return "visit body"
}
