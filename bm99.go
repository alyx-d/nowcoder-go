package main

/*
1 2 3
4 5 6
7 8 9

7 4 1
8 5	2
9 6	3

7 8 9
4 5 6
1 2 3
*/
func RotateMatrix(mat [][]int, n int) [][]int {
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; {
			mat[j][i], mat[k][i] = mat[k][i], mat[j][i]
			j++
			k--
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}
	return mat
}
