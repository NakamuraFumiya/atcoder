package main

import (
	"fmt"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N)
	fmt.Scan(&S)

	t, a := 0, 0
	for i := 0; i < N; i++ {
		if string(S[i]) == "T" {
			t++
		} else if string(S[i]) == "A" {
			a++
		}
	}
	if t > a {
		fmt.Println("T")
	} else if t < a {
		fmt.Println("A")
	} else {
		count := t
		t, a = 0, 0
		for i := 0; i < N; i++ {
			if string(S[i]) == "T" {
				t++
			} else if string(S[i]) == "A" {
				a++
			}
			if t == count {
				fmt.Println("T")
				break
			} else if a == count {
				fmt.Println("A")
				break
			}
		}
	}
}
