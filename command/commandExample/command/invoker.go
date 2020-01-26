package command

import "fmt"

type Invoker struct {
	orders []Command
}

func (b *Invoker) AddOrder(command Command) {
	b.orders = append(b.orders, command)
}

func (b *Invoker) ExecuteOrders() {
	for _, command := range b.orders {
		fmt.Println(">>>>处理命令<<<<")
		command.ExecCommand()
		fmt.Println("记录日志")
		fmt.Println(">>>>处理结束<<<<")

	}
}
