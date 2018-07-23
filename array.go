package main

import "fmt"

func main() {
	//元素的类型和长度都是数组类型的一部分,数组默认是0值,对于int数组就是0
	var arr [5]int
	fmt.Println("array:", arr)

	//设置/获取数组某位的值
	arr[2] = 100
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[2])
	fmt.Println("len:", len(arr))

	//一行内初始化一个数组
	b := [5]int{13, 90, 28, 67, 34}
	fmt.Println("dcl:", b)

	//多纬数组
	var twoA [2][3]int
	for i := 0;i < 2;i++ {
		for j := 0;j < 3;j++ {
			twoA[i][j] = i * j
		}
	}
	fmt.Println("2Array: ", twoA)
	fmt.Println("2Array len:", len(twoA))
}
