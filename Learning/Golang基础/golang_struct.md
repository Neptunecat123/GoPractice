# Golang struct

## tag用法

struct tag是结构体中某个字段别名，可以定义多个，空格分隔。

    type Student struct {
        Name string `ak:"av" bk:"bv" ck:"cv"`
    }

tag的作用相当于该字段的一个属性标签，一些包通过tag做相应判断。例如

    type Student struct {
        Name string `json:"name"`
    }

    s1 := Student {Name: "s1"}

    v, err := json.Marshal(s1)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(v)) // []byte转string，json格式

string(v)的结果，"Name"转成"name"

    {
        "name": "s1"
    }

## 常用tag

* json json序列化和反序列化时字段的名称
* db sqlx模块中对应的数据库字段名
* form gin框架中对应的前端数据字段名
* binding 搭配form使用，如果没查到结构体中某个字段不报错值为空，binding为required代表没找到返回错误给前端