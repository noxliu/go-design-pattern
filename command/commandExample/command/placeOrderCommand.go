package command

type PlaceOrderCommand struct {
	receiver OrderReceiver
}

func (c *PlaceOrderCommand) ExecCommand() {
	c.receiver.execOrder()
}
