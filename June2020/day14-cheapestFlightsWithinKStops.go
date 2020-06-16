// There are n cities connected by m flights. Each flight starts from city u and arrives at v with a price w.
// Now given all the cities and flights, together with starting city src and the destination dst, your task is to find the cheapest price from src to dst with up to k stops. If there is no such route, output -1.
// Example 1:
// Input:
// n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
// src = 0, dst = 2, k = 1
// Output: 200
// Explanation:
// The graph looks like this:
//         0
//  (100) / \ (500)
//       1 - 2
//       (100)
// The cheapest price from city 0 to city 2 with at most 1 stop costs 200, as marked red in the picture.
// Example 2:
// Input:
// n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
// src = 0, dst = 2, k = 0
// Output: 500
// Explanation:
// The graph looks like this:
//         0
//  (100) / \ (500)
//       1 - 2
//       (100)
// The cheapest price from city 0 to city 2 with at most 0 stop costs 500, as marked blue in the picture.
// Constraints:
// The number of nodes n will be in range [1, 100], with nodes labeled from 0 to n - 1.
// The size of flights will be in range [0, n * (n - 1) / 2].
// The format of each flight will be (src, dst, price).
// The price of each flight will be in the range [1, 10000].
// k is in the range of [0, n - 1].
// There will not be any duplicated flights or self cycles.

package main

import (
	"math"
)

func sliceFill(s []int, v int) {
	for i := 0; i < len(s); i++ {
		s[i] = v
	}
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	dist := make([][]int, 2)
	dist[0], dist[1] = make([]int, n), make([]int, n)
	sliceFill(dist[0], math.MaxInt32)
	sliceFill(dist[1], math.MaxInt32)
	dist[0][src], dist[1][src] = 0, 0

	for i := 0; i < K+1; i++ {
		for _, e := range flights {
			s, d, w := e[0], e[1], e[2]
			dU, dV := dist[1-i&1][s], dist[i&1][d]

			if dU+w < dV {
				dist[i&1][d] = dU + w
			}
		}
	}

	if dist[K&1][dst] < math.MaxInt32 {
		return dist[K&1][dst]
	}

	return -1
}
