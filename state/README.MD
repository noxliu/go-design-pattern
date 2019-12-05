### 状态模式（State Pattern）

##### 行为模式（Behavioural Pattern）

#### 定义

***Ties object circumstances to its behavior, allowing the object to behave in different ways based upon its internal state.***

***。***

![State Pattern UML](../images/state_pattern.png)

状态模式主要解决的问题就是当控制一个对象状态转换的条件表达式过于复杂时的情况。即把状态的判断逻辑转移到标识不同状态的一系列类当中。

状态模式在开发中非常有用，从开关按钮，到任务和购物车，都存在很多种不同的状态。

以为一个电梯程序为例，假设要控制电梯的四种状态：0代表电梯门打开，1代表电梯门关闭，2代表电梯在运行，3代表电梯停止（电梯门没有打开）

电梯能够在四种状态中进行一定的切换，但是有些状态之间的切换是不合理的，比如不能从 0 切换至 2，因为电梯门没有关闭的话，不能直接开动。

仔细想来，上面的代码也不难写，只要理清逻辑，明确哪些状态之间能够切换，然后进行合理的判断，就能写出代码，于是写出了下面的程序：

```go
var currentState = 0

//四种状态，0代表电梯门打开，1代表电梯门关闭，2代表电梯在运行，3代表电梯停止（电梯门没有打开）
var OPEN = 0
var CLOSE = 1
var RUN = 2
var STOP = 3

func stateChange(nextState int) {
	if currentState == OPEN {
		fmt.Println("关闭电梯门")
		currentState = CLOSE
	} else if currentState == CLOSE {
		fmt.Println("电梯开始运行")
		currentState = RUN
	} else if currentState == RUN {
		fmt.Print("停止电梯")
		currentState = STOP
	} else if currentState == STOP {
		fmt.Println("打开电梯门")
		currentState = OPEN
	}
}
```

上面的代码先不说是否有bug，但是看上去已经存在太多的if判断嵌套。或许可以将if换为switch case语句，但是也改变不了嵌套复杂的问题。

而且这还不是最糟糕的，当状态的数量和切换变得更加复杂之后，麻烦会更加严重的增加。比如将RUN状态改为UP或者DOWN两种状态，分别为上升运行和下降运行，然后再加入维护状态和紧急状态等状态。可以想象的是，上面的代码将会有非常复杂的if嵌套。

代码的可读性和可维护性将非常低，即便在耗费大量精力之后保证逻辑正确，一旦要增加或者删减状态，将又要花费很大的精力来保证改动后的代码的正确性。

state模式正是为了解决状态切换的需求而存在的。

状态模式中三个角色：

1 上下文环境（Context）：它定义了客户程序需要的接口并维护一个具体状态角色的实例，将与状态相关的操作委托给当前的Concrete State对象来处理。

2 抽象状态（State）：定义一个接口以封装使用上下文环境的的一个特定状态相关的行为。

3 具体状态（Concrete State）：实现抽象状态定义的接口。




### 优点
1. 封装和控制了业务的不变部分。
2. 实现代码复用，将公共功能部分代码抽取出来，更易于维护。

### 例子分析

//电梯例子

##### 下载工具案例

### 参考文档：

https://en.wikipedia.org/wiki/Template_method_pattern

https://dzone.com/articles/design-patterns-template-method