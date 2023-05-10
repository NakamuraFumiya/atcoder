package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	count := 700
	if fmt.Sprintf("%c", S[0]) == "o" {
		count = count + 100
	}
	if fmt.Sprintf("%c", S[1]) == "o" {
		count = count + 100
	}
	if fmt.Sprintf("%c", S[2]) == "o" {
		count = count + 100
	}
	fmt.Println(count)
}
