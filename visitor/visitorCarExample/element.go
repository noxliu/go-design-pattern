package visitorCarExample

type Element interface {
	Accept(visitor Visitor)
}
