package sol

func countBits(n int) []int {
	k := 0
	ans := []int{0}
	for pos := 1; pos <= n; pos++ {
		if pos == (0b1 << k) {
			ans = append(ans, 1)
			k++
		} else {
			ans = append(ans, 1+ans[pos-(0b1<<(k-1))])
		}
	}
	return ans
}
