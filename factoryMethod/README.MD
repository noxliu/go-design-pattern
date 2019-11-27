The idea behind the Factory Method pattern is that it allows for the case where a client doesn't know what concrete classes it will be required to create at runtime, but just wants to get a class that will do the job. The FactoryMethod builds on the concept of a simple Factory, but lets the subclasses decide which implementation of the concrete class to use.  You'll see factories used in logging frameworks, and in a lot of scenarios where the client doesn't need to know about the concrete implementations. It's a good approach to encapsulation.

优点：

1. 