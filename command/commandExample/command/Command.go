package command

type Command struct {
	Receiver Receiver
}

func (c *Command) ExecCommand() {
	c.Receiver.DoSomething()
}
