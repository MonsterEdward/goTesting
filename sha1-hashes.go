package main

import(
	"crypto/sha1"
	"fmt"
	"crypto/md5"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	s1 := "sha1 this string"
	m := md5.New()
	m.Write([]byte(s1))
	ms := m.Sum(nil)
	fmt.Println(s1)
	fmt.Println(ms)
}
