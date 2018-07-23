package main

import "fmt"

func main() {
	//迭代slice
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {//不需索引时,使用空值定义符_来忽略
		sum += num
	}
	
	fmt.Println("sum:", sum)

	fmt.Println("-----------")

	for i, num := range nums {//需要索引i
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	fmt.Println("----------")

	//迭代map
	kvs := map[string]string{"a": "Apple", "b": "Banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	fmt.Println("-----------")

	//迭代字符串(unicode编码),第1个返回值为rune起始字节位置,第2个是rune自己
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

