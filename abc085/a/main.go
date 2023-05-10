package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	var s string
	for i := range S {
		if i == 3 {
			s = s + "8"
		} else {
			s = s + fmt.Sprintf("%c", S[i])
		}
	}
	fmt.Println(s)
}
