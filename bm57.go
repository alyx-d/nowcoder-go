package main

func Bm57Solve(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	if m == 0 {
		return 0
	}
	var dfs func(int, int)
	dfs = func(x, y int) {
		grid[x][y] = '0'
		if y-1 >= 0 && grid[x][y-1] == '1' {
			dfs(x, y-1)
		}
		if y+1 < n && grid[x][y+1] == '1' {
			dfs(x, y+1)
		}
		if x-1 >= 0 && grid[x-1][y] == '1' {
			dfs(x-1, y)
		}
		if x+1 < m && grid[x+1][y] == '1' {
			dfs(x+1, y)
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

type Pos struct {
	X int
	Y int
}

func Bm57Solve2(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	if m == 0 {
		return 0
	}
	queue := make([]Pos, 0)
	bfs := func() {
		for len(queue) > 0 {
			p := queue[0]
			grid[p.X][p.Y] = '0'
			queue = queue[1:]
			if p.X-1 >= 0 && grid[p.X-1][p.Y] == '1' {
				queue = append(queue, Pos{p.X - 1, p.Y})
			}
			if p.X+1 < m && grid[p.X+1][p.Y] == '1' {
				queue = append(queue, Pos{p.X + 1, p.Y})
			}
			if p.Y-1 >= 0 && grid[p.X][p.Y-1] == '1' {
				queue = append(queue, Pos{p.X, p.Y - 1})
			}
			if p.Y+1 < n && grid[p.X][p.Y+1] == '1' {
				queue = append(queue, Pos{p.X, p.Y + 1})
			}
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				queue = append(queue, Pos{i, j})
				bfs()
			}
		}
	}
	return count
}

type UnionFind struct {
	fx    []int
	count int
}

func (uf *UnionFind) init(grid [][]byte) {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				uf.count++
				uf.fx = append(uf.fx, i*n+j)
			} else {
				uf.fx = append(uf.fx, -1)
			}
		}
	}
}

func (uf *UnionFind) find(x int) int {
	if uf.fx[x] != x {
		uf.fx[x] = uf.find(uf.fx[x])
		return uf.fx[x]
	}
	return x
}

func (uf *UnionFind) union(a, b int) {
	x := uf.find(a)
	y := uf.find(b)
	if x != y {
		uf.fx[x] = y
		uf.count--
	}
}

func Bm57Solve3(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	if m == 0 {
		return 0
	}
	uf := new(UnionFind)
	uf.init(grid)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				if i-1 >= 0 && grid[i-1][j] == '1' {
					uf.union(i*n+j, (i-1)*n+j)
				}
				if i+1 < m && grid[i+1][j] == '1' {
					uf.union(i*n+j, (i+1)*n+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					uf.union(i*n+j, i*n+j-1)
				}
				if j+1 < n && grid[i][j+1] == '1' {
					uf.union(i*n+j, i*n+j+1)
				}
			}
		}
	}
	return uf.count
}
