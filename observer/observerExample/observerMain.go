package main

import (
	. "go-design-pattern/observer/observerExample/observer"
)

func main() {
	subject := NewSubject()
	reader1 := NewObserver("reader1")
	reader2 := NewObserver("reader2")
	reader3 := NewObserver("reader3")

	subject.Register(*reader1)
	subject.Register(*reader2)
	subject.Register(*reader3)

	subject.UpdateContext("来自Subject的消息，数据处理完成")
}
