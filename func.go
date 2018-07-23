package main

import "fmt"

//func定义
func plus(x int, y int) int {
	return x + y
}

func vals() (int, int) { //多返回值
	return 3, 7
}

//变参函数
func sum(nums ...int) { //使用任意数目的int作参数
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

//closure闭包,通过闭包使用匿名函数
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1) //直到fact(0)一直调用自身
}

func main() {
	res := plus(10, 90) //func调用
	fmt.Println("10+90 =", res)

	fmt.Println("------")

	//多赋值操作使用这两个不同返回值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//使用空白定义符_，返回部分返回值
	_, c := vals()
	fmt.Println(c)

	fmt.Println("------")

	sum(9)
	sum(1, 2)
	sum(1, 2, 3)
	
	nums := []int{1, 2, 3, 4}
	sum(nums...) //如果slice已有多个值,想把它们作变参使用,func(slice...)

	fmt.Println("*********")

	nextInt := intSeq()
	fmt.Println(nextInt()) //每次调用nextInt都会更新i的值
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println("+++++++")

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println("--------")

	fmt.Println(fact(10))
}
