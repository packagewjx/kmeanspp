package kmeanspp

import (
	"github.com/packagewjx/kmeanspp/internal"
)

// K-Means++ Entrance function. k is the number of classes. round is the number of iteration. data array contains
// all instances data. data[i][j] is jth feature of ith instance.
func KMeansPP(k, round int, data [][]float32) (centers [][]float32, class []int) {
	if k <= 0 || round <= 0 {
		return [][]float32{}, []int{}
	}

	centerIndices := internal.InitialCenterPureGo(k, data)
	centers = make([][]float32, k)
	for i, ci := range centerIndices {
		centers[i] = data[ci]
	}

	for r := 0; r < round; r++ {
		centers, class = internal.CalculateNewCentersPureGo(data, centers)
	}

	return centers, class
}
