package main

import (
	. "go-design-pattern/memento/mementoExample/memento"
)

func main() {
	game := &Game{
		HP: 10,
		SP: 10,
		MP: 10,
	}

	game.Status()
	progress := game.Save()

	game.Play(-2, -3, -5)
	game.Status()

	game.Load(progress)
	game.Status()

}
