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

上面的代码的功能是根据用户类型或者其他参数，来选择算法逻辑进行计算，上面的版本还没有if嵌套，一旦嵌套将更加复杂。

更大的麻烦是，如果其中一些逻辑发生变化，势必需要重构代码，但具备一定开发经验的同行都知道，一旦if嵌套过多，改动起来要去梳理整个所有的if结构，麻烦会很大。很容易带来新的BUG。

策略模式的核心思想就是将上面的算法封装成独立的类，一个算法一个类，以为根据上下文灵活的选择算法的需求，也能实现算法的灵活切换。

### 优点
1. 算法可以自由切换
2. 避免使用多重条件判断（如果不用策略模式我们可能会使用多重条件语句，不利于维护）
3. 扩展性良好，增加一个策略只需实现接口即可

### 例子分析
工资计算系统，根据不同的员工类型，要选择不同的算法计算工资。

代码实现：

算法接口：
```go
type SalaryAlgorithm interface {
	Calculate(b BaseInfo) int
}

type BaseInfo struct {
	UserId     string
	Name       string
	BaseSalary int
}
```

```go

```

##### 数据处理系统案例



### 参考文档：

