### 建造者模式（Builder Pattern）

##### 创建模式（Creational Pattern）

#### 定义

***Allows for the dynamic creation of objects based upon easily interchangeable algorithms.***

***使用多个简单的对象一步一步构建成一个复杂的对象。建造者模式，就是把复杂，易出错的进行封装。便于创建对象。将创建和展现分开***

将一个复杂的构建与表示分离，使得同样的材料，因为有了不同的构建过程，会有不一样的构建结果（不一样的表示）。

![Builder Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/builder_pattern.png)


### 优点
1. 建造者模式有利于创造复杂对象，利用不同的执行（建造）顺序，可以制造出不同的结果。
2. 可以将多个部件组装到同一个对象中，但是因为不同的建造顺序，可以制造出不同的结果。

### 例子分析

##### 电脑案例
1. 一台电脑的装配需要多种零部件配合。

2. 通过一些相同的零部件，加上一些差异化的零部件，可以分别组装出不同类型的电脑，比如台式电脑和笔记本电脑。

3. 代码如下：


创建者接口
```go
type ComputerBuilder interface {
	SetCpu(cpu string) ComputerBuilder
	SetMemory(memory string) ComputerBuilder
	SetHardDisk(hardDisk string) ComputerBuilder
	SetScreenSize(screenSize int) ComputerBuilder
	GetComputerProduct() ComputerItems //显示结果
	SetConstruct(construct Construct) ComputerBuilder
	Build() ComputerItems
}
```

创建者实现
```go

type ComputerBuilderImplement struct {
	ComputerItems ComputerItems
}

func (c ComputerBuilderImplement) SetCpu(cpu string) ComputerBuilder {
	c.ComputerItems.Cpu = cpu
	return c
}

func (c ComputerBuilderImplement) SetMemory(memory string) ComputerBuilder {
	c.ComputerItems.Memory = memory
	return c
}

func (c ComputerBuilderImplement) SetHardDisk(hardDisk string) ComputerBuilder {
	c.ComputerItems.HardDisk = hardDisk
	return c
}

func (c ComputerBuilderImplement) SetScreenSize(screenSize int) ComputerBuilder {
	c.ComputerItems.ScreenSize = screenSize
	return c
}

func (c ComputerBuilderImplement) GetComputerProduct() ComputerItems {
	return c.ComputerItems
}

func (c ComputerBuilderImplement) SetConstruct(construct Construct) ComputerBuilder {
	c.ComputerItems.Construct = construct
	return c
}

func (c ComputerBuilderImplement) Build() ComputerItems {
	return c.ComputerItems.Construct.BuildComputer(c.ComputerItems)
}

```

Construct接口
```go
type Construct interface {
	BuildComputer(ComputerItems ComputerItems) ComputerItems
}
```

台式机Construct实现
```go
type ComputerConstruct struct {
	ComputerBuilder ComputerBuilder
}

func (con ComputerConstruct) BuildComputer(ComputerItems ComputerItems) ComputerItems {
	ComputerItems.ComputerProduct = "台式电脑：" +
		ComputerItems.Cpu + " " +
		ComputerItems.Memory + " " +
		ComputerItems.HardDisk + " " +
		strconv.Itoa(ComputerItems.ScreenSize) + "寸显示屏"
	return ComputerItems
}
```

笔记本Construct实现，可以看到和上面的台式电脑在构建显示屏的时候有差异
```go
type LaptopConstruct struct {
	ComputerBuilder ComputerBuilder
}

func (con LaptopConstruct) BuildComputer(ComputerItems ComputerItems) ComputerItems {
	ComputerItems.ComputerProduct = "笔记本电脑：" +
		ComputerItems.Cpu + " " +
		ComputerItems.Memory + " " +
		ComputerItems.HardDisk + " " +
		strconv.Itoa(ComputerItems.ScreenSize) + "寸触摸显示屏"
	return ComputerItems
}
```

测试代码
```go
func main() {
	builder := ComputerBuilderImplement{}
	ComputerConstruct := ComputerConstruct{}
	builder.SetCpu("Intel i7").
		SetMemory("16GB").
		SetHardDisk("1TSSD").
		SetScreenSize(21).
		SetConstruct(ComputerConstruct).
		Build().PrintComputerInfo()

	LaptopConstruct := LaptopConstruct{}
	builder.SetCpu("Intel i7").
		SetMemory("16GB").
		SetHardDisk("1TSSD").
		SetScreenSize(21).
		SetConstruct(LaptopConstruct).
		Build().PrintComputerInfo()
}
```

