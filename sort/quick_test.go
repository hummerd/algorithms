package sort_test

import (
	stdsort "sort"
	"testing"

	"github.com/hummerd/algorithms/sort"
)

func TestQuick(t *testing.T) {
	tests := []struct {
		name string
		args []int
	}{
		{
			name: "case 1",
			args: []int{12, 45, 3, 1, 0, 78, 122, 45},
		},
		{
			name: "case 2",
			args: []int{12, 45, 3, 1, 250, 78, 122, 45},
		},
		{
			name: "case 3",
			args: []int{12, 45, 3, 1, 56, 78, 122, 45},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Quick(tt.args)

			if !stdsort.IntsAreSorted(tt.args) {
				t.Log(tt.args)
				t.Fail()
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Quick2(tt.args)

			if !stdsort.IntsAreSorted(tt.args) {
				t.Log(tt.args)
				t.Fail()
			}
		})
	}
}

func FuzzQuick(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte) {
		sort.Quick(a)

		if !stdsort.IsSorted(Byte64Slice(a)) {
			t.Fail()
		}
	})
}

type Byte64Slice []byte

func (x Byte64Slice) Len() int { return len(x) }

func (x Byte64Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x Byte64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
