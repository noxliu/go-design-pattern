### 享元模式（Fly Weight Pattern）

##### 结构模式（Structural Pattern）

#### 定义

***Facilitates the reuse of many fine grained objects, making the utilization of large numbers of objects more efficient.***

***提高对象的重用率，减少大量重复对象在内存中的创建开销和资源占用开销***

![Fly_Weight Pattern UML](../images/fly_weight_pattern.png)


### 优点
1. 尽可能实现对象复用，降低内存开销。
2. 因为对象复用，也减低了重新创建对象的性能开销。

### 例子分析

##### 邮件模版案例
1. 某电子商务系统会发送消息提醒给客户以告知信息，不同类型的消息套用不同的模版，模版的数量有多种，下面列举几种以举例。

模版1
```go
感谢你的支持, 你购买的商品<<ProductName>>, 订单号<<OrderNo>>,已经发货, 快递单号为<<TrackNo>>
```

模版2
```go
客户你好, 你关注的商品<<ProductName>>, 正在促销打折! 数量不多!
```

模版3
```go
客户你好, 我们已经你对订单<<OrderNo>>的评价, 感谢!
```

2. 当需要给客户发送信息时候，只需要调用正确的模版然后套用相关数据发送信息即可。

3. 代码示例

message 结构体是不可共享的外部状态，是消息的变化的内容。
```go
type Message struct {
	ProductName string
	OrderNo     string
	TrackNo     string
	TemplateNo  int
}

func (m *Message) Merge(MessageTemplate string) string {
	message := strings.ReplaceAll(MessageTemplate, "<<ProductName>>", m.ProductName)
	message = strings.ReplaceAll(message, "<<OrderNo>>", m.OrderNo)
	message = strings.ReplaceAll(message, "<<TrackNo>>", m.TrackNo)
	return message
}
```




### 参考文档：
