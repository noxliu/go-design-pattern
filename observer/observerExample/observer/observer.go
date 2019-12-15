package observer

import "fmt"

type Observer struct {
	name string
}

func NewObserver(name string) *Observer {
	return &Observer{
		name: name,
	}
}

func (r *Observer) Update(s *Subject) {
	fmt.Printf("%s receivd message: [%s]\n", r.name, s.context)
}
