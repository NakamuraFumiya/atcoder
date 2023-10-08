package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var s string
	fmt.Fscan(r, &s)

	yesFlg := true
	ss := strings.Split(s, "")
	for i, v := range ss {
		i++
		if i <= 1 {
			continue
		}

		if i%2 == 0 {
			if v == "1" {
				yesFlg = false
			}
		}
	}

	if yesFlg {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
