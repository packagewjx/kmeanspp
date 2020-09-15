package utils

import "math"

func DistanceSquare(p, q []float32) float32 {
	sum := float32(0)
	i := 0
	for ; i+3 < len(p); i += 4 {
		d1 := p[i] - q[i]
		d2 := p[i+1] - q[i+1]
		d3 := p[i+2] - q[i+2]
		d4 := p[i+3] - q[i+3]

		sum += d1*d1 + d2*d2 + d3*d3 + d4*d4
	}

	for ; i < len(p); i++ {
		d := p[i] - q[i]
		sum += d * d
	}

	return sum
}

func SumFloat32(arr []float32) float32 {
	i := 0
	sum := float32(0)
	for ; i+3 < len(arr); i += 4 {
		sum += arr[i]
		sum += arr[i+1]
		sum += arr[i+2]
		sum += arr[i+3]
	}

	for ; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func SumToA(a, b []float32) {
	i := 0
	for ; i+3 < len(a); i += 4 {
		a[i] += b[i]
		a[i+1] += b[i+1]
		a[i+2] += b[i+2]
		a[i+3] += b[i+3]
	}

	for ; i < len(a); i++ {
		a[i] += b[i]
	}
}

func SumToA2DimensionFloat32(a, b [][]float32) {
	for i := 0; i < len(a); i++ {
		SumToA(a[i], b[i])
	}
}

func SumToAInt(a, b []int) {
	i := 0
	for ; i+3 < len(a); i += 4 {
		a[i] += b[i]
		a[i+1] += b[i+1]
		a[i+2] += b[i+2]
		a[i+3] += b[i+3]
	}
	for ; i < len(a); i++ {
		a[i] += b[i]
	}
}

// Find the first position equal to or bigger than k
func BinarySearchFloat32(arr []float32, k float32) int {
	if len(arr) == 0 {
		return 0
	}
	if arr[len(arr)-1] < k {
		return len(arr)
	}

	l := 0
	r := len(arr) - 1
	for l < r {
		m := (l + r) / 2
		if arr[m] >= k {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func Make2DimensionFloat32Array(m, n int) [][]float32 {
	res := make([][]float32, m)
	for i := 0; i < m; i++ {
		res[i] = make([]float32, n)
	}

	return res
}

func ClosestPointDistance(p []float32, data [][]float32) float32 {
	minDist := float32(math.MaxFloat32)
	for _, q := range data {
		dist := DistanceSquare(p, q)
		// ensure p != q
		if dist == 0 {
			continue
		}
		if minDist > dist {
			minDist = dist
		}
	}

	return float32(math.Sqrt(float64(minDist)))
}
