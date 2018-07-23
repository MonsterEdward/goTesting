package main

import "fmt"

func main() {
	//单个循环条件,最常用
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("-----")

	//初始化/条件/后续,经典
	for j := 7;j <= 9;j++ {
		fmt.Println(j)
	}
	fmt.Println("-----")

	//无条件的for循环,一直执行直到循环体内使用break或return跳出
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("-----")
}
