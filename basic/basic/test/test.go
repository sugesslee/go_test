package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func main() {
	var s1 = [2]string{"hello", "world"}

	s2 := [...]string{"1", "3", "5"}

	var line1 [2]image.Point

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(line1)
	fmt.Println()

	var a = [...]int{1, 2, 3}
	var b = &a

	fmt.Println(a[0], a[1])
	a[1] = 0
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])
	fmt.Println()

	b[1] = 0
	fmt.Println(b[0], b[1])

	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println()

	// 函数数组
	var decoder1 [2]func(io.Reader) (image.Image, error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}

	fmt.Println(decoder1)
	fmt.Println(decoder2)
	fmt.Println()

	// 接口数组
	var unknow1 [2]interface{}
	var unknow2 = [...]interface{}{123, "您好"}

	fmt.Println(unknow1)
	fmt.Println(unknow2)
	fmt.Println()

	// 通道数组
	var chanList = [2]chan int{}

	fmt.Println(chanList)
	fmt.Println()

	c1 := make(chan [0]int)

	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1

}
