package state

type State interface {
	SetState(s *State)
	Run()
	Stop()
}
