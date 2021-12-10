# Golang面向对象编程

Is Go an object-oriented language?

**Yes and no**. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concecpt of "interface" in Go provides a different approach that we believe is esay to use and in some ways more general.

Also, the lack of a type hierarchy makes "objects" in Go feel much more lightweight than in languages such as C++ or Java.

## 封装数据和行为

### 封装数据

结构体定义

```go
type Employee struct {
    Id string
    Name string
    Age int
}
```

初始化

```go
e := Employee{"001", "Bob", 20}
e1 := Employee{Name:"Mike", Age:30}
e2 := new(Employee) //new这里返回的是引用（指针），相当于e:=&Employee{}
e2.Id = "003"
e2.Age = 22
e2.Name = "Rose"
```

### 封装行为

```go
// 实例对应的方法被调用时，实例的成员会进行值复制，内存地址不一致
func (e Employee) String() string {
    return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
}

// 通常情况下避免内存拷贝，使用这种方式，内存地址一致
func (e *Employee) String() string {
    return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
} 
```

## 接口相关

对象之间交互协议

Go接口为非入侵性，实现不依赖于接口定义，接口的定义可以包含在接口使用者包内。

接口变量

自定义类型

```go
type FuncInt func (int op) int //  定义一个类型FuncInt，这个类型是个函数，需要一个int入参，返回一个int
```

## 扩展


## 多态


## 空接口

1. 空接口可以表示任何类型。
2. 通过断言【.()】来将空接口转换为指定类型

```go
v, ok := p.(int) // ok==true时转换成功
```