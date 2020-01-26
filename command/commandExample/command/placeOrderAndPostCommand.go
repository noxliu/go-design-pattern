package command

type PlaceOrderAndPostCommand struct {
	receiver     OrderReceiver
	postReceiver PostReceiver
}

func (c *PlaceOrderAndPostCommand) ExecCommand() {
	c.receiver.execOrder()
	c.postReceiver.execPost()
}
