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

	var H, W int
	fmt.Fscan(r, &H)
	fmt.Fscan(r, &W)

	strSlice := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Fscan(r, &strSlice[i])
	}

	// iは行、jは列
	boundaryNum := 4
Loop:
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if string(strSlice[i][j]) == "s" {
				// 左
				if (j >= boundaryNum && H >= boundaryNum+1) && string(strSlice[i][j-1]) == "n" && string(strSlice[i][j-2]) == "u" && string(strSlice[i][j-3]) == "k" && string(strSlice[i][j-4]) == "e" {
					i++
					j++
					fmt.Printf("%d %d\n", i, j)
					fmt.Printf("%d %d\n", i, j-1)
					fmt.Printf("%d %d\n", i, j-2)
					fmt.Printf("%d %d\n", i, j-3)
					fmt.Printf("%d %d\n", i, j-4)
					break Loop
				}
				// 左上
				if (i >= boundaryNum && H >= boundaryNum+1) && (j >= boundaryNum && W >= boundaryNum) && string(strSlice[i-1][j-1]) == "n" && string(strSlice[i-2][j-2]) == "u" && string(strSlice[i-3][j-3]) == "k" && string(strSlice[i-4][j-4]) == "e" {
					i++
					j++
					fmt.Printf("%d %d\n", i, j)
					fmt.Printf("%d %d\n", i-1, j-1)
					fmt.Printf("%d %d\n", i-2, j-2)
					fmt.Printf("%d %d\n", i-3, j-3)
					fmt.Printf("%d %d\n", i-4, j-4)
					break Loop
				}
				// 上
				if (i >= boundaryNum && H >= boundaryNum+1) && string(strSlice[i-1][j]) == "n" && string(strSlice[i-2][j]) == "u" && string(strSlice[i-3][j]) == "k" && string(strSlice[i-4][j]) == "e" {
					i++
					j++
					fmt.Printf("%d %d\n", i, j)
					fmt.Printf("%d %d\n", i-1, j)
					fmt.Printf("%d %d\n", i-2, j)
					fmt.Printf("%d %d\n", i-3, j)
					fmt.Printf("%d %d\n", i-4, j)
					break Loop
				}
				// 右上
				if (i >= boundaryNum && H >= boundaryNum+1) && (j >= boundaryNum && W >= boundaryNum) && string(strSlice[i-1][j+1]) == "n" && string(strSlice[i-2][j+2]) == "u" && string(strSlice[i-3][j+3]) == "k" && string(strSlice[i-4][j+4]) == "e" {
					i++
					j++
					fmt.Printf("%d %d\n", i, j)
					fmt.Printf("%d %d\n", i-1, j+1)
					fmt.Printf("%d %d\n", i-2, j+2)
					fmt.Printf("%d %d\n", i-3, j+3)
					fmt.Printf("%d %d\n", i-4, j+4)
					break Loop
				}
				// 右
				if (j >= boundaryNum && W >= boundaryNum) && string(strSlice[i][j+1]) == "n" && string(strSlice[i][j+2]) == "u" && string(strSlice[i][j+3]) == "k" && string(strSlice[i][j+4]) == "e" {
					i++
					j++
					fmt.Printf("%d %d\n", i, j)
					fmt.Printf("%d %d\n", i, j+1)
					fmt.Printf("%d %d\n", i, j+2)
					fmt.Printf("%d %d\n", i, j+3)
					fmt.Printf("%d %d\n", i, j+4)
					break Loop
				}
				// 右下
				if (i >= boundaryNum && H >= boundaryNum+1) && (j >= boundaryNum && W >= boundaryNum) && string(strSlice[i+1][j+1]) == "n" && string(strSlice[i+2][j+2]) == "u" && string(strSlice[i+3][j+3]) == "k" && string(strSlice[i+4][j+4]) == "e" {
					i++
					j++
					fmt.Printf("%d %d\n", i, j)
					fmt.Printf("%d %d\n", i+1, j+1)
					fmt.Printf("%d %d\n", i+2, j+2)
					fmt.Printf("%d %d\n", i+3, j+3)
					fmt.Printf("%d %d\n", i+4, j+4)
					break Loop
				}
				// 下
				if (j >= boundaryNum && W >= boundaryNum) && string(strSlice[i+1][j]) == "n" && string(strSlice[i+2][j]) == "u" && string(strSlice[i+3][j]) == "k" && string(strSlice[i+4][j]) == "e" {
					i++
					j++
					fmt.Printf("%d %d\n", i, j)
					fmt.Printf("%d %d\n", i+1, j)
					fmt.Printf("%d %d\n", i+2, j)
					fmt.Printf("%d %d\n", i+3, j)
					fmt.Printf("%d %d\n", i+4, j)
					break Loop
				}
				// 左下
				if (i >= boundaryNum && H >= boundaryNum+1) && (j >= boundaryNum && W >= boundaryNum) && string(strSlice[i+1][j-1]) == "n" && string(strSlice[i+2][j-2]) == "u" && string(strSlice[i+3][j-3]) == "k" && string(strSlice[i+4][j-4]) == "e" {
					i++
					j++
					fmt.Printf("%d %d\n", i, j)
					fmt.Printf("%d %d\n", i+1, j-1)
					fmt.Printf("%d %d\n", i+2, j-2)
					fmt.Printf("%d %d\n", i+3, j-3)
					fmt.Printf("%d %d\n", i+4, j-4)
					break Loop
				}
			}
		}
	}
}
