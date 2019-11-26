package visitorCarExample

type VisitBody struct {
}

func (*VisitBody) VisitBody(body Body) string {
	return body.BodyColor
}
