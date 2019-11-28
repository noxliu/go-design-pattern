The idea behind the Factory Method pattern is that it allows for the case where a client doesn't know what concrete classes it will be required to create at runtime, but just wants to get a class that will do the job. The FactoryMethod builds on the concept of a simple Factory, but lets the subclasses decide which implementation of the concrete class to use.  You'll see factories used in logging frameworks, and in a lot of scenarios where the client doesn't need to know about the concrete implementations. It's a good approach to encapsulation.

优点：

1. 在运行时候，只是期望一个抽象的对象来执行任务，并不清晰的知道到底需要什么对象？在运行时并不明确知道，应该创建什么类型的问题。

you separated the use of Pizzas from its creation.

The advantage: you separated the use of Pizzas from its creation. You only want Pizza, independend of its creation process. You could enjoy Pizza without baking one yourself.

The downside: If you have several subtypes of the given flavours of Pizza (e.g. ROMANABerlinStyle, ROMANAHamburgStyle) you have two choices.

https://softwareengineering.stackexchange.com/questions/254923/what-are-the-disadvantages-of-a-simple-factory


https://zhidao.baidu.com/question/131823823.html


》》》》》》

https://blog.51cto.com/zero01/2067822