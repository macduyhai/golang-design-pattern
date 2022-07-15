package main

import (
	"fmt"
	"gotest/design_pattern/Creational-pattern/singleton"
)

func main() {
	s := singleton.NewInstance()
	fmt.Printf("%p\n", &s)
	// fmt.Println(s.AddOne())
	// fmt.Println(s.AddOne())
	s2 := singleton.NewInstance()
	fmt.Printf("%p\n", &s2)
	fmt.Println(s2.AddOne())
}
