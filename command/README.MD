### 命令模式（Strategy Pattern）

##### 行为模式（Behavioural Pattern）

#### 定义

***Encapsulates a request allowing it to be treated as an object. This allows the request to be handled in traditionally object based relationships such as queuing and callbacks.***

******

![Strategy Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/command_pattern.png)

命令模式的核心在解耦请求者与执行者。

使用场景：
- 需要支持回滚操作
- 请求者对命令的执行顺序没有要求
- 请求的历史在需要的时候可以被记录为日志（请求历史可追溯）
- 命令的请求者和执行者需要被解耦

如果不使用命令模式，我们的做法无非就是:

```go

```
如果考虑到命令模式的思路，可以在客户端的请求发送类和请求处理类之间增加一个Invoker类（调用者），将请求发送类发送的所有请求封装成对象，然后让Invoker类去管理这些请求对象，并决定这些请求是否允许执行、何时执行、按什么顺序执行。

请求者：发起要求
调用者：安排命令
执行者：执行命令

### 优点

### 例子分析

##### 数据处理系统案例

### 参考文档：
