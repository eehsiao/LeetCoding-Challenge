// Given a set of N people (numbered 1, 2, ..., N), we would like to split everyone into two groups of any size.
// Each person may dislike some other people, and they should not go into the same group.
// Formally, if dislikes[i] = [a, b], it means it is not allowed to put the people numbered a and b into the same group.
// Return true if and only if it is possible to split everyone into two groups in this way.
// Example 1:
// Input: N = 4, dislikes = [[1,2],[1,3],[2,4]]
// Output: true
// Explanation: group1 [1,4], group2 [2,3]
// Example 2:
// Input: N = 3, dislikes = [[1,2],[1,3],[2,3]]
// Output: false
// Example 3:
// Input: N = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
// Output: false
// Note:
// 1 <= N <= 2000
// 0 <= dislikes.length <= 10000
// 1 <= dislikes[i][j] <= N
// dislikes[i][0] < dislikes[i][1]
// There does not exist i != j for which dislikes[i] == dislikes[j].

package main

import "fmt"

var (
	graph [][]int
	color map[int]int
)

func dfs(n, c int) bool {
	if v, ok := color[n]; ok {
		return v == c
	}
	color[n] = c

	for _, nei := range graph[n] {
		if !dfs(nei, c^1) {
			return false
		}
	}
	return true
}

func possibleBipartition(N int, dislikes [][]int) bool {
	graph = make([][]int, N+1)
	color = make(map[int]int)
	for i := 1; i <= N; i++ {
		graph[i] = make([]int, 0)
	}

	for _, edge := range dislikes {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	fmt.Println(graph)

	for n := 1; n <= N; n++ {
		if _, ok := color[n]; !ok && !dfs(n, 0) {
			fmt.Println(color)
			return false
		}
	}
	fmt.Println(color)
	return true
}
