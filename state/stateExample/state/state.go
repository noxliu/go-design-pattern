package state

type State interface {
	SetState(s State)
	Open()
	Close()
	Run()
	Stop()
}
