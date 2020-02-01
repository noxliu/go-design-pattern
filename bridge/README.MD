### 适配器模式（Adapter Pattern）

##### 结构模式（Structural Pattern）

#### 定义

***Permits classes with disparate interfaces to work together by creating a common object by which they may communicate and interact.***

***将一个类的接口转换成客户希望的另外一个接口。使原本由于接口不兼容而不能一起工作的那些类可以一起工作。***

![Memento Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/adapter_pattern.png)

目标角色（target）：这是客户锁期待的接口。目标可以是具体的或抽象的类，也可以是接口

适配者角色（adaptee）：已有接口，但是和客户器期待的接口不兼容。

适配器角色（adapter）：将已有接口转换成目标接口。

### 优点
将适配器和适配者类解耦, 避免了因为要适配某个借口，直接修改适配者源码。


### 例子分析

##### 打印工具案例
1. 目前有一打印工具，只能打印字符串类型数据。
2. 但目前除一字符串生成器以外，还有另外两个生成器，分别生成int类型和int数组，这两种生成器生成的内容无法匹配打印工具，无法打印。

打印工具代码如下，可以看到，只有一个StringPrinter方法，其无法接受stringGeneratorI接口以外类型：
```go
type StringPrinterI interface {
	StringPrinter(string string)
}

type StringPrinter struct {
}

func (*StringPrinter) PrintString(string stringGeneratorI) {
	fmt.Println(string.stringGenerator())
}
```

stringGenerator代码如下：
```go
type stringGeneratorI interface {
	stringGenerator() string
}

type StringGenerator struct {
}

func (*StringGenerator) stringGenerator() string {
	return "Hello world"
}
```

intGenerator代码如下：
```go
type IntGeneratorI interface {
	IntGenerator() int
}

type IntGenerator struct {
}

func (*IntGenerator) IntGenerator() int {
	return 19982727
}
```

intArrayGenerator代码如下：
```go
type GeneratorForIntArrayI interface {
	IntArrayGenerator() int
}

type GeneratorForIntArray struct {
}

func (*GeneratorForIntArray) IntArrayGenerator() []int {
	array1 := []int{1, 2, 3, 999, 877777, 1212, 221212}
	return array1
}
```

intGenerator的适配器intToStringAdapter代码如下：
```go
type IntToStringAdapter struct {
	IntValue int
}

func (i *IntToStringAdapter) stringGenerator() string {
	return strconv.Itoa(i.IntValue)
}
```


### 参考文档：
https://en.wikipedia.org/wiki/Adapter_pattern