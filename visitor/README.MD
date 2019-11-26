### 访问者模式（Visitor Pattern）

### 行为模式（Behavioural Pattern）

### 定义

***Represent an operation to be performed on elements of an object structure. Visitor lets you define a new operation without changing the classes of the elements on which it operates.***

***访问者模式用于操作某个结构中各个对象. 它可以不用改变对象本身, 对对象的新操作定义在访问者自己内部.***

![Visitor Pattern UML](../images/visitor_pattern.png ''Visitor Pattern UML'')

在访问者模式中

被访问的对象称为元素对象element

而访问对象就是访问者visitor

元素对象表示具备一定属性、数据、内容的对象，其不能因为有外部对象希望访问（读取）其内容，就改动自己的代码以提供对外方法，这样做的的缺陷是会污染到原对象的代码：新增加的给予外部对象提供数据的方法会和元素对象本身的业务代码混合在一起。

应对以上场景，访问者模式的解决方案是，让元素对象提供一个accept方法，允许接受者访问其内部对象。

然后把访问者需要实现的数据访问逻辑，全部移到访问方法"visitMethod"内部，而访问方法由访问者自己实现。

因此，不管增加多少个访问逻辑，只需要在需要访问的对象中增加accept方法以接受访问者的访问，然后在访问者中增加对数据的访问逻辑。

## 优点

综上所述，访问者模式的核心逻辑以及优势在于：

1. 将希望访问某对象数据并进行计算或者其他业务处理的操作封装到访问者自己内部，这样最大程度的降低了被访问对象的污染。



2. 在本例的代码中，Car结构体是作为要被访问的汽车对象。如果需要暴露某一个部件的访问权限，则实现

```greenplum
type Element interface {
	Accept(visitor Visitor)
}
```

方法即可


参考文档：

https://en.wikipedia.org/wiki/Visitor_pattern

https://dzone.com/articles/design-patterns-visitor