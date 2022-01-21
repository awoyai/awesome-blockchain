package fft

func FFT(p []int) []int {
	n := len(p)
	if n == 1 {
		return p
	}
	// TODO: 创建欧米伽的计算算法
	w := 0
	Pe, Po := splitByOddAndEven(p)
	ye, yo := FFT(Pe), FFT(Po)
	y := make([]int, n)
	for i := 0; i < n/2; i++ {
		y[i] = ye[i] + w*yo[i]
		y[i+n/2] = ye[i] - w*yo[i]
	}
	return y
}

type Omg struct {
	Sita int
	N int
}

func NewOmg(n int) Omg {
	return Omg{N: n}
}



func splitByOddAndEven(p []int) ([]int, []int) {
	n := len(p)
	even := make([]int, n/2)
	odd := make([]int, n/2)
	for i := 0; i < n; i+=2 {
		even[i/2] = p[i+1]
		odd[i/2] = p[i]
	}
	return even, odd
}


