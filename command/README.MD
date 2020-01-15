### 抽象工厂模式（Builder Pattern）

##### 创建模式（Creational Pattern）

#### 定义

******

******


![Builder Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/builder_pattern.png)

抽象工厂提出了产品族概念，这和工厂方法模式是有差别的。

在工厂方法模式中，我们的设计是建立在所有的是同一类型且类型单一的基础上的，比如要么都是车，要么都是杯子。

但如果发生产品之间还要划分子类别的时候，工厂方法模式就显得不是那么舒服了。

比如，我们有一个生产电脑的接口，比如生产dell电脑，hp电脑，lenovo电脑等品牌电脑，但同时，我们又生产这些品牌的鼠标。

如果用工厂方法模式，会是什么样子呢？

各种工厂：
```go
func (*DellPcFactory) GetPc() Pc {}
func (*HpPcFactory) GetPc() Pc {}
func (*LenovoPcFactory) GetPc() Pc {}
func (*DellMouseFactory) GetMouse() Mouse {}
func (*HpMouseFactory) GetMouse() Mouse {}
func (*LenovoMouseFactory) GetMouse() Mouse {}
```



这样我们有分别三个电脑产品，三个鼠标产品，就有了六个工厂。但显然，同一品牌的电脑和鼠标是有关系的，上面的代码不光产生了太多的类型，代码的内聚性也不好。

抽象工厂模式则是为了解决以上问题提出的。

### 优点

### 例子分析

##### 电脑案例


### 参考文档：