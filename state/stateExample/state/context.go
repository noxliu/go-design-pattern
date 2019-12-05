package state

type Context struct {
	currentState State
}

type ContextStateNeedChange interface {
	SetState(s *State)
	Run()
	Stop()
}

func (c *Context) SetState(s *State) {
	c.currentState = *s
}

func (c *Context) Run() {
	c.currentState.Run()
}

func (c *Context) Stop() {
	c.currentState.Stop()
}
