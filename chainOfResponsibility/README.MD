### 享元模式（Fly Weight Pattern）

##### 结构模式（Structural Pattern）

#### 定义

***Facilitates the reuse of many fine grained objects, making the utilization of large numbers of objects more efficient.***

***提高对象的重用率，减少大量重复对象在内存中的创建开销和资源占用开销***

![Fly_Weight Pattern UML](../images/fly_weight_pattern.png)

享元模式中提出了两个重要的概念，内部状态intrinsic和外部状态extrinsic。

内部状态是可以被共享的，外部状态是不可以被共享的，外部状态的数据应该是被传入的。

上面两句话可能有些拗口，很多时候享元模式都会以在线游戏来举例。游戏中一些角色，在外观上是一样的，拥有相同的外形外形，这就是其内部状态，在内存中可以只需要一份，这样来实现了数据的共享；而这种游戏角色会在地图上不同的地方多次出现，出现的位置、血量、其他数值会有差异，这就是外部状态，需要在内存中维护多份来体现差异。

### 优点
1. 尽可能实现对象复用，降低内存开销。
2. 因为对象复用，也减低了重新创建对象的性能开销。

### 例子分析

##### 邮件模版案例
1. 某电子商务系统会发送消息提醒给客户以告知信息，不同类型的消息套用不同的模版，模版的数量有多种，下面列举几种以举例。

模版1
```go
感谢你的支持, 你购买的商品<<ProductName>>, 订单号<<OrderNo>>,已经发货, 快递单号为<<TrackNo>>
```

模版2
```go
客户你好, 你关注的商品<<ProductName>>, 正在促销打折! 数量不多!
```



### 参考文档：
