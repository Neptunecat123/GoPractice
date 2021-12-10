# 00

* GOROOT: Go语言安装根目录的路径
* GOPATH: （若干）工作区目录的路径，自己定义的工作空间
* GOBIN: GO程序生成的可执行文件路径 

Go语言的源码是以代码包为基本组织单位，在文件系统中，这些代码包其实是与目录一一对应的，由于目录可以有子目录，所以代码包也可以有子包。代码包的名称一般会与源码文件所在目录同名，如果不同名，那么在构建、安装的过程中会以代码包名称为准。

每个代码包都会有导入路径，代码包的导入路径是其他代码在该包中的程序实体时需要引入的路径。在工作区中，一个代码包的导入路径实际上就是从src子目录，到该包的实际存储位置的相对路径。

    import "github.com/labstack/echo"

源码文件放在src子目录下；安装后产生的归档文件（以".a"为扩展名的文件），就会放进该工作区的pkg子目录；如果产生了可执行文件，会放在该工作区的bin子目录。

源码文件会以代码包的形式组织起来，一个代码包其实就对应一个目录，安装某个代码包而产生的归档文件是与这个代码包同名的。放置的相对目录就是该代码包的导入路径的直接父级。执行命令：

构建和安装：构建使用命令 go build，安装使用go install，构建和安装代码包的时候都会执行编译、打包等操作，这些操作生成的任何文件都会被保存到某个临时目录中。如果构建的是库源码文件，那么操作后产生的结果文件只会存在于临时目录中，这里构建的意义在于检查和验证。如果构建的是命令源码文件，那么操作的结果文件会被搬运到源码文件所在的目录中，安装会先执行构建，然后还会进行链接操作，并把结果文件搬运到指定目录。

如果安装的是库源码文件，那么结果文件会被搬运到它所在工作区的pkg目录下的某个子目录中。如果安装的是命令源码文件，那么结果文件会被搬运到他所在工作区的bin目录中，或者环境变量GOBIN指向的目录中。

## 01 常量和变量

### 变量

#### 变量的声明

    var v1 int
    var v2 string
    var v3 []int
    var v4 struct {
        f int
    }
    var v5 *int
    var v6 map[string]int
    var v7 func(a int) int

    var (
        v1 int
        v2 string
    )

#### 变量初始化

    var v1 int = 10
    var v2 = 10
    v3 := 10 //左侧的变量不能是被声明过的

`:=`声明的变量只能在函数里，`var`声明的变量可以在函数外作为全局变量

#### 变量赋值

    var v1 int
    v1 = 10 //赋值

    i, j = j, i //变量交换

#### 匿名变量

    func GetName() (firstName, lastName, nickname string) {
        return "May", "Chan", "Chibi Maruko"
    }
    _, _, nickName := GetName() //_代替不需要的变量，不用定义没用的变量

### 常量

#### 字面常量

程序中的硬编码

    -12 
    3.1415  //浮点
    3.2+12i //复数
    true    //布尔
    "foo"   //字符串

#### 常量的定义

    const Pi float64 = 3.1415926
    const zero = 0.0
    const (
        size int64 = 1024
        eof = -1
    )

    const u, v float32 = 0, 3 //u=0.0, v=3.0
    const a, b, c = 3, 4, "foo"

const用于定义不会改变的数据，常量是在编译时被创建的，即使定义在函数内部也是如此。常量的值必须是能够在编译时就能够确定的，可以在其赋值表达式中涉及计算过程，但是所有用于计算的值必须在编译期间就能获得。

#### 预定义常量

预定义常量：`true` `false` `iota`

**iota**

在每一个const关键字出现时被重置为0，然后再下一个`const`出现之前，每出现一次`iota`，其所代表的数字自增1

    const (
        c0 = iota // 0
        c1 = iota // 1
        c2 = iota // 2
    )
    const (
        c3 = 1 << iota // iota在const开头被置为0，左移0位，c3 = 1
        c4 = 1 << iota // iota为1，c4 = 2
        c5 = 1 << iota // iota为2，c5 = 4
    )
    const (
        c6 = iota*42         // c6 = 0
        c7 float64 = iota*42 // c7 = 42.0
    )
    const c8 = iota // c8 = 0
    const c9 = iota // c9 = 0

