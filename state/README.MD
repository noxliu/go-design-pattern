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

```

### 优点
1. 封装和控制了业务的不变部分。
2. 实现代码复用，将公共功能部分代码抽取出来，更易于维护。

### 例子分析

//电梯例子

##### 下载工具案例

### 参考文档：

https://en.wikipedia.org/wiki/Template_method_pattern

https://dzone.com/articles/design-patterns-template-method