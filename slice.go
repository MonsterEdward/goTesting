package main

import "fmt"

func main() {
	//slice的类型仅由它所包含的元素决定(不像array还需元素的个数)
	s := make([]string, 3)	//要创建一个长度非0的slice需使用内建方法make
	fmt.Println("emp:", s, len(s))

	s[0] = "P"
	s[1] = "I"
	s[2] = "H"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "J")	//内建append方法
	s = append(s, "V", "Z")
	fmt.Println("apd:", s, len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]	//s[2]开始,总共5-2=3个元素
	fmt.Println("sl1:", l)

	l = s[:4] //s[0]开始,总4-0=4个元素
	fmt.Println("sl2:", l)

	l = s[4:] //s[4]开始,总共len(s) - 4 = 2个元素
	fmt.Println("sl3:", l)

	t := []string{"3", "7", "9"} //一行内声明slice
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3) //多纬slice
	for i := 0;i < 3;i++ {
		innerLen := i + 1 //内部长度可不同
		twoD[i] = make([]int, innerLen)
		for j := 0;j < innerLen;j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
