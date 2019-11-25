package element

import . "go-design-pattern/visitor/visitor"

type Element interface {
	Accept(visitor Visitor)
}
