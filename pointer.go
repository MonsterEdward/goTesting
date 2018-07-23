package main

import "fmt"

//zeroval有一个int型参数,所以使用值传递,zeroval()将从调用它的那个函数中得到一个ival形参的拷贝
func zeroval(ival int) {
	ival = 0
}

//zeroptr有一个*int参数,意味着它使用了一个int指针,zeroptr()函数体内的*iptr解引用这个指针,从它内存地址得到这个地址对应的当前值.对一个解引用的指针赋值,将会改变这个指针引用的真实地址的值
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i) //zeroval()在main()中不能改变i的值

	zeroptr(&i)
	fmt.Println("zeroptr:", i) //zeroptr()可在main()中改变i的值,因为有一个这个变量的内存地址的引用

	fmt.Println("pointer:", &i) //通过&i取得i的内存地址
}
