package fft

import (
	"fmt"
	"testing"
)

func Test_splitByOddAndEven(t *testing.T) {
	even, odd := splitByOddAndEven([]int{0, 1, 2, 3, 4, 5})
	fmt.Println(even)
	fmt.Println(odd)
}
