package _92canWinNim

func canWinNim(n int) bool {
	if (n-1)%4 == 0 || (n-2)%4 == 0 || (n-3)%4 == 0 {
		return true
	}
	return false
}
