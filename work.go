// work
package main

import (
	"fmt"
	"unsafe"
)

func getMaxValue(a, b int) int {
	var temp int
	if a > b {
		temp = a
	} else {
		temp = b
	}
	return temp
}
func getMaxIndexValue(a *int, b *int) int {
	var temp int
	if *a > *b {
		temp = *a
	} else {
		temp = *b
	}
	return temp
}
func swep(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

type animals struct {
	name  string
	age   int
	boots int
}

func main() {
	//	var a = 1
	//	var b = 9
	//	//	var c int
	//	c := a
	//	var d, e = "string", 1
	//	var f string = d
	//	var g = a
	//	var h = "string"
	//	fmt.Println(a+b, c, &d, &h, e, &f, &a, &e, *(&g))

	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)

	const age int = 18
	const width int = 10
	const height int = 10
	var ageStr int = 18
	var widthStr int = 10
	var maxValue int
	maxValue = getMaxValue(age, width)
	fmt.Println(maxValue)
	var maxIndex int
	maxIndex = getMaxIndexValue(&ageStr, &widthStr)

	fmt.Println(maxIndex)
	fmt.Println(&ageStr, &widthStr, &maxIndex)
	swep(&ageStr, &widthStr)
	fmt.Println(&ageStr, &widthStr, ageStr, widthStr)
	var area int
	area = width * height
	fmt.Printf("面积为：%d，%d，", width, area)
	fmt.Println(age)

	const (
		i = "123456"
		j = len(i)
		k = unsafe.Sizeof(j)
	)
	fmt.Println(i, j, k)
	//	var l *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int32(0)))))
	var m int32
	var n int64
	var o int
	var p int16
	fmt.Println(unsafe.Sizeof(m))
	fmt.Println(unsafe.Sizeof(n))
	fmt.Println(unsafe.Sizeof(o))
	fmt.Println(unsafe.Sizeof(p))

	const (
		q = iota //0
		r        //1
		s        //2
		aa
		t = "ha" //独立值，iota += 1
		u        //"ha"   iota += 1
		v = 100  //iota +=1
		w        //100  iota +=1
		y = iota //7,恢复计数
		z        //8
	)
	fmt.Println(q, r, s, aa, t, u, v, w, y, z)

	var forDo int = 10
	var fori int
	for fori < forDo {
		fori++
		fmt.Println(fori)
	}
	//	for fori := 0; fori < forDo; fori++ {
	//		fmt.Println(fori)
	//	}
	var point *int
	var even int
	even = 10
	point = &even
	fmt.Println(&even, point, "-----", *point, &point)
	var bird animals
	var dark animals
	//	var monkey animals

	bird.name = "大鹏"
	bird.age = 18
	bird.boots = 2

	dark.name = "鸭子"
	dark.age = 16
	dark.boots = 2
	fmt.Println("有一只", bird.name, "它已经", bird.age, "岁了，", "它有", bird.boots, "只脚")
	fmt.Println("有一只", dark.name, "它已经", dark.age, "岁了，", "它有", dark.boots, "只脚")
	getAnimal(&dark)
}
func getAnimal(animal *animals) {
	fmt.Println(animal.name)
}
