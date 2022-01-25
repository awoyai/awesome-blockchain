package fft

import (
	"fmt"
	"github.com/mjibson/go-dsp/fft"
	"testing"
)

func Test_splitByOddAndEven(t *testing.T) {
	even, odd := splitByOddAndEven([]complex128{0, 1, 2, 3, 4, 5})
	fmt.Println(even)
	fmt.Println(odd)
}

func Test_fft(t *testing.T) {
	y := []complex128{1, 2, 3, 4}
	//y := []complex128{1, 2, 3, 4, 5, 6, 7, 8}
	res := FFT(y)
	fmt.Println(res)
	fmt.Println(IFFT(res))
	ff := fft.FFT(y)
	fmt.Println(ff)
	fmt.Println(fft.IFFT(ff))
}

func Test_math(t *testing.T) {
	c := complex(0, 1)
	c2 := complex(-4, 2)
	fmt.Println(c * c2)
	fmt.Println(complex(-2, 0) - c*c2)
	fmt.Println(c2 / complex(-2, 1))
}
