package main

import "fmt"

type person struct { //struct是各字段,字段类型的集合
	name string
	age int
}

type rect struct {
	width, height int
}

func (r *rect) area() int { //area()方法有一个接收器类型rect
	return r.width * r.height
}

func (r rect) perim() int { //可以为值类型或指针类型的接收器定义方法,这是一个值类型接收器的例子
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Println(person{"Bob", 26}) //创建一个新的结构体元素
	
	fmt.Println(person{name: "Alice", age: 30}) //初始化结构体时指定字段值

	fmt.Println(person{name: "Jack"}) //省略的字段被初始化为0

	fmt.Println(&person{name: "Tom", age: 19}) //&生成一个结构体指针

	s := person{name: "Sean", age: 100}
	fmt.Println(s.name) //.访问结构体字段

	sp := &s
	fmt.Println(sp.age) //可对结构体指针使用.指针会被自动解引用

	sp.age = 99
	fmt.Println(sp.age) //结构体可变

	fmt.Println("-----------")

	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area()) //调用上面结构体定义的两个方法
	fmt.Println("perim: ", r.perim())

	fmt.Println("----------")

	rp := &r
	fmt.Println("area: ", rp.area()) //Go自动处理方法调用时的值和指针之间的转化,可以使用指针来调用方法来避免在方法调用时产生一个拷贝,或者让方法能改变接收的数据
	fmt.Println("perim: ", rp.perim())
}
