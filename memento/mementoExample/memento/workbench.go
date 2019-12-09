package memento

type Workbench struct {
	Words     string
	SavedTime string
}

func (w *Workbench) Save() Memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}
