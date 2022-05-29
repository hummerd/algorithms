package dynamic

func MatrixMult(p []int) [][]int {
	l := make([][]int, len(p)-1)

	mm := make([]int, len(p)-1)
	l[0] = mm

	for i := 0; i < len(p)-1; i++ {
		mm[i] = 0
	}

	for i := 2; i <= len(p)-1; i++ { // loop over length of subset with length i
		mm := make([]int, len(p)-i)
		l[i-1] = mm

		for j := 0; j < len(p)-i; j++ { // subset starts with j

			mmin := -1
			for k := 1; k < i; k++ { // now check every subset and find min
				ll := l[k-1][j]                   // count for subset of length k-1, that starts with j
				lr := l[i-k-1][j+k]               // count for subset of length i-k-1, that starts with j+k
				m := ll + lr + p[j]*p[j+k]*p[j+i] // sum of subset counts + count for mult of subsets

				if mmin == -1 || mmin > m {
					mmin = m
				}
			}

			mm[j] = mmin
		}
	}

	return l
}
