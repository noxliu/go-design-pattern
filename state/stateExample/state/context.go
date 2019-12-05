package state

type Context struct {
	currentState State
}

func (c *Context) GetCurrentState() State {
	return c.currentState
}

func (c *Context) SetState(s State) {
	c.currentState = s
}

func (c *Context) Open() {
	c.currentState.Open()
	c.currentState = Open{}
}

func (c *Context) Close() {
	c.currentState.Close()
}

func (c *Context) Run() {
	c.currentState.Run()
}

func (c *Context) Stop() {
	c.currentState.Stop()
}
