package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1s2s3 string
	fmt.Scan(&s1s2s3)
	var count int
	for _, v := range strings.Split(s1s2s3, "") {
		if v == "1" {
			count++
		}
	}
	fmt.Println(count)
}
