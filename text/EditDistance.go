package text

const (
	maxDistance = 0x3fffffff
)

type Row []int

// MinEditDistance 计算两个字符串的最短编辑距离
func MinEditDistance(str1, str2 string) int {
	r1, r2 := []rune(str1), []rune(str2)
	return minEditDistance(r1, r2)
}

func StrEditSimilarity(str1, str2 string) float64 {
	r1, r2 := []rune(str1), []rune(str2)
	med := minEditDistance(r1, r2)
	return float64(med*2) / float64(len(str1)+len(str2))
}

func minEditDistance(r1, r2 []rune) int {
	l1, l2 := len(r1), len(r2)
	a := genArray(l1, l2)
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if r1[i-1] == r2[j-1] {
				a[i][j] = a[i-1][j-1]
			} else {
				a[i][j] = min(a[i-1][j-1], min(a[i-1][j], a[i][j-1])) + 1
			}
		}
	}
	return a[l1][l2]
}

func genArray(l1, l2 int) []Row {
	a := make([]Row, l1+1)
	for i := 0; i <= l1; i++ {
		a[i] = make(Row, l2+1)
		for j := 0; j <= l2; j++ {
			a[i][j] = maxDistance
			if i == 0 || j == 0 {
				a[i][j] = max(i, j)
			}
		}
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
