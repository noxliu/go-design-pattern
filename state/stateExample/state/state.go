package state

type State interface {
	Open()
	Close()
	Run()
	Stop()
}
