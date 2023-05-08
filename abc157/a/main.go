package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	if N%2 == 0 {
		fmt.Printf("%d\n", N/2)
	} else {
		fmt.Printf("%d\n", N/2+1)
	}
}
