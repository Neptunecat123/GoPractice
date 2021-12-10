package main

import (
	"fmt"
	"unsafe"
)

// 自定义类型
type (
	IZ int
	FZ float32
	STR string
)
// 类型别名
type IA = int

// 可以给自定义类型添加方法
func (iz IZ) sum(i int) {
	fmt.Println(int(iz) + i)
}

func main() {
	// 基本类型
	// bool
	v1 := true
	v2 := (1 != 0)
	fmt.Println(v1, v2)

	// int
	var i1 int = 1 // 本机环境int的长度和int64一样
	var i2 int8 = 1
	var i3 int16 = 1
	var i4 int32 = 1
	var i5 int64 = 1
	fmt.Println(unsafe.Sizeof(i1))
	fmt.Println(unsafe.Sizeof(i2))
	fmt.Println(unsafe.Sizeof(i3))
	fmt.Println(unsafe.Sizeof(i4))
	fmt.Println(unsafe.Sizeof(i5))

	// string
	str := "Hello world."
	ch := str[0]
	fmt.Printf("str is \"%s\". len(str) is %d. ch is %c\n", str, len(str), ch)
	str2 := "123"
	fmt.Println(str + str2)

	// byte, rune
	//[]rune(str), 它可以将字符串转化成 unicode 
	str = "hello 世界"
	r := []rune(str)
	fmt.Println([]byte(str))
	fmt.Println([]rune(str)) // 一个中文字符长度是3个字节
	fmt.Println(string(r[6:])) // rune可以按照unicode字符输出字符


	// 自定义类型, 系统会判定自定义类型和原类型不一样
	var iz IZ = 10
	var i int = 10
	var ia IA = 10
	fmt.Println(iz)
	fmt.Printf("iz IZ %T %v\n", iz, iz)
	fmt.Printf("i int %T %v\n", i, i)
	fmt.Printf("ia int %T %v\n", ia, ia)

	// 类型别名可以与同类型进行操作；自定义类型不行，自定义类型与原类型相当于不同类型
	fmt.Println(i + ia)

	// 类型强转
	fmt.Println(int(iz))
	f := 1.23
	fmt.Println(int(f))

	// 调用自定义类型的方法
	iz.sum(5)

	// const
	// const常量只能定义boolean, number (integer, float or complex) or string类型
	const c1 = "string"
	const c2 = false
	fmt.Println(c1)
	fmt.Println(c2)

	// const实现枚举
	const (
		sun = iota //从0开始，后面依次递增1
		mon
		tue
		wed
		thu
		fri
		sat
	)
	fmt.Println(mon, tue, wed, thu, fri, sat, sun)

	// array
	a := [5]int{1, 2, 3, 4, 5}
	a2 := [...]string{"hello", "world", "!"}
	fmt.Println("array values:", a)

	for i := 0; i < len(a2); i++ {
		fmt.Println(a2[i])
	}

	a[0] = 10
	fmt.Println("array2 values:", a)

	// slice
  // 创建slice方法1:从数组中切片
	sli1 := a[1:4] 
	for i:=0; i<len(sli1); i++ {
		fmt.Printf("%d ", sli1[i])
	}
	fmt.Println()

	// 创建slice方法2:make([]type, len, cap)
	sli2 := make([]int, 10, 20) 
	for i:=0; i<len(sli2); i++ {
		sli2[i] = 5 * i
	}
	fmt.Println(sli2)

	// copy and append
	sl_to := make([]int, 15)
	n := copy(sl_to, sli2)
	fmt.Println(sl_to)
	fmt.Printf("copied %d elements.\n", n)

	sli3 := append(sl_to, sli2...) // append一个切片
	sli3 = append(sli3, 6, 6, 6) // append指定几个值
	fmt.Println(sli3)

	st := "abcde"
	fmt.Println(st[len(st)/2:] + st[:len(st)/2])

	// map
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(m["one"])
	m["two"] = 222
	fmt.Println(m)
	
}