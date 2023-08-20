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

	arr := make([]int, 8)
	for i := 0; i < 8; i++ {
		var num int
		fmt.Fscan(r, &num)
		arr[i] = num
	}

	ans := true

	// 1
	small := arr[0]
	for _, v := range arr {
		if small <= v {
			small = v
		} else {
			ans = false
		}
	}

	// 2
	for _, v := range arr {
		if v >= 100 && v <= 675 {
		} else {
			ans = false
		}
	}

	// 3
	for _, v := range arr {
		if v%25 == 0 {
		} else {
			ans = false
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
