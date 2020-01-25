### 命令模式（Strategy Pattern）

##### 行为模式（Behavioural Pattern）

#### 定义

***Encapsulates a request allowing it to be treated as an object. This allows the request to be handled in traditionally object based relationships such as queuing and callbacks.***

***将一个请求封装为一个对象，从而使你可用不同的请求对客户进行参数化；对请求排队或记录请求日志，以及支持可撤销的操作。***

![Strategy Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/command_pattern.png)

命令模式的核心在解耦请求者与执行者。

比如有一个订单系统，客户端程序如果直接调用服务端程序进行下单，下单之前要检查库存等一系列操作，然后还要记录下单日志，时间等。

看上去也能实现功能，但是客户端接触到服务端的代码就过多了（过于耦合）

使用命令模式，把每一笔订单提交交给一个调用者(invoker)来安排执行任务。这样就实现了客户端和服务端的解耦。

使用场景：
- 需要支持命令撤销/回滚操作
- 请求者对命令的执行顺序没有要求
- 请求的历史在需要的时候可以被记录为日志（请求历史可追溯）
- 命令的请求者和执行者需要被解耦

在回到上面提到的订单系统例子中，如果不是使用命令模式，要实现命令撤销/回滚操作的话，需要

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
https://cloud.tencent.com/developer/article/1056075
http://www.ayqy.net/blog/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F%E4%B9%8B%E5%91%BD%E4%BB%A4%E6%A8%A1%E5%BC%8F%EF%BC%88command-pattern%EF%BC%89/