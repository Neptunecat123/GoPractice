# 函数：一等公民

## 01

* 可以返回多个值
* 所有参数都是值传递（slice，map，channel也是值传递，有引用传递的错觉）
* 函数可以作为变量的值
* 函数可以作为参数和返回值

## 02

可变长参数：不需要指定参数个数，但每个参数类型是一致的。

```go
func sum(ops ...int) int {
    s := 0
    for _, op := range(ops) {
        s += op
    }
    return s
}
```

## 03

defer函数：defer的函数不会立即执行，会在其所在的函数返回前执行defer函数。

```go
func TestDefer(t *testing.T) {
    defer func() {
        t.Log("Clear resources")
    }()
    t.Log("Start")
    panic("Fatal error") // defer的函数会在panic之后仍然执行
}
```
## 04

init函数

* init先于main函数执行
* init不能被其他函数调用
* init没有入参、返回值
* 包的每个源文件可以有多个init函数
* 同一个包的init执行顺序不定，不要依赖；不同包的init函数按照包导入的依赖关系决定执行顺序。


golang程序初始化顺序：

1. 初始化导入的包
2. 初始化包作用域的变量
3. 执行包的init函数



