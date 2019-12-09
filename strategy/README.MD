### 策略模式（Strategy Pattern）

##### 行为模式（Behavioural Pattern）

#### 定义

***Defines a set of encapsulated algorithms that can be swapped to carry out a specific behavior.***

***定义了一组算法，将每个算法都封装起来，并且使它们之间可以互换。***

![Strategy Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/observer_pattern.png)

观察者模式，顾名思义，是对象或者线程对某一对象或者线程进行持续性的关注，一旦被关注的对象或线程发生任何改变，能够及时的通知观察者，已做出正确的应对。

观察者模式在开发中非常常见，诸如某些管理系统，或者电子商务系统，创造观察者线程对需要关注的线程进行关注，一旦某种需要告知更多人知晓的事件发生，则可以通知观察者做出反应。比如某工厂的生产管理系统，一旦某些重要指标超过警戒线，需要及时的通知观察者线程，可以依据不同的问题状况构建不同的观察者，不同的观察者可以以不同的方式或渠道通知用户，比如发送短信，发送邮件，推送手机APP消息等，避免生产事故的发生。

### 优点

### 例子分析

##### 数据处理系统案例



### 参考文档：

