package template

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	// 做西红柿
	tomato := &Tomato{}
	doCook(tomato)

	fmt.Println("\n=====> 做另外一道菜")
	// 做炒鸡蛋
	egg := &Egg{}
	doCook(egg)
}
