package memento

import "fmt"

type Game struct {
	HP int
	SP int
	MP int
}

func (g *Game) Play(HP int, SP int, MP int) {
	g.HP += HP
	g.SP += SP
	g.MP += MP
}

func (g *Game) Save() Memento {
	return Memento{
		HP: g.HP,
		SP: g.SP,
		MP: g.MP,
	}
}

func (g *Game) Load(m Memento) {
	g.HP = m.HP
	g.SP = m.SP
	g.MP = m.MP
}

func (g *Game) Status() {
	fmt.Printf("当前生命值 HP: %d，当前耐力: %d, 当前魔法力: %d\n", g.HP, g.SP, g.MP)
}
