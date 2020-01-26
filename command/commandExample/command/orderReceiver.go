package command

import "fmt"

type OrderReceiver struct {
}

func (*OrderReceiver) execOrder() {
	fmt.Println("下单")
}
