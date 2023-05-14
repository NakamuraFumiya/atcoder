package main

import (
	"bufio"
	"fmt"
	"os"
)

func ain() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var s string
	fmt.Fscan(r, &s)

	max, count, count2 := 0, 0, 0
	countedFlg := false
	for _, v := range s {
		if !countedFlg && string(v) == "A" {
			count = 0
			count2 = 0
			count++
			countedFlg = true
		} else if countedFlg && string(v) != "Z" {
			count++
		} else if countedFlg && string(v) == "Z" {
			count++
			if count > max {
				max = count
			}
			countedFlg = false
		} else if !countedFlg && string(v) == "z" {
			if count+count2 > max {
				max = count + count2
			}
		} else if !countedFlg {
			count2++
		}
	}
	fmt.Println(max)
}
