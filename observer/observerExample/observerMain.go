package main

import (
	. "go-design-pattern/observer/observerExample/observer"
)

func main() {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")

	subject.Register(*reader1)
	subject.Register(*reader2)
	subject.Register(*reader3)

	subject.UpdateContext("Message from subject")
}
