### 观察者模式（Observer Pattern）

##### 行为模式（Behavioural Pattern）

#### 定义

***Lets one or more objects be notified of state changes in other objects within the system.***

***观察者模式又叫发布-订阅模式，定义对象间一（主题）对多（观察者）的依赖关系，使得每当主题对象改变状态，则所有依赖（观察）它的对象都会得到通知并自动更新。***

![Observer Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/observer_pattern.png)

观察者模式，顾名思义，是对象或者线程对某一对象或者线程进行持续性的关注，一旦被关注的对象或线程发生任何改变，能够及时的通知观察者，已做出正确的应对。

观察者模式在开发中非常常见，诸如某些管理系统，或者电子商务系统，创造观察者线程对需要关注的线程进行关注，一旦某种需要告知更多人知晓的事件发生，则可以通知观察者做出反应。比如某工厂的生产管理系统，一旦某些重要指标超过警戒线，需要及时的通知观察者线程，可以依据不同的问题状况构建不同的观察者，不同的观察者可以以不同的方式或渠道通知用户，比如发送短信，发送邮件，推送手机APP消息等，避免生产事故的发生。

### 优点
避免了直接在原有业务逻辑中加入通知相关业务逻辑造成代码耦合。

### 例子分析

##### 数据处理系统案例
1. 某数据处理系统，数据处理好之后可以供业务人员使用，但是数据处理时间较长，不可能一直等待。
2. 业务人员可以作为观察者，一直观察数据处理任务（主题），一旦数据数据里任务完成，则通知所有的业务人员，可以查看已经处理完成的数据了。

代码如下：

消息发布方（主题）：
```go
type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}
```

消息订阅方
```go
type Observer struct {
	name string
}

func NewObserver(name string) *Observer {
	return &Observer{
		name: name,
	}
}

func (r *Observer) Update(s *Subject) {
	fmt.Printf("%s receivd message: [%s]\n", r.name, s.context)
}
```

测试代码
```go
func main() {
	subject := NewSubject()
	observer1 := NewObserver("observer1")
	observer2 := NewObserver("observer2")
	observer3 := NewObserver("observer3")

	subject.Register(*observer1)
	subject.Register(*observer2)
	subject.Register(*observer3)

	subject.UpdateContext("来自Subject的消息，数据处理完成")
}
```

执行结果
```go
observer1 receivd message: [来自Subject的消息，数据处理完成]
observer2 receivd message: [来自Subject的消息，数据处理完成]
observer3 receivd message: [来自Subject的消息，数据处理完成]
```


### 参考文档：

https://www.tutorialspoint.com/design_pattern/observer_pattern.htm