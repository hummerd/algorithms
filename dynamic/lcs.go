package dynamic

const (
	L = iota + 1
	T
	LT
)

func Lcs(a, b string) string {
	spl := make([][]int, len(a)+1)
	spd := make([][]int, len(a))

	spl[0] = make([]int, len(b))

	for i := 1; i <= len(a); i++ {
		l := make([]int, len(b)+1)
		spl[i] = l
		spd[i-1] = make([]int, len(b))

		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				l[j] = 1 + spl[i-1][j-1]
				spd[i-1][j-1] = LT
			} else {
				if spl[i-1][j] >= spl[i][j-1] {
					l[j] = spl[i-1][j]
					spd[i-1][j-1] = T
				} else {
					l[j] = spl[i][j-1]
					spd[i-1][j-1] = L
				}
			}
		}
	}

	return rpath(a, b, spd)
}

func rpath(a, b string, spd [][]int) string {
	s := ""

	i, j := len(a)-1, len(b)-1

	for {
		if i < 0 || j < 0 {
			break
		}

		d := spd[i][j]
		if d == LT {
			s = b[j:j+1] + s
			i--
			j--
		} else if d == L {
			j--
		} else {
			i--
		}

	}

	return s
}

// func main() {
// 	r := lcs("abcbdab", "bdcaba")
// 	fmt.Println(r)
// }
