package command

type PlaceOrderCommand struct {
	Receiver Receiver
}

func (c *PlaceOrderCommand) ExecCommand() {
	c.Receiver.DoSomething()
}