执行结果：
```go
台式电脑：Intel i7 16GB 1TSSD 21寸显示屏
笔记本电脑：Intel i7 16GB 1TSSD 21寸触摸显示屏
```



### 参考文档：

https://www.jianshu.com/p/a200c63adf2e


































### 建造者模式（Builder Pattern）

##### 创建模式（Creational Pattern）

#### 定义

***Allows for the dynamic creation of objects based upon easily interchangeable algorithms.***

***使用多个简单的对象一步一步构建成一个复杂的对象。建造者模式，就是把复杂，易出错的进行封装。便于创建对象。将创建和展现分开***

将一个复杂的构建与表示分离，使得同样的材料，因为有了不同的构建过程，会有不一样的构建结果（不一样的表示）。

![Builder Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/builder_pattern.png)


### 优点
1. 建造者模式有利于创造复杂对象，利用不同的执行（建造）顺序，可以制造出不同的结果。
2. 可以将多个部件组装到同一个对象中，但是因为不同的建造顺序，可以制造出不同的结果。

### 例子分析

##### 电脑案例
1. 一台电脑的装配需要多种零部件配合。

2. 通过一些相同的零部件，加上一些差异化的零部件，可以分别组装出不同类型的电脑，比如台式电脑和笔记本电脑。

3. 代码如下：


创建者接口
```go
type ComputerBuilder interface {
	SetCpu(cpu string) ComputerBuilder
	SetMemory(memory string) ComputerBuilder
	SetHardDisk(hardDisk string) ComputerBuilder
	SetScreenSize(screenSize int) ComputerBuilder
	GetComputerProduct() ComputerItems //显示结果
	SetConstruct(construct Construct) ComputerBuilder
	Build() ComputerItems
}
```

创建者实现
```go

type ComputerBuilderImplement struct {
	ComputerItems ComputerItems
}

func (c ComputerBuilderImplement) SetCpu(cpu string) ComputerBuilder {
	c.ComputerItems.Cpu = cpu
	return c
}

func (c ComputerBuilderImplement) SetMemory(memory string) ComputerBuilder {
	c.ComputerItems.Memory = memory
	return c
}

func (c ComputerBuilderImplement) SetHardDisk(hardDisk string) ComputerBuilder {
	c.ComputerItems.HardDisk = hardDisk
	return c
}

func (c ComputerBuilderImplement) SetScreenSize(screenSize int) ComputerBuilder {
	c.ComputerItems.ScreenSize = screenSize
	return c
}

func (c ComputerBuilderImplement) GetComputerProduct() ComputerItems {
	return c.ComputerItems
}

func (c ComputerBuilderImplement) SetConstruct(construct Construct) ComputerBuilder {
	c.ComputerItems.Construct = construct
	return c
}

func (c ComputerBuilderImplement) Build() ComputerItems {
	return c.ComputerItems.Construct.BuildComputer(c.ComputerItems)
}

```

Construct接口
```go
type Construct interface {
	BuildComputer(ComputerItems ComputerItems) ComputerItems
}
```

台式机Construct实现
```go
type ComputerConstruct struct {
	ComputerBuilder ComputerBuilder
}

func (con ComputerConstruct) BuildComputer(ComputerItems ComputerItems) ComputerItems {
	ComputerItems.ComputerProduct = "台式电脑：" +
		ComputerItems.Cpu + " " +
		ComputerItems.Memory + " " +
		ComputerItems.HardDisk + " " +
		strconv.Itoa(ComputerItems.ScreenSize) + "寸显示屏"
	return ComputerItems
}
```

笔记本Construct实现，可以看到和上面的台式电脑在构建显示屏的时候有差异
```go
type LaptopConstruct struct {
	ComputerBuilder ComputerBuilder
}

func (con LaptopConstruct) BuildComputer(ComputerItems ComputerItems) ComputerItems {
	ComputerItems.ComputerProduct = "笔记本电脑：" +
		ComputerItems.Cpu + " " +
		ComputerItems.Memory + " " +
		ComputerItems.HardDisk + " " +
		strconv.Itoa(ComputerItems.ScreenSize) + "寸触摸显示屏"
	return ComputerItems
}
```

