package main

import (
	"fmt"
)

type Animal interface {
	bark()
}

type Dog struct {
	name string
}

func (d Dog) bark() {
	fmt.Println("wang wang wang")
}

type Cat struct {
	age int
}

func (c Cat) bark() {
	fmt.Println("miao miao miao")
}

type Animals []Animal

func main() {
	animals := Animals{Dog{"dog"}, Cat{3}}
	for _, animal := range animals {
		fmt.Println(animal)
	}
}
