### 责任链模式（Chain Of Responsibility Pattern）
         
##### 行为模式（Behavioural Pattern）

#### 定义

***Gives more than one object an opportunity to handle a request by linking receiving objects together.***

***每个接收者都包含对另一个接收者的引用。如果一个对象不能处理该请求，那么它会把相同的请求传给下一个接收者，依此类推。***

![Chain_Of_Responsibility Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/chain_of_responsibility_pattern.png)

责任链模式在现实开发中很常见的一个模式，比如我们工作中的上下级处理任务，某个级别没有权限处理，则之后提交到更高一级的处理，以此类推。

### 优点
1. 降低耦合度。将请求的发送者和接收者解耦。 
2. 简化了对象。使得对象不需要知道链的结构。 
3. 增强给对象指派职责的灵活性。通过改变链内的成员或者调动它们的次序，允许动态地新增或者删除责任。 
4. 增加新的请求处理类很方便。

### 例子分析

##### 报销程序案例
1. 报销程序，暂定有三级领导：项目经理project manager, 主管经理 department manager, 总经理 general manager。
2. 项目经理能处理200元以下的报销任务，如果处理不了则转移到更高一级经理，也就是主管经理处处理。
3. 主管经理能处理500元以下的报销任务，如果处理不了则转移到总经理处，总经理能处理大于500元的任务报销任务。
4. 代码如下：

经理接口
```go
type Manager interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}
```

项目经理实现
```go
type ProjectManager struct {
}

func (*ProjectManager) HaveRight(money int) bool {
	return money < 200
}

func (*ProjectManager) HandleFeeRequest(name string, money int) bool {
	fmt.Printf("项目经理处理 %s 的 %d 报销请求\n", name, money)
	return true

}
```

主管经理实现
```go
type DepartmentManager struct {
}

func (*DepartmentManager) HaveRight(money int) bool {
	return money < 500
}

func (*DepartmentManager) HandleFeeRequest(name string, money int) bool {
	fmt.Printf("主管经理处理 %s 的 %d 报销请求\n", name, money)
	return true
}
```

总经理实现
```go
type GeneralManager struct {
}

func (*GeneralManager) HaveRight(money int) bool {
	return true
}

func (*GeneralManager) HandleFeeRequest(name string, money int) bool {
	fmt.Printf("总经理处理 %s 的 %d 报销请求\n", name, money)
	return true
}
```

责任链实现
```go
type RequestChain struct {
	Manager
	successor *RequestChain
}

func (r *RequestChain) SetSuccessor(m *RequestChain) {
	r.successor = m
}

func (r *RequestChain) HandleFeeRequest(name string, money int) bool {
	if r.Manager.HaveRight(money) {
		return r.Manager.HandleFeeRequest(name, money)
	}
	if r.successor != nil {
		return r.successor.HandleFeeRequest(name, money)
	}
	return false
}

func (r *RequestChain) HaveRight(money int) bool {
	return true
}
```

测试代码
```go
func main() {
	projectManager := ProjectManager{}
	projectManagerChain := RequestChain{}
	projectManagerChain.Manager = &projectManager

	departmentManager := DepartmentManager{}
	departmentManagerChain := RequestChain{}
	departmentManagerChain.Manager = &departmentManager

	generalManager := GeneralManager{}
	generalManagerChain := RequestChain{}
	generalManagerChain.Manager = &generalManager

	departmentManagerChain.SetSuccessor(&generalManagerChain)
	projectManagerChain.SetSuccessor(&departmentManagerChain)

	projectManagerChain.HandleFeeRequest("张三", 100)
	projectManagerChain.HandleFeeRequest("张三", 600)
	projectManagerChain.HandleFeeRequest("李四", 499)
	projectManagerChain.HandleFeeRequest("王五", 1099)
	projectManagerChain.HandleFeeRequest("李四", 199)

}
```




### 参考文档：
https://www.runoob.com/design-pattern/chain-of-responsibility-pattern.html