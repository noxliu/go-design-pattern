### 适配器模式（Adapter Pattern）

##### 结构模式（Structural Pattern）

#### 定义

***Permits classes with disparate interfaces to work together by creating a common object by which they may communicate and interact.***

***将一个类的接口转换成客户希望的另外一个接口。使原本由于接口不兼容而不能一起工作的那些类可以一起工作。***

![Memento Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/adapter_pattern.png)

目标角色（target）：这是客户锁期待的接口。目标可以是具体的或抽象的类，也可以是接口

适配者角色（adaptee）：已有接口，但是和客户器期待的接口不兼容。

适配器角色（adapter）：将已有接口转换成目标接口。

### 优点
将适配器和适配者类解耦, 避免了因为要适配某个借口，直接修改适配者源码。


### 例子分析

##### 打印工具案例




### 参考文档：
https://en.wikipedia.org/wiki/Adapter_pattern