如果const赋值语句的表达式是一样的，可以省略后面的赋值表达式：

    const (
        a = iota // a = 0
        b        // b = 1
        c        // c = 2
        d        // d = 3
    )

**枚举**

    const (
        Sunday = iota
        Monday
        Tuesday
        Wednesday
        Thursday
        Friday
        Saturday
        numberOfDays
    )

大写字母开头的常量包外可见，`numberOfDays`为包内私有。


## 02 数据类型

### 基础类型

* 布尔：bool

    布尔类型不接受其他类型的赋值，不支持自动或强制类型转换。

* 整型：int8, byte, int16, int, uint, uintptr

    |类型|字节|值范围|
    |----|----|----|
    |int8|1|-128-127
    |uint8|1|0-255|
    |int16|2|-32768-32767|
    |uint16|2|0-65535|
    |int32|4||
    |uint32|4||
    |int64|8||
    |uint64|8||
    |int|平台相关||
    |uint|平台相关||
    |uintptr|同指针|32位平台下4字节，64位平台下8字节|

    `int`和`int32`是不同类型，需要强制类型转换。

* 浮点：float32, float64

        fv := 12.0 // 默认fv是float64

* 复数：complex64, complex128
* 字符串：string

        var str string

        str = "Hello World."

    字符串内容不能再初始化后被修改。

        "Hello" + "123" // Hello123
        len("Hello") // 5
        "Hello"[1] // 'e'

    遍历字符串（含中文），如果按照字节数组的方式遍历，中文在utf8中占3个字节，每个字节的类型是`byte`。

    以unicode字符方式遍历时，1个中文算1一个unicode字符，每个字符的类型是`rune`

* 字符：rune

    rune等同与int32，用来处理代表unicode字符

    byte和rune之间可以转换，byte转向rune时不会出错，但是如果rune表示的字符只占用一个字符，不超过 uint8 时不会出错；超过时直接转换编译无法通过，可以通过引用转换，但是会舍去超出的位，出现错误结果

* 错误：error

### 复合类型

* 指针：pointer
* 数组：array

    [32]byte
    [2*N] struct {x, y int32}
    [1000] *float64


        for i, v := range array {
            fmt.Println("Array element[", i "]=", v)
        }

    range有两个返回值，第一个返回值是元素的数组下标，第二个返回值是元素的值。

* 切片：slice

    切片的数据结构可抽象为以下3个变量：

    1. 一个指向原生数组的指针
    2. 数组切片中元素的个数
    3. 数组切片已分配的存储空间

    创建数组切片：

    1. 基于数组
    
      myslice = myArray[first:last]
      mySlice = myArray[:5] // 前5个（0-4）
      mySlice = myArray[5:] // 第5个到最后包含第5个


    2. 直接创建

      mySlice1 := make([]int 5)
      mySlice2 := make([]int 5, 10) // 元素个数为5，初始值为0，预留10个元素的存储空间
      mySlice3 := []int{1,2,3,4,5}

    动态增减数组

    capacity

    to be continue

    内容复制

    copy()

    to be continue

* 字典：map

      var myMap map[string] PersonInfo // 变量声明。string是键的类型，PersonInfo是其中存放值的类型

      myMap = make(map[string] PersonInfo, 100) //初始存储能力为100的map

      myMap["Jack"] = PersonInfo{"1", "Jack", "Room 101"} // 赋值

      delete(myMap, "Jack") // 删除元素

      // 元素查找
      val, ok := myMap["Mary"]
      if ok {
          // 处理value
      }

* 通道：chan

    to be continue

* 结构体：struct

    to be continue

* 接口：interface

    to be continue

## 03 流程控制

### switch

    name := "Frank"
	switch name {
	case "Tom":
		fmt.Println("hahahah")
	case "Jim":
		fmt.Println("hehehehehe")
	case "Frank":
		fallthrough
	case "mary":
		fmt.Println("it's mary")
	case "Bob":
		fmt.Println("nonononon")
	}

    //output: it's mary

* Golang不需要用break来明确退出一个case
* 明确添加fallthrough关键字，会继续执行紧跟的下一个case

### 循环

不支持while和do while，只支持for

## 04 函数

