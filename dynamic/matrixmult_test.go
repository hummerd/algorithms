package dynamic

import (
	"fmt"
	"testing"
)

func TestMatrixBestMultiplication(t *testing.T) {
	r := MatrixMult([]int{30, 35, 15, 5, 10, 20, 25})
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i])
	}
}
