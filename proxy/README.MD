### 代理模式（Proxy Pattern）

##### 结构模式（Structural Pattern）

#### 定义

***Allows for object level access control by acting as a pass through entity or a placeholder object.***

*** ***

![Proxy Pattern UML](../images/proxy_pattern.png)

享元模式中提出了两个重要的概念，内部状态intrinsic和外部状态extrinsic。

内部状态是可以被共享的，外部状态是不可以被共享的，外部状态的数据应该是被传入的。

上面两句话可能有些拗口，很多时候享元模式都会以在线游戏来举例。游戏中一些角色是一样的，比如其外形、贴图等，这就是其内部状态，在内存中，可以只需要一份，这就实现了数据的共享；而这种游戏角色会在地图上出现多次，只是出现的位置、血量、其他数值会有差异，这就是外部状态，需要在内存中维护多份来体现差异。

### 优点
1. 尽可能实现对象复用，降低内存开销。
2. 因为对象复用，也减低了重新创建对象的性能开销。

### 例子分析

##### 邮件模版案例


### 参考文档：

https://en.wikipedia.org/wiki/Template_method_pattern

https://dzone.com/articles/design-patterns-template-method