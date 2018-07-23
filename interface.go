package main

import "fmt"
import "math"

type geometry interface { //interface是方法特征的命名集合
	area() float64
	perim() float64
}

type rect struct { //让rect和circle实现这个interface
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 { //在Go中实现一个interface,需实现interface中的所有func
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) { //如果一个变量是interface类型,可调用这个被命名的interface中的func,这里的measure()函数,可以用在任何geometry()
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
