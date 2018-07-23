package main

import "fmt"

func main () {//关联数组,(哈希/字典)
	m := make(map[string]int)
	
	m["k1"] = 6
	m["k2"] = 7
	fmt.Println("map:", m)

	v1 := m["k1"]
	v2 := m["k2"]
	fmt.Println("v1: ", v1)
	fmt.Println("v2: ", v2)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
