### 单例模式（Singleton Pattern）

##### 创建模式（Creational Pattern）

#### 定义

***Ensures that only one instance of a class is allowed within a system.***

***保证一个类仅有一个实例，并提供一个访问它的全局访问点。***

![singleton Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/singleton_pattern.png)

单例模式在开发中是一个很常用也很重要的模式，比如设计一个简单的缓存，一般来说整个系统进程中尽量只能有一个缓存实例存在（如果非要有多个缓存示例存在就要考虑多个缓存的同步问题），那么这个缓存就是单例的。

还有个例子，就是多方同时编辑一个文件，那么该文件在服务端内存中也最好是单独的一份，不然会造成两边同时编辑的差异问题。

### 优点
1. 闲置了内存中只有一个实例，降低了内存的存储开销和创建开销，尤其是当开销成本很大时，单例模式的优势就非常明显。
2. 避免对资源或者对象的多重占用（比如写文件操作，只有一个实例时，避免对同一个资源文件同时写操作），简单来说就是对唯一实例的受控访问。

### 例子分析

##### 一. 各种单例例子
饿汉模式，在加载时就创建，这样的缺点是如果单例一直用不到，会造成内存的开销。

```go
type eagerSingleton struct{}

var eagerSingletonInstance *eagerSingleton = &eagerSingleton{}

func GetEagerSingletonInstance() *eagerSingleton {
	return eagerSingletonInstance
}
```

懒汉模式，也就是懒加载，在需要时候再创建，但是缺点是如果有两个或者两个以上线程同时请求，有可能会因为线程不安全造成多个示例产生，单例失败。
```go
type lazySingleton struct{}

var lazySingletonInstance *lazySingleton

func GetLazySingletonInstance() *lazySingleton {
	if lazySingletonInstance == nil {
		lazySingletonInstance = &lazySingleton{}
	}
	return lazySingletonInstance
}
```

带锁的懒汉模式，这样避免了上面懒汉模式的现成不安全问题，但是缺点是获取单例实例的时候都要加锁，这样还有锁开销。
```go
type lazySingletonWithLock struct{}

var lazySingletonWithLockInstance *lazySingletonWithLock
var mu sync.Mutex

func GetLazySingletonWithLockInstance() *lazySingletonWithLock {
	mu.Lock()
	defer mu.Unlock()

	if lazySingletonWithLockInstance == nil {
		lazySingletonWithLockInstance = &lazySingletonWithLock{}
	}
	return lazySingletonWithLockInstance
}
```

带判断的线程安全的模式，避免了每次加锁
```go
type threadSafeSingleton struct{}

var threadSafeInstance *threadSafeSingleton
var mu1 sync.Mutex

func GetThreadSafeInstance() *threadSafeSingleton {
	if threadSafeInstance == nil {
		mu1.Lock()
		defer mu1.Unlock()
		if threadSafeInstance == nil {
			threadSafeInstance = &threadSafeSingleton{}
		}
	}
	return threadSafeInstance
}
```

基于sync.Once的单例模式
```go
type syncOnceSingleton struct{}

var syncOnceInstance *syncOnceSingleton
var once sync.Once

func GetSyncOnceInstance() *syncOnceSingleton {
	once.Do(func() {
		syncOnceInstance = &syncOnceSingleton{}
	})
	return syncOnceInstance
}
```


### 参考文档：

https://studygolang.com/articles/11444

https://www.journaldev.com/1377/java-singleton-design-pattern-best-practices-examples