package main

import (
	. "go-design-pattern/memento/mementoExample/memento"
)

func main() {
	game := &Game{
		Hp: 10,
		Mp: 10,
	}

	game.Status()
	progress := game.Save()

	game.Play(-2, -3)
	game.Status()

	game.Load(progress)
	game.Status()

}
