### 装饰模式（Decorator pattern）
 
##### 创建模式（Creational Pattern）
 
 
#### 定义
 
***Allows for the dynamic wrapping of objects in order to modify their existing responsibilities and behaviors.***
 
***动态地给一个对象添加一些额外的职责或者功能。***
 
![Decorator Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/decorator_method.png)

装饰器模式在日常研发中心是一个出场率非常高的模式，因为有太多的场景适用于该模式了。

装饰器模式常常用来动态的为已存在的功能增加新的功能，而不用对原有代码进行任何修改以造成影响。

比如，我们现在有一个订票功能OrderTicket。但是我们打算给他增加一些功能，比如订票之后发短信通知，或者发邮件通知。

听起来，似乎新建两个函数，SendMsgAfterOrderTicket（订票之后发短信），和SendEmailAfterOrderTicket（订票之后发邮件）。

这样我们在新的函数中实现发短信或者发邮件的功能，先调用原有的OrderTicket实现订票，然后在订票之后发送短信或者邮件。

但是如果我们需要既发送短信又发送邮件呢？如果在增加一个自动推送手机APP消息呢？如果希望只是选择性的选择其中的一种或者两种或者全部的通知呢？

重写的方式势必会比较多。如果继续增加通知方式，带来的更多麻烦则更不用说了。

这就是装饰器模式显示作用的地方了。

### 优点
 
 
 
### 例子分析
 
##### 购票案例


### 总结


### 参考文档：
