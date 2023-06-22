package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	name string
	age  uint64
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	persons := []Person{}
	for i := 0; i < n; i++ {
		var name string
		var age uint64
		fmt.Fscan(r, &name, &age)
		persons = append(persons, Person{name, age})
	}

	minIndex := 0
	min := persons[0].age
	for i, v := range persons {
		if v.age < min {
			minIndex = i
		}
	}

	for _, v := range persons[minIndex:] {
		fmt.Println(v.name)
	}
	for _, v := range persons[:minIndex] {
		fmt.Println(v.name)
	}
}
