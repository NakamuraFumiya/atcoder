package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var s string
	fmt.Fscan(r, &s)

	aIndex, zIndex := 0, 0
	for i := range s {
		if string(s[i]) == "A" {
			aIndex = i
			break
		}
	}
	for i := range s {
		if string(s[i]) == "Z" {
			zIndex = i
		}
	}
	fmt.Println(zIndex - aIndex + 1)
}
