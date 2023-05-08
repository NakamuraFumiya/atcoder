package main

import (
	"fmt"
	"strconv"
)

func main() {
	var r, g, b uint8
	fmt.Scanf("%d %d %d", &r, &g, &b)
	num, _ := strconv.Atoi(fmt.Sprintf("%d%d%d", r, g, b))
	if num%4 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
