package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int
	fmt.Fscan(r, &N)

	s := make([]int, N)
	for i := range s {
		fmt.Fscan(r, &s[i])
	}

	ss := make([]int, 0)
	for i := 0; i < N; i++ {
		ss = append(ss, s[i])
		j := i + 1
		if j >= N {
			break
		}
		if s[i] > s[j] {
			abs := int(math.Abs(float64(s[i] - s[j])))
			if abs == 1 {
				continue
			} else {
				// sliceの操作
				tmp := make([]int, abs-1)
				for k := range tmp {
					if k == 0 {
						tmp[k] = s[i] - 1
					} else {
						tmp[k] = tmp[k-1] - 1
					}
				}

				ss = append(ss, tmp...)
			}
		} else if s[i] < s[j] {
			abs := int(math.Abs(float64(s[i] - s[j])))
			if abs == 1 {
				continue
			} else {
				// sliceの操作
				tmp := make([]int, abs-1)
				for k := range tmp {
					if k == 0 {
						tmp[k] = s[i] + 1
					} else {
						tmp[k] = tmp[k-1] + 1
					}
				}
				ss = append(ss, tmp...)
			}
		}
	}

	res := ""
	for i := range ss {
		if i == 0 {
			res += fmt.Sprintf("%d", ss[i])
		} else {
			res += fmt.Sprintf(" %d", ss[i])
		}
	}
	fmt.Println(res)
}
