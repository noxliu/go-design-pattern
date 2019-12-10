package main

import (
	"fmt"
	. "go-design-pattern/memento/mementoExample/memento"
)

func main() {
	game := &Game{
		HP: 10,
		SP: 8,
		MP: 6,
	}

	game.Status()
	progress := game.Save()
	game.Play(-10, -8, -6)
	game.Status()
	fmt.Println("游戏Game Over, 重新读档")
	game.Load(progress)
	game.Status()

}
