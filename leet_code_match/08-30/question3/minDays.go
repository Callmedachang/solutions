package question3

func minDays(grid [][]int) int {
	if
}

func IsFL(grid [][]int, i, j int) bool {
	if i == 0 {
		if j == 0 {
			if grid[i+1][j] == 0 && grid[i][j+1] == 0 {
				return true
			}
		} else {
			if grid[i][j-0] == 0 && grid[i+1][j] == 0 && grid[i][j+1] == 0 {
				return true
			}
		}
	}
	if i == len(grid)-1 {

	}
	if j == 0 {
		if grid[i][j-0] == 0 && grid[i+1][j] == 0 && grid[i][j+1] == 0 {
			return true
		}
	}
}
