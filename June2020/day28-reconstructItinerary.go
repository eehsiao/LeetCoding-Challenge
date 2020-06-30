// Given a list of airline tickets represented by pairs of departure and arrival airports [from, to], reconstruct the itinerary in order. All of the tickets belong to a man who departs from JFK. Thus, the itinerary must begin with JFK.
// Note:
// If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when read as a single string. For example, the itinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].
// All airports are represented by three capital letters (IATA code).
// You may assume all tickets form at least one valid itinerary.
// One must use all the tickets once and only once.
// Example 1:
// Input: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
// Output: ["JFK", "MUC", "LHR", "SFO", "SJC"]
// Example 2:
// Input: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
// Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
// Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"].
//              But it is larger in lexical order.

package main

import (
	"sort"
)

var (
	adj map[string][]string
	res []string
)

func findItinerary(tickets [][]string) []string {
	adj = make(map[string][]string)
	res = make([]string, 0)

	for _, v := range tickets {
		adj[v[0]] = append(adj[v[0]], v[1])
	}
	for k, _ := range adj {
		sort.Strings(adj[k])
	}
	dfs("JFK")
	return res
}

func dfs(s string) {
	if _, ok := adj[s]; ok {
		for len(adj[s]) > 0 {
			v := adj[s][0]
			adj[s] = adj[s][1:len(adj[s])]
			dfs(v)
		}
	}

	res = append([]string{s}, res...)
}
