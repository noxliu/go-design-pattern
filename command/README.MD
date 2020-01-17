### 访问者模式（Visitor Pattern）

##### 行为模式（Behavioural Pattern）

#### 定义

***Represent an operation to be performed on elements of an object structure. Visitor lets you define a new operation without changing the classes of the elements on which it operates.***

***访问者模式用于操作某个结构中各个对象. 它可以不用改变对象本身, 对对象的新操作定义在访问者自己内部.***

![Visitor Pattern UML](../images/visitor_pattern.png)

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

### 参考文档：