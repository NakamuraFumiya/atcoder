package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)

	problems := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &problems[i])
	}

	playerScores := make([]int, n)
	maxScore := -1
	// keyがプレイヤーのindex, valueがまだ解いてない問題のindexes
	notSolvedProblemsMap := make(map[int][]int)
	for i := 0; i < n; i++ {
		var solvedProblemsStr string
		fmt.Fscan(r, &solvedProblemsStr)

		solvedProblems := strings.Split(solvedProblemsStr, "")
		score := 0
		for j, v := range solvedProblems {
			if v == "o" {
				score += problems[j]
			} else {
				notSolvedProblemsMap[i] = append(notSolvedProblemsMap[i], j)
			}
		}
		score += i + 1
		playerScores[i] = score

		if score > maxScore {
			maxScore = score
		}
	}

	// 答えを出力
	for i := 0; i < n; i++ {
		notSolvedProblems := make([]int, 0)
		for _, v := range notSolvedProblemsMap[i] {
			notSolvedProblems = append(notSolvedProblems, problems[v])
		}
		// まだ解いてない問題のスコアの降順のリストにする
		sort.Slice(notSolvedProblems, func(i, j int) bool { return notSolvedProblems[i] > notSolvedProblems[j] })

		count := numberOfProblemsToSolve(playerScores[i], maxScore, notSolvedProblems, 0, 0)
		fmt.Println(count)
	}
}

func numberOfProblemsToSolve(score int, maxScore int, problems []int, index int, count int) int {
	if score >= maxScore {
		return count
	}
	score += problems[index]
	return numberOfProblemsToSolve(score, maxScore, problems, index+1, count+1)
}
