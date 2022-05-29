// Insertion sort - Chapter 2.1.
package sort

func Insertion(a []int) {
	if len(a) <= 1 {
		return
	}

	// invariant: everuthing at left of i is sorted
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j-1] <= a[j] {
				break
			}
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
}
