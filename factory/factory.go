package factory

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hi, MyName is %s", p.Name)
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// 抽象工厂
type Student interface {
	GreetStudent()
}

type student struct {
	name string
	age  int
}

func (s student) GreetStudent() {
	fmt.Printf("Hi! My name is %s", s.name)
}

func newStudent(name string, age int) Student {
	return student{name: name, age: age}
}
