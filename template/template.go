package template

import "fmt"

type Cooker interface {
	fire()
	cooke()
	outfire()
}

// 类似于一个抽象类
type CookMenu struct {
}

func (CookMenu) fire() {
	fmt.Println("开火")
}

// 做菜，交给具体的子类实现
func (CookMenu) cooke() {
}

func (CookMenu) outfire() {
	fmt.Println("关火")
}

// 封装具体步骤
func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outfire()
}

type Tomato struct {
	CookMenu
}

func (*Tomato) cooke() {
	fmt.Println("做西红柿")
}

type Egg struct {
	CookMenu
}

func (Egg) cooke() {
	fmt.Println("做炒鸡蛋")
}
