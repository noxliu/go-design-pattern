package command

type Invoker struct {
	orders []Command
}

func (b *Invoker) AddOrder(command Command) {
	b.orders = append(b.orders, command)
}

func (b *Invoker) ExecuteOrders() {
	for _, command := range b.orders {
		command.ExecCommand()
	}
}
