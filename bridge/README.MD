### 策略模式（Strategy Pattern）

##### 行为模式（Behavioural Pattern）

#### 定义

***Defines a set of encapsulated algorithms that can be swapped to carry out a specific behavior.***

***定义了一组算法，将每个算法都封装起来，并且使它们之间可以互换。***

![Strategy Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/strategy_pattern.png)

在现实研发中，我们或许有这样的场景存在，根据不同的角色或者对象，选择不同的算法，比如：

```go
	if userType == "A" {
		//业务算法逻辑A
	} else if userType == "B" {
		//业务算法逻辑B
	} else if userType == "C" {
		//业务算法逻辑C
	} else if userType == "D" {
		//业务算法逻辑D
	} else {
		//..
	}
```

### 优点

### 例子分析

##### 数据处理系统案例



### 参考文档：
https://www.runoob.com/design-pattern/strategy-pattern.html