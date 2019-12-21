### 代理模式（Proxy Pattern）

##### 结构模式（Structural Pattern）

#### 定义

***Allows for object level access control by acting as a pass through entity or a placeholder object.***

***为其他对象提供一种代理以控制对这个对象的访问。***

![Proxy Pattern UML](../images/proxy_pattern.png)

代理模式在程序编写中是一种非常常见的模式。

代理模式的目标是通过代理类来隔离访问者对目标类的直接访问，访问者只能访问代理类，这样能够保证访问者对目标类的访问是受限的。

比如访问某些资源，如果访问者不具备访问的权限或条件，则代理类可以拒绝访问者的访问。这是常见实现权限的方式之一。

### 优点
1. 职责清晰, 业务代码专注于实现实际的业务逻辑，不用关心其他非本职责的事务，通过后期的代理完成一件完整事务，附带的结果就是编程简洁清晰。
2. 代理对象可以在客户端和目标对象之间起到中介的作用，这样起到了中介的作用和保护了目标对象的作用。

### 例子分析

##### 文件读取案例

1. 有一个文件读取的代码，能够读取磁盘上的文件。
2. 但是需要在读取之前判断文件读取者是否具有相应的权限，如果权限不够则拒绝访问。
3. 示例代码

文件访问接口
```go
type FileAccessor interface {
	ReadFile(filePath string)
}
```

文件访问业务代码
```go
type fileAccessorImpl struct {
}

func (*fileAccessorImpl) ReadFile(filePath string) {
	fmt.Println("访问文件操作")
}
```

代理文件访问
```go
type FileAccessorProxy struct {
	UserType string
}

func (f *FileAccessorProxy) ReadFile(filePath string) {
	if f.UserType == "user" {
		fmt.Printf("当前用户: %s \n", f.UserType)
		fileAccessor := fileAccessorImpl{}
		fileAccessor.ReadFile(filePath)
	} else if f.UserType == "guest" {
		fmt.Printf("当前用户: %s \n", f.UserType)
		fmt.Println("对不起，你没有访问权限")
	}
}
```



### 参考文档：
