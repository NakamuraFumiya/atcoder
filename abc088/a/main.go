package main

import "fmt"

func main() {
	var N, A int
	fmt.Scan(&N)
	fmt.Scan(&A)
	n := N - (500 * (N / 500))
	if n <= A {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
