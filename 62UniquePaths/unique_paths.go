package main

import "log"

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	m, n = m+n-2, m-1

	var up, down2 int64 = 1, 1
	c := 0
	if n > m/2 {
		c = n
	} else {
		c = m - n
	}

	for i := c+1; i <= m; i++ {
		up = up * int64(i)
	}

	for i := 1; i <= m-c; i++ {
		down2 = down2 * int64(i)
	}
	return int(up /  down2)
}

func main() {
	log.Println(uniquePaths(7, 3))
}
