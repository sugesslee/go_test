package main

import "fmt"

type Rect struct {
	x, y           float64 //类只包含属性，没有方法
	weight, height float64
}

// 为Rect类型绑定Area的方法，*Rect为指针引用可以修改传入参数的值
func (r *Rect) Area() float64 {
	// 方法归属于类型，不归属于具体的对象，声明该类型的对象即可调用该类型的方法
	return r.height * r.weight
}

func main() {
	var rect Rect
	rect.weight = 10
	rect.height = 10
	fmt.Println(rect.Area())
}