### 自定义包需要调用的包以及函数

### 不定参数

函数传入的参数个数不确定。

    func myfunc(arg ...int) {
        for _, arg := range arg {
            fmt.Println(arg)
        }
    }

用interface{}传递任意类型数据。

    func MyPrintf(args ...interface{}) { 
        for _, arg := range args { 
            switch arg.(type) { 
            case int: 
            fmt.Println(arg, "is an int value.") 
            case string: 
            fmt.Println(arg, "is a string value.") 
            case int64: 
            fmt.Println(arg, "is an int64 value.") 
            default: 
            fmt.Println(arg, "is an unknown type.") 
            } 
        } 
    }

### 匿名函数，闭包



## 05 go常用命令

go build

go run

go get

go install

## 06 包管理

## 07 GC？反射？内存管理？

## 08 并发编程

协程（Coroutine）本质是一种用户态线程，Go语言在语言级别支持轻量级线程，叫goroutine，Golang标准库提供的所有系统调用操作（包括所有同步IO操作），都会让出CPU给其他goroutine。

### goroutine

goroutine是Golang中轻量级线程的实现，在函数调用前加上`go`关键字，本次调用就会在一个新的goroutine中并发执行。当被调用的函数返回时，goroutine也自动结束了，如果这个函数有返回值，那么这个返回值会被丢弃。
 
### 并发通信

两种常见的并发通信模型：共享数据和消息，Golang使用消息机制作为通信方式。

channel是Golang提供的goroutine之间的通信方式，可以在两个或多个goroutine之间传递消息。channel是进程内的通信方式，因此通过channel传递对象的过程和调用函数时的参数传递行为比较一致。

如果要跨进程通信，建议用分布式系统的方法解决，如使用Socket或者HTTP等通信协议。



## 09 context？CGO？unsafe

## 10 逃逸分析？GC优化？sync.pool

## 11 内存分配模型

## 12 常用第三方包

### excelize

## 13 常用包

### 13.1 fmt

### 13.2 strings

strings包提供utf-8类型字符串的基本操作。

#### 13.2.1 Compare

#### 13.2.2 Contains

`func Contains(s, substr string) bool`

功能：字符串s中是否包含子串substr，返回bool值。

	fmt.Println(strings.Contains("seafood", "foo"))  // true
	fmt.Println(strings.Contains("seafood", "bar"))  // false
	fmt.Println(strings.Contains("seafood", ""))     // true
	fmt.Println(strings.Contains("", ""))            // true

(ContainsAny, ContainsRune)

#### 13.2.3 Count

`func Count(s, str string) int`

功能：计算字符串str在字符串s中出现的非重叠次数

    fmt.Println(strings.Count("ffffff", "f")) // 6
    fmt.Println(strings.Count("ffffff", "fff")) // 2
    fmt.Println(strings.Count("ffffff", "ffff")) // 1
    fmt.Println(strings.Count("ffffff", "ffffff")) // 1
    fmt.Println(strings.Count("ffffff", "fffffffff")) // 0
    fmt.Println(strings.Count("ffffff", "a")) // 0

#### 13.2.4 Repeat

`func Repeat(s string, count int) string`

功能：重复s字符串count次，返回重复后的新字符串

    var s string = "Hello"
	fmt.Println(strings.Repeat(s, 3)) // HelloHelloHello

#### 13.2.5 Split,Fields

`func Fields(s string) []string`

功能：用1个或多个空白符号作为动态长度的分隔符将字符串分割成若干块，返回一个slice。如果字符串只包含空白字符，则返回一个长度为0的slice。

`func Split(s, seq string) []string`

功能：把s字符串用seq分割，返回slice

    var ip string = "192.168.13.1"
	var test string = "haha hehe heihei"
	fmt.Println(strings.Fields(test))   // [haha hehe heihei]
	fmt.Println(strings.Split(ip, ".")) // [192 168 13 1]


#### EqualFold

#### HasPrefix

#### HasSuffix

#### Index

IndexAny, IndexByte, IndexFunc, IndexRune

#### Join

#### LastIndex

#### Map

#### Replace

### io

### bufio

### strconv

### os

### sync

### flag

### encoding/json

### http


## 14 框架：gin

## 15 kratos




