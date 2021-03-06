// There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.
// Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]
// Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?
// Example 1:
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: true
// Explanation: There are a total of 2 courses to take.
//              To take course 1 you should have finished course 0. So it is possible.
// Example 2:
// Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
// Output: false
// Explanation: There are a total of 2 courses to take.
//              To take course 1 you should have finished course 0, and to take course 0 you should
//              also have finished course 1. So it is impossible.
// Constraints:
// The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
// You may assume that there are no duplicate edges in the input prerequisites.
// 1 <= numCourses <= 10^5
//    Hide Hint #1
// This problem is equivalent to finding if a cycle exists in a directed graph. If a cycle exists, no topological ordering exists and therefore it will be impossible to take all courses.
//    Hide Hint #2
// Topological Sort via DFS - A great video tutorial (21 minutes) on Coursera explaining the basic concepts of Topological Sort.
//    Hide Hint #3
// Topological sort could also be done via BFS.

package main

type gNode struct {
	inDegrees int
	outNodes  []int
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		ttlDeps, edge int   = len(prerequisites), 0
		nodepCourses  []int = make([]int, 0)
	)
	if ttlDeps == 0 {
		return true
	}

	graph := make(map[int]*gNode, 0)

	for _, r := range prerequisites {
		p := createGNode(graph, r[1])
		n := createGNode(graph, r[0])
		p.outNodes = append(p.outNodes, r[0])
		n.inDegrees++
	}

	for v, m := range graph {
		if m.inDegrees == 0 {
			nodepCourses = append(nodepCourses, v)
		}
	}
	for len(nodepCourses) > 0 {
		n := len(nodepCourses) - 1
		course := nodepCourses[n]
		nodepCourses = nodepCourses[:n] // pop

		for _, next := range graph[course].outNodes {
			child := graph[next]
			child.inDegrees--
			edge++
			if child.inDegrees == 0 {
				nodepCourses = append(nodepCourses, next)
			}
		}
	}

	for edge != ttlDeps {
		return false
	}

	return true
}

func createGNode(g map[int]*gNode, c int) *gNode {
	var node *gNode = nil
	if n, ok := g[c]; ok {
		node = n
	} else {
		node = &gNode{
			inDegrees: 0,
			outNodes:  make([]int, 0),
		}
		g[c] = node
	}

	return node
}
