package main

import (
	. "go-design-pattern/observer/observerExample/observer"
)

func main() {
	subject := NewSubject()
	observer1 := NewObserver("observer1")
	observer2 := NewObserver("observer2")
	observer3 := NewObserver("observer3")

	subject.Register(*observer1)
	subject.Register(*observer2)
	subject.Register(*observer3)

	subject.UpdateContext("来自Subject的消息，数据处理完成")
}
