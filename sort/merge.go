// Merge sort - chapter 2.3.
package sort

func Merge(a []int) {
	s := 1

	tmp := make([]int, len(a))
	for s < len(a) {
		i := 0
		for i < len(a) {
			fi := a[i : i+s]
			se := a[i+s : i+2*s]
			merge(fi, se, tmp[i:])
			i += 2 * s
		}
		a, tmp = tmp, a
		s = s * 2
	}
	copy(tmp, a)
}

func merge(a, b, c []int) {
	i, j, k := 0, 0, 0
	for {
		if a[i] > b[j] {
			c[k] = b[j]
			j++
		} else {
			c[k] = a[i]
			i++
		}

		k++

		if i == len(a) {
			for ; j < len(b); j++ {
				c[k] = b[j]
				k++
			}
			return
		}

		if j == len(b) {
			for ; i < len(a); i++ {
				c[k] = a[i]
				k++
			}
			return
		}
	}
}
