package memento

import "fmt"

type Game struct {
	Hp, Mp int
}

func (g *Game) Play(mpDelta, hpDelta int) {
	g.Mp += mpDelta
	g.Hp += hpDelta
}

func (g *Game) Save() Memento {
	return &gameMemento{
		Hp: g.Hp,
		Mp: g.Mp,
	}
}

func (g *Game) Load(m Memento) {
	gm := m.(*gameMemento)
	g.Mp = gm.Mp
	g.Hp = gm.Hp
}

func (g *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", g.Hp, g.Mp)
}
