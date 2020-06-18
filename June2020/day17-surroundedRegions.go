// Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
// A region is captured by flipping all 'O's into 'X's in that surrounded region.
// Example:
// X X X X
// X O O X
// X X O X
// X O X X
// After running your function, the board should be:
// X X X X
// X X X X
// X X X X
// X O X X
// Explanation:
// Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.

package main

func DFS(board [][]byte, i, j, m, n int) {
	if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'i'
	DFS(board, i-1, j, m, n)
	DFS(board, i+1, j, m, n)
	DFS(board, i, j-1, m, n)
	DFS(board, i, j+1, m, n)
}

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			DFS(board, i, 0, m, n)
		}
		if board[i][n-1] == 'O' {
			DFS(board, i, n-1, m, n)
		}
	}

	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			DFS(board, 0, j, m, n)
		}
		if board[m-1][j] == 'O' {
			DFS(board, m-1, j, m, n)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'i' {
				board[i][j] = 'O'
			}
		}
	}
}
