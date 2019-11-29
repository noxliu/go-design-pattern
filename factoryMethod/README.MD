The idea behind the Factory Method pattern is that it allows for the case where a client doesn't know what concrete classes it will be required to create at runtime, but just wants to get a class that will do the job. The FactoryMethod builds on the concept of a simple Factory, but lets the subclasses decide which implementation of the concrete class to use.  You'll see factories used in logging frameworks, and in a lot of scenarios where the client doesn't need to know about the concrete implementations. It's a good approach to encapsulation.

优点：

1. 在运行时候，只是期望一个抽象的对象来执行任务，并不清晰的知道到底需要什么对象？在运行时并不明确知道，应该创建什么类型的问题。

you separated the use of Pizzas from its creation.

The advantage: you separated the use of Pizzas from its creation. You only want Pizza, independend of its creation process. You could enjoy Pizza without baking one yourself.

The downside: If you have several subtypes of the given flavours of Pizza (e.g. ROMANABerlinStyle, ROMANAHamburgStyle) you have two choices.

https://softwareengineering.stackexchange.com/questions/254923/what-are-the-disadvantages-of-a-simple-factory


https://zhidao.baidu.com/question/131823823.html


》》》》》》

https://blog.51cto.com/zero01/2067822


### 工厂方法模式（Factory Method pattern）

##### 创建模式（Creational Pattern）

#### 定义

***Exposes a method for creating objects, allowing subclasses to control the actual creation process.***

***定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法模式使一个类的实例化延迟到其子类。***

![Visitor Pattern UML](../images/visitor_pattern.png)


在讨论工厂方法模式之前，我们首先要讨论的一定是简单工厂(Simple Factory)，严格的说，简单工厂不是Gof23设计模式中的一种。

首先我们看看不使用简单工厂或者工厂方法之前，我们要创建对象的方式。




在访问者模式中

被访问的对象称为**元素对象**element

而访问对象称为**访问者**visitor

元素对象表示具备一定属性、数据、内容的对象，其不能因为有外部对象希望访问（读取）其内容，就改动自己的代码以提供对外方法，这样做的的缺陷是会污染到原对象的代码：新增加的给予外部对象提供数据的方法会和元素对象本身的业务代码混合在一起。

1. 应对以上场景，访问者模式的解决方案是，当有访问需要访问元素对象的时候，元素对象提供一个accept方法，使得元素对象具备了可访问性，允许接受者访问其内部对象。

```go
type Element interface {
	Accept(visitor Visitor)
}

func (e Engine) Accept(visitor Visitor) {
	visitor.VisitEngine(e)
}
```

2. 访问者将数据访问逻辑全部移到访问方法"visitMethod"内部，访问方法由访问者自己实现，这样就避免了对元素对象的污染。

```go
func (*VisitEngine) VisitEngine(engine Engine) string {
	return engine.EngineType
}
```

### 优点
实现了元素本身的可访问性，也保证了逻辑的清晰隔离：外部的访问逻辑不会和原有的业务逻辑混在一起。

### 例子分析

##### 一. 汽车案例
某个二手汽车相关的业务系统，汽车对象有若干字段，比如轮胎尺寸、轴距、马力、使用时间（年）等。

目前有个需求是，在某节日，可能会根据汽车的某些特性进行附加的礼品赠送，这就要求能够对汽车对象的某些属性能够访问以判断是否满足赠送条件。

首先能够想到的第一个做法，就是在客户下单之前，在程序中取出汽车的相关字段进行计算，如果达到条件，则进行礼品赠送。但是这里的坏处就是要修改原有的业务逻辑，而且一旦赠送活动结束，此处增加的业务代码可能成为无用的废代码，污染了原有程序。

另外一个做法则是使用访问者模式，此处只要实现能够访问汽车对象（通过增加accept方法使汽车对象具备可访问性），则可以在不破坏原有业务逻辑的前提下，编写并上线访问者相关的程序，按照访问者模式的设计，新增加的业务逻辑全部增加在访问者的visit方法中，即便某一套折扣逻辑未来要下线停用或者更改赠送逻辑，也是在访问者者相关的代码中更改，总之不会影响到现有的标准逻辑。

具体代码详见 

##### 二. 购物车案例
某供应链系统，部分员工的工作是在该系统在线采购公司所需要的原材料，不同种类材料有不同价格以及其他复杂属性。业务经理作为资深领导，有时候要实时的关注员工正在采购的工作，如果发现问题需要进行帮助和建议。

因为员工采购的材料可能有多个属性，需要一些附加的业务计算逻辑帮经理判断员工正在采购的工作的状况，如果直接让经理的访问逻辑加入到员工的业务逻辑代码中以帮助经理监督过程，这样会侵入原有的代码。

所以，此处最合适的方法也依然是实现访问者模式，让员工对象实现可访问性。

### 参考文档：

https://en.wikipedia.org/wiki/Visitor_pattern

https://dzone.com/articles/design-patterns-visitor