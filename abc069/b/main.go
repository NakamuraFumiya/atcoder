package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(fmt.Sprintf("%c", s[0]) + fmt.Sprintf("%d", len(s)-2) + fmt.Sprintf("%c", s[len(s)-1]))
}
