package main

import (
	"fmt"
)

// 定义结构体
type Person struct {
	Name string
	Age int
	Gender string
	Id string
}

// Student继承Person
type Student struct {
	Grade int
	Person
}

func (p *Person) GetPersonInfo() string {
	Info := fmt.Sprintf("Name: %s, Age: %d, Gender: %s, Id: %s", p.Name, p.Age, p.Gender, p.Id)
  return Info
}

func main() {
	p := new(Person)
	p.Name = "Mike"
	p.Age = 20
	p.Gender = "M"

	fmt.Println(p)
	fmt.Println(p.GetPersonInfo())

	stu := &Student{3, Person{"Mary", 21, "F", "61011300000000"}}
	fmt.Println(stu.GetPersonInfo())
}