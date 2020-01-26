### 命令模式（Strategy Pattern）

##### 行为模式（Behavioural Pattern）

#### 定义

***Encapsulates a request allowing it to be treated as an object. This allows the request to be handled in traditionally object based relationships such as queuing and callbacks.***

***将一个请求封装为一个对象，从而使你可用不同的请求对客户进行参数化；对请求排队或记录请求日志，以及支持可撤销的操作。***

![Strategy Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/command_pattern.png)

命令模式的核心在解耦请求者与执行者。

比如有一个订单系统，客户端程序如果直接调用服务端程序进行下单，下单之前要检查库存等一系列操作，然后还要记录下单日志，时间等。

看上去也能实现功能，但是客户端接触到服务端的代码就过多了（过于耦合）

而使用命令模式，把每一笔订单提交交给一个调用者(invoker)来安排执行任务。这样就实现了请求者和执行者的解耦。

命令模式使用场景：
- 需要支持命令撤销/回滚操作
- 请求者对命令的执行顺序没有要求
- 请求的历史在需要的时候可以被记录为日志（请求历史可追溯）
- 命令的请求者和执行者需要被解耦

模式中三角色职能划分:
- 请求者：发起要求
- 调用者：安排命令
- 执行者：执行命令

### 优点
将请求和执行的两端进行解耦，使代码中不同部分的权责更加清晰。

### 例子分析
在回到上面提到的订单系统例子中，如果不是使用命令模式，而要求实现命令撤销/回滚操作的话，则需要客户端实现和维护相关撤销/回滚代码。如果多个客户端在同时提交命令，需要对这些命令进行优先级排队（某些客户端可能应该具备较高的执行权重），那么也是相当困难的事情，因为这需要多个客户端之间进行协调。

如果需要记录执行日志，也是一个麻烦工作，因为客户端直接调用了服务端的一些方法来完成了任务，需要重新维护一套日志功能。

比如服务端有三个功能：

1. 完成订单。
2. 将订单信息提交到快递商安排发货。

有两个不同的客户端要调用该服务端，第一个需要顺序执行1，2；而第二个不需要发货，只需要执行1。

那么记录日志有两种办法，一种是规定一个统一的接口，在客户端执行完任务之后，调用该接口提交日志；
一种是服务端自己维护部分代码，首先判断客户端调用的类型，然后确认客户端执行完成之后（有可能是12骤，有可能只有1步骤，所以需要判断），然后记录日志。

上面两种日志方式，都因为客户端和服务端（请求者和执行者）的耦合造成了日志功能的实现显得很不优雅。

而命令模式很好的解决了该问题，因为命令模式在请求者和执行者之间增加了调用者，可以由调用者来安排命令的执行，同时也很好的完成了命令的日志记录功能。

所以命令模式的思路，是在客户端的请求发送类和请求处理类之间增加一个Invoker类（调用者），将请求发送类发送的所有请求封装成对象，然后让Invoker类去管理这些请求对象，并决定这些请求是否允许执行、何时执行、按什么顺序执行。

##### 订单处理案例
订单处理receiver
```go
type OrderReceiver struct {
}

func (*OrderReceiver) execOrder() {
	fmt.Println("下单")
}
```

订单发货receiver
```go
type PostReceiver struct {
}

func (*PostReceiver) execPost() {
	fmt.Println("发货")
}
```

command接口
```go
type Command interface {
	ExecCommand()
}
```

仅下单功能命令
```go
type PlaceOrderCommand struct {
	receiver OrderReceiver
}

func (c *PlaceOrderCommand) ExecCommand() {
	c.receiver.execOrder()
}
```

先下单再发货命令
```go
type PlaceOrderAndPostCommand struct {
	receiver     OrderReceiver
	postReceiver PostReceiver
}

func (c *PlaceOrderAndPostCommand) ExecCommand() {
	c.receiver.execOrder()
	c.postReceiver.execPost()
}
```

invoker
```go
type Invoker struct {
	orders []Command
}

func (b *Invoker) AddOrder(command Command) {
	b.orders = append(b.orders, command)
}

func (b *Invoker) ExecuteOrders() {
	for _, command := range b.orders {
		fmt.Println(">>>>处理命令<<<<")
		fmt.Println("判断库存")
		command.ExecCommand()
		fmt.Println("记录日志")
		fmt.Println(">>>>处理结束<<<<")
	}
}
```

例子代码
```go
func main() {
	onlyPlaceOrdercommand := PlaceOrderCommand{}
	placeOrderAndPostCommand := PlaceOrderAndPostCommand{}
	invoker := Invoker{}
	invoker.AddOrder(&onlyPlaceOrdercommand)
	invoker.AddOrder(&placeOrderAndPostCommand)
	invoker.ExecuteOrders()
}
```

执行结果
```go
>>>>处理命令<<<<
判断库存
下单
记录日志
>>>>处理结束<<<<
>>>>处理命令<<<<
判断库存
下单
发货
记录日志
>>>>处理结束<<<<
```

### 参考文档：
http://www.ayqy.net/blog/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F%E4%B9%8B%E5%91%BD%E4%BB%A4%E6%A8%A1%E5%BC%8F%EF%BC%88command-pattern%EF%BC%89/