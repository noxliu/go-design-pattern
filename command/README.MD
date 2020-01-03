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

以上代码是没有使用任何工厂模式的用法，虽然并没有什么bug，但是如果代码中这样的接口实例化和使用存在多处:


```go
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
 
    ...

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

    ...
```

假设上面的代码片段大量存在于各种方法和函数中，正常情况下没有任何问题。但是假如汽车接口一旦发生了变化，比如对原有汽车结构体中的Size进行了拆分，分为了两个字段（wheelbase和length）,那么这个重构就会造成所有调用点都要改动，都要将size改成两个参数。

```go
	carTypeBmw := CarType{
		CarName: "bmw",
		Size:    9,
	}

>>

	carTypeBmw := CarType{
		CarName: "bmw",
		wheelbase:    2840,
        	length      : 4704,
	}
```

所有使用该汽车对象的地方都需要改动（有100处使用，就要改动100处）。

所以，上述代码也没有实现代码的复用，创建汽车类型的代码在很多地方重复，这是造成如果使用了100处就要改动这100个地方的原因。

我们使用简单工厂模式对上述代码进行改动。

增加一个工厂函数，用于产生汽车对象。
```go
func GetCar(carType string) Car {
	carType1 := CarType{
		carName: carType,
		size:    10,
	}
	if carType == "benz" {
		benz := Benz{CarType: carType1}
		return benz
	} else if carType == "bmw" {
		bmw := Bmw{CarType: carType1}
		return bmw
	} else {
		return nil
	}
}
```


#### 定义

***Exposes a method for creating objects, allowing subclasses to control the actual creation process.***

***定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法模式使一个类的实例化延迟到其子类。***

![Factory_Method Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/factory_method_patern.png)

### 优点

将对象的创建和使用完全良好分离，很好的提高了代码的重用度。

### 例子分析

##### 一. 汽车案例




### 总结
工厂方法模式是代码封装性的一个良好体现，很好的将对象的创建和使用隔离开。在开发中合理的使用，会体会到开闭原则的真正意义。

### 参考文档：
