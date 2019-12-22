### 工厂方法模式（Factory Method pattern）

##### 创建模式（Creational Pattern）


在讨论工厂方法模式之前，我们首先要讨论的一定是简单工厂(Simple Factory)，严格的说，简单工厂不是Gof23设计模式中的一种。

首先我们看看不使用简单工厂或者工厂方法之前，我们要创建对象的方式:

假设我们有汽车类型接口及其两种实现：一个benz，一个bwm

汽车类型car：

```go
type Car interface {
	Drive() string
}

type CarType struct {
	CarName string
	Size    int
}
```

两种实现：
```go
type Benz struct {
	CarType CarType
}

func (b Benz) Drive() string {
	return b.CarType.CarName
}

type Bmw struct {
	CarType CarType
}

func (b Bmw) Drive() string {
	return b.CarType.CarName
}
```

对两种汽车类型的调用：
```go
func main() {
	carTypeBenz := CarType{
		CarName: "benz",
		Size:    10,
	}

	benz := Benz{CarType: carTypeBenz}
	fmt.Println(benz.Drive())

	carTypeBmw := CarType{
		CarName: "bmw",
		Size:    9,
	}

	bmw := Benz{CarType: carTypeBmw}
	fmt.Println(bmw.Drive())
}
```


#### 定义

***Exposes a method for creating objects, allowing subclasses to control the actual creation process.***

***定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法模式使一个类的实例化延迟到其子类。***

![Factory_Method Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/factory_method_patern.png)

### 优点


### 例子分析

##### 一. 汽车案例

### 总结

### 参考文档：
