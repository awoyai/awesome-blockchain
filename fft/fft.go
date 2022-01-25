package fft

import "math"

func FFT(p []complex128) []complex128 {
	n := len(p)
	if n == 1 {
		return p
	}
	// TODO: 创建欧米伽的计算算法
	omg := NewOmg(n)
	Pe, Po := splitByOddAndEven(p)
	ye, yo := FFT(Pe), FFT(Po)
	y := make([]complex128, n)
	for i := 0; i < n/2; i++ {
		pow := omg.Pow(float64(i))
		y[i] = yo[i] + pow*ye[i]
		y[i+n/2] = yo[i] - pow*ye[i]
	}
	return y
}

func IFFT(p []complex128) []complex128 {
	n := len(p)
	if n == 1 {
		return p
	}
	// TODO: 创建欧米伽的计算算法
	omg := NewOmg(n)
	Pe, Po := splitByOddAndEven(p)
	ye, yo := IFFT(Pe), IFFT(Po)
	y := make([]complex128, n)
	for i := 0; i < n/2; i++ {
		pow := omg.IPow(float64(i))/complex(float64(n), 0)
		y[i] = yo[i] + pow*ye[i]
		y[i+n/2] = yo[i] - pow*ye[i]
	}
	return y
}

type Omg struct {
	N int
}

func NewOmg(n int) Omg {
	return Omg{N: n}
}

func (o *Omg) Pow(i float64) (res complex128) {
	sita := 360 * i / float64(o.N)
	cos45 := math.Sqrt(2) / 2
	switch sita {
	case 0:
		res = complex(1, 0)
	case 45:
		res = complex(cos45, -cos45)
	case 90:
		res = complex(0, -1)
	case 135:
		res = complex(-cos45, -cos45)
	case 180:
		res = complex(-1, 0)
	case 225:
		res = complex(-cos45, cos45)
	case 270:
		res = complex(0, 1)
	case 315:
		res = complex(cos45, cos45)
	}
	return
}

func (o *Omg) IPow(i float64) (res complex128) {
	sita := 360 * i / float64(o.N)
	cos45 := math.Sqrt(2) / 2
	switch sita {
	case 0:
		res = complex(1, 0)
	case 45:
		res = complex(cos45, cos45)
	case 90:
		res = complex(0, 1)
	case 135:
		res = complex(cos45, -cos45)
	case 180:
		res = complex(-1, 0)
	case 225:
		res = complex(-cos45, -cos45)
	case 270:
		res = complex(0, -1)
	case 315:
		res = complex(cos45, -cos45)
	}
	return
}

func splitByOddAndEven(p []complex128) ([]complex128, []complex128) {
	n := len(p)
	even := make([]complex128, n/2)
	odd := make([]complex128, n/2)
	for i := 0; i < n; i += 2 {
		even[i/2] = p[i+1]
		odd[i/2] = p[i]
	}
	return even, odd
}